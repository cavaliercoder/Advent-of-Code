#!/usr/bin/env python3

import logging

from enum import Enum
from typing import Any, Dict, List, Optional, Sequence, Tuple


Data = List[int]

Address = int


class AddressMode(Enum):
    POSITION = 0
    IMMEDIATE = 1
    RELATIVE = 2

    def __str__(self) -> str:
        return self.name


class Opcode(Enum):
    UNKNOWN = 0
    ADD = 1
    MULTIPLY = 2
    INPUT = 3
    OUTPUT = 4
    JCC = 5
    JNC = 6
    LT = 7
    EQ = 8
    RELBASE = 9
    HALT = 99


OPCODE_SIGNATURES: Dict[Opcode, Tuple[str, List[str]]] = {
    Opcode.UNKNOWN: ("unknown", []),
    Opcode.ADD: ("add", ["a", "b", "addr"]),
    Opcode.MULTIPLY: ("multiply", ["a", "b", "addr"]),
    Opcode.INPUT: ("input", ["addr"]),
    Opcode.OUTPUT: ("output", ["addr"]),
    Opcode.JCC: ("jump-if-true", ["v", "addr"]),
    Opcode.JNC: ("jump-if-not-true", ["v", "addr"]),
    Opcode.LT: ("less-than", ["a", "b", "addr"]),
    Opcode.EQ: ("equal", ["a", "b", "addr"]),
    Opcode.RELBASE: ("adjust-rel-base", ["offset"]),
    Opcode.HALT: ("halt", []),
}


logger = logging.getLogger(__name__)


def decode(s: str) -> Data:
    return [int(opcode) for opcode in s.split(",")]


class IntcodeVM:
    def __init__(
        self, data: Optional[Data] = None, stdin: Optional[Data] = None
    ) -> None:
        self.handlers = {
            Opcode.ADD: self._op_add,
            Opcode.MULTIPLY: self._op_multiply,
            Opcode.INPUT: self._op_input,
            Opcode.OUTPUT: self._op_output,
            Opcode.JCC: self._op_jump_if_true,
            Opcode.JNC: self._op_jump_if_not_true,
            Opcode.LT: self._op_less_than,
            Opcode.EQ: self._op_equal,
            Opcode.RELBASE: self._op_relbase,
        }
        self.reset(data, stdin)

    def reset(self, data: Optional[Data], stdin: Optional[Data] = None) -> None:
        # program counter
        self.reg_pc = Address(0)

        # dedicated register for the relative offset base
        self.reg_rel_base = Address(0)

        # vm memory
        self.data = data.copy() if data else []

        # input buffer
        self.stdin: Data = stdin.copy() if stdin else []

        # output buffer
        self.stdout: Data = []

        # current instruction opcode
        self.opcode: Opcode = Opcode.UNKNOWN

        # current instruction parameter address modes
        self.param_modes: List[AddressMode] = [AddressMode.POSITION] * 4

    def _trace(
        self,
    ) -> None:
        sig = OPCODE_SIGNATURES.get(self.opcode)
        if sig is None:
            sig = OPCODE_SIGNATURES[Opcode.UNKNOWN]
        opcode = sig[0]
        params = sig[1]

        def encode_param(i: int) -> str:
            addr = self.reg_pc + 1 + i
            mode = self.param_modes[i]
            name = params[i]
            v = self.data[addr]
            if mode == AddressMode.POSITION:
                return f"{name}={v}"
            if mode == AddressMode.IMMEDIATE:
                return f"{name}=#{v}"
            if mode == AddressMode.RELATIVE:
                return f"{name}={v}+${self.reg_rel_base}"
            raise RuntimeError(f"Unrecognized address mode: {mode}")

        params_str = ", ".join([encode_param(i) for i in range(len(params))])
        raw_str = ",".join(
            str(i) for i in self.data[self.reg_pc : self.reg_pc + len(params) + 1]
        )
        logger.debug(f"[{self.reg_pc:04}] {opcode}({params_str})  # {raw_str}")

    @property
    def halted(self) -> bool:
        return self.opcode == Opcode.HALT

    def step(self) -> bool:
        """
        Fetch, decode and execute a single instruction.
        """
        if self.halted:
            raise RuntimeError("Halted")

        # fetch instruction
        v = self.read(self.reg_pc)

        # decode opcode and parameter address modes
        self.opcode = Opcode(v % 100)
        v = int(v / 100)
        for i in range(len(self.param_modes)):
            self.param_modes[i] = AddressMode(v % 10)
            v = int(v / 10)
        self._trace()

        # dispatch
        if self.opcode == Opcode.HALT:
            return False
        if self.opcode not in self.handlers:
            raise RuntimeError(f"Unrecognized opcode: {self.opcode}")
        self.handlers[self.opcode]()
        return True

    def run(self) -> None:
        """
        Call step repeatedly until the VM is halted.
        """
        while self.step():
            pass

    def io_push(self, v: int) -> None:
        """
        Push a new value to the VM's input buffer.
        """
        self.stdin.append(v)

    def io_pop(self) -> int:
        """
        Pop a value from the VM's output buffer.
        """
        v = self.stdout[0]
        self.stdout = self.stdout[1:]
        return v

    def effective_addr(self, addr: int, mode: AddressMode) -> int:
        if mode == AddressMode.IMMEDIATE:
            return addr
        if mode == AddressMode.POSITION:
            return self.data[addr]
        if mode == AddressMode.RELATIVE:
            return self.data[addr] + self.reg_rel_base
        raise RuntimeError(f"Unrecognized address mode: {mode}")

    def grow_mem(self, addr: int) -> int:
        old_size = len(self.data)
        if addr < old_size:
            return old_size
        new_size = 64
        while new_size < addr:
            new_size *= 2
        if new_size > 4096:
            raise RuntimeError("Too much memory allocated")
        new_data = [0] * new_size
        for i in range(old_size):
            new_data[i] = self.data[i]
        self.data = new_data
        logger.debug(f"  grow_mem(addr={addr}) -> {new_size}")
        return new_size

    def read(self, addr: int, mode=AddressMode.IMMEDIATE) -> int:
        if addr != AddressMode.IMMEDIATE:
            addr = self.effective_addr(addr, mode)
        assert addr >= 0, f"Invalid address: {addr}"
        if addr >= len(self.data):
            v = 0
        else:
            v = self.data[addr]
        logger.debug(f"  read(addr={addr}) -> {v}")
        return v

    def write(self, addr: int, v: int) -> None:
        self.grow_mem(addr)
        self.data[addr] = v
        logger.debug(f"  write(addr={addr}, v={v})")

    def _op_add(self) -> None:
        """
        Opcode 1 adds together numbers read from two positions and stores the result in a third
        position. The three integers immediately after the opcode tell you these three positions
         - the first two indicate the positions from which you should read the input values, and the
        third indicates the position at which the output should be stored.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        b = self.read(self.reg_pc + 2, self.param_modes[1])
        addr = self.effective_addr(self.reg_pc + 3, self.param_modes[2])
        v = a + b
        self.write(addr, v)
        self.reg_pc += 4

    def _op_multiply(self) -> None:
        """
        Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding
        them. Again, the three integers after the opcode indicate where the inputs and outputs are,
        not their values.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        b = self.read(self.reg_pc + 2, self.param_modes[1])
        addr = self.effective_addr(self.reg_pc + 3, self.param_modes[2])
        v = a * b
        self.write(addr, v)
        self.reg_pc += 4

    def _op_input(self) -> None:
        if self.stdin:
            v = self.stdin[0]
            self.stdin = self.stdin[1:]
        else:
            v = int(input(f"[{self.reg_pc:04}] Enter input: "))
        addr = self.effective_addr(self.reg_pc + 1, self.param_modes[0])
        self.write(addr, v)
        self.reg_pc += 2

    def _op_output(self) -> None:
        v = self.read(self.reg_pc + 1, self.param_modes[0])
        self.stdout.append(v)
        self.reg_pc += 2

    def _op_jump_if_true(self) -> None:
        """
        Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction
        pointer to the value from the second parameter. Otherwise, it does nothing.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        addr = self.read(self.reg_pc + 2, self.param_modes[1])
        v = a != 0
        if v:
            self.reg_pc = addr
        else:
            self.reg_pc += 3

    def _op_jump_if_not_true(self) -> None:
        """
        Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer
        to the value from the second parameter. Otherwise, it does nothing.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        addr = self.read(self.reg_pc + 2, self.param_modes[1])
        v = a == 0
        if v:
            self.reg_pc = addr
        else:
            self.reg_pc += 3

    def _op_less_than(self) -> None:
        """
        Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1
        in the position given by the third parameter. Otherwise, it stores 0.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        b = self.read(self.reg_pc + 2, self.param_modes[1])
        addr = self.effective_addr(self.reg_pc + 3, self.param_modes[2])
        v = int(a < b)
        self.write(addr, v)
        self.reg_pc += 4

    def _op_equal(self) -> None:
        """
        Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in
        the position given by the third parameter. Otherwise, it stores 0.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        b = self.read(self.reg_pc + 2, self.param_modes[1])
        addr = self.effective_addr(self.reg_pc + 3, self.param_modes[2])
        v = int(a == b)
        self.write(addr, v)
        self.reg_pc += 4

    def _op_relbase(self) -> None:
        """
        Opcode 9 adjusts the relative base by the value of its only parameter. The relative base
        increases (or decreases, if the value is negative) by the value of the parameter.
        """
        a = self.read(self.reg_pc + 1, self.param_modes[0])
        self.reg_rel_base += a
        self.reg_pc += 2


def run(data: Data, stdin: Optional[Data] = None) -> Data:
    vm = IntcodeVM(data=data, stdin=stdin)
    vm.run()
    return vm.stdout


def load_and_run(file: str, stdin: Optional[Data] = None) -> Data:
    with open(file, "r") as fp:
        data = decode(fp.readline())
    return run(data, stdin=stdin)

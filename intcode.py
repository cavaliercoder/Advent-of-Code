#!/usr/bin/env python3

import logging

from typing import List, Optional, Sequence, Tuple


ADDR_MODE_POS = 0
ADDR_MODE_IMM = 1
ADDR_MODE_REL = 2

OPCODE_ADD = 1
OPCODE_MULTIPLY = 2
OPCODE_INPUT = 3
OPCODE_OUTPUT = 4
OPCODE_JCC = 5
OPCODE_JNC = 6
OPCODE_LT = 7
OPCODE_EQ = 8
OPCODE_RELBASE = 9
OPCODE_HALT = 99


Data = List[int]


logger = logging.getLogger(__name__)


def decode(s: str) -> Data:
    return [int(opcode) for opcode in s.split(",")]


class Instruction:
    def __init__(self, v: int) -> None:
        self.opcode: int = v % 100
        self.params: Sequence[int] = []

        v = int(v / 100)
        while v:
            self.params.append(v % 10)
            v = int(v / 10)

    def get_param(self, n: int) -> int:
        if n >= len(self.params):
            return 0
        return self.params[n]


class IntcodeVM:
    def __init__(
        self, data: Optional[Data] = None, stdin: Optional[Data] = None
    ) -> None:
        self.handlers = {
            OPCODE_ADD: self.op_add,
            OPCODE_MULTIPLY: self.op_multiply,
            OPCODE_INPUT: self.op_input,
            OPCODE_OUTPUT: self.op_output,
            OPCODE_JCC: self.op_jcc,
            OPCODE_JNC: self.op_jnc,
            OPCODE_LT: self.op_lt,
            OPCODE_EQ: self.op_eq,
            OPCODE_RELBASE: self.op_relbase,
        }
        self.reset(data, stdin)

    @property
    def halted(self) -> bool:
        fn = Instruction(self.data[self.ptr])
        return fn.opcode == OPCODE_HALT

    def reset(self, data: Optional[Data], stdin: Optional[Data] = None) -> None:
        self.ptr = 0
        self.reg_rel_base = 0
        self.data = data.copy() if data else []
        self.stdin: Data = stdin.copy() if stdin else []
        self.stdout: Data = []

    def step(self) -> bool:
        """
        Fetch, decode and execute a single instruction.
        """
        fn = Instruction(self.data[self.ptr])
        if fn.opcode == OPCODE_HALT:
            logger.debug(f"[{self.ptr:04}] HALT")
            return False
        if fn.opcode not in self.handlers:
            raise RuntimeError(f"Unrecognized opcode: {fn.opcode}")
        self.handlers[fn.opcode](fn)
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

    def decode_addr(self, addr: int, mode: int) -> int:
        if mode == ADDR_MODE_IMM:
            return addr
        if mode == ADDR_MODE_POS:
            return self.data[addr]
        if mode == ADDR_MODE_REL:
            return self.data[addr] + self.reg_rel_base
        raise RuntimeError(f"Unrecognized address mode: {mode}")

    def grow_mem(self, addr: int) -> int:
        old_size = len(self.data)
        if addr < old_size:
            return old_size
        new_size = 64
        while new_size < addr:
            new_size *= 2
        new_data = [0] * new_size
        for i in range(old_size):
            new_data[i] = self.data[i]
        self.data = new_data
        return new_size

    def read(self, addr: int, mode=ADDR_MODE_IMM) -> int:
        addr = self.decode_addr(addr, mode)
        if addr >= len(self.data):
            return 0
        return self.data[addr]

    def write(self, addr: int, v: int, mode=ADDR_MODE_POS) -> None:
        # addr = self.decode_addr(addr, mode)
        self.grow_mem(addr)
        self.data[addr] = v

    def op_add(self, fn: Instruction) -> None:
        """
        Opcode 1 adds together numbers read from two positions and stores the result in a third
        position. The three integers immediately after the opcode tell you these three positions
         - the first two indicate the positions from which you should read the input values, and the
        third indicates the position at which the output should be stored.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        addr = self.decode_addr(self.ptr + 3, fn.get_param(2))
        v = a + b
        self.write(addr, v, fn.get_param(2))
        logger.debug(f"[{self.ptr:04}] ADD({a}, {b}) -> {v}")
        self.ptr += 4

    def op_multiply(self, fn: Instruction) -> None:
        """
        Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding
        them. Again, the three integers after the opcode indicate where the inputs and outputs are,
        not their values.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        addr = self.decode_addr(self.ptr + 3, fn.get_param(2))
        v = a * b
        self.write(addr, v,  fn.get_param(2))
        logger.debug(f"[{self.ptr:04}] MUL({a}, {b}) -> {v})")
        self.ptr += 4

    def op_input(self, fn: Instruction) -> None:
        if self.stdin:
            v = self.stdin[0]
            self.stdin = self.stdin[1:]
        else:
            v = int(input(f"[{self.ptr:04}] Enter input: "))
        addr = self.decode_addr(self.ptr + 1, fn.get_param(0))
        self.write(addr, v)
        logger.debug(f"[{self.ptr:04}] INPUT({addr}) -> {v}")
        self.ptr += 2

    def op_output(self, fn: Instruction) -> None:
        v = self.read(self.ptr + 1, fn.get_param(0))
        self.stdout.append(v)
        logger.debug(f"[{self.ptr:04}] OUTPUT() -> {v}")
        self.ptr += 2

    def op_jcc(self, fn: Instruction) -> None:
        """
        Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction
        pointer to the value from the second parameter. Otherwise, it does nothing.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        result = a != 0
        logger.debug(f"[{self.ptr:04}] JCC({a}, {b}) -> {result}")
        if result:
            self.ptr = b
        else:
            self.ptr += 3

    def op_jnc(self, fn: Instruction) -> None:
        """
        Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer
        to the value from the second parameter. Otherwise, it does nothing.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        result = a == 0
        logger.debug(f"[{self.ptr:04}] JNC({a}, {b}) -> {result}")
        if result:
            self.ptr = b
        else:
            self.ptr += 3

    def op_lt(self, fn: Instruction) -> None:
        """
        Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1
        in the position given by the third parameter. Otherwise, it stores 0.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        addr = self.decode_addr(self.ptr + 3, fn.get_param(2))
        v = int(a < b)
        self.write(addr, v)
        logger.debug(f"[{self.ptr:04}] LT({a}, {b}, {addr}) -> {v}")
        self.ptr += 4

    def op_eq(self, fn: Instruction) -> None:
        """
        Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in
        the position given by the third parameter. Otherwise, it stores 0.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        addr = self.decode_addr(self.ptr + 3, fn.get_param(2))
        v = int(a == b)
        self.write(addr, v)
        logger.debug(f"[{self.ptr:04}] EQ({a}, {b}, {addr}) -> {v}")
        self.ptr += 4

    def op_relbase(self, fn: Instruction) -> None:
        """
        Opcode 9 adjusts the relative base by the value of its only parameter. The relative base
        increases (or decreases, if the value is negative) by the value of the parameter.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        self.reg_rel_base += a
        logger.debug(f"[{self.ptr:04}] RELBASE({a}) -> {self.reg_rel_base}")
        self.ptr += 2

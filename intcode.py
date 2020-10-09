#!/usr/bin/env python3

from typing import List, Optional, Sequence, Tuple


ADDR_MODE_POS = 0
ADDR_MODE_IMM = 1

OPCODE_ADD = 1
OPCODE_MULTIPLY = 2
OPCODE_INPUT = 3
OPCODE_OUTPUT = 4
OPCODE_JCC = 5
OPCODE_JNC = 6
OPCODE_LT = 7
OPCODE_EQ = 8
OPCODE_HALT = 99


Data = List[int]


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
    def __init__(self, data: Optional[Data] = None, stdin: Optional[Data] = None) -> None:
        self.handlers = {
            OPCODE_ADD: self.op_add,
            OPCODE_MULTIPLY: self.op_multiply,
            OPCODE_INPUT: self.op_input,
            OPCODE_OUTPUT: self.op_output,
            OPCODE_JCC: self.op_jcc,
            OPCODE_JNC: self.op_jnc,
            OPCODE_LT: self.op_lt,
            OPCODE_EQ: self.op_eq,
        }
        self.reset(data, stdin)
    
    @property
    def halted(self) -> bool:
        fn = Instruction(self.data[self.ptr])
        return fn.opcode == OPCODE_HALT

    def reset(self, data: Optional[Data], stdin: Optional[Data] = None) -> None:
        self.ptr = 0
        self.data = data.copy() if data else []
        self.stdin: Data = stdin.copy() if stdin else []
        self.stdout: Data = []

    def step(self) -> bool:
        fn = Instruction(self.data[self.ptr])
        if fn.opcode == OPCODE_HALT:
            print(f"[{self.ptr:04}] HALT")
            return False
        if fn.opcode not in self.handlers:
            raise RuntimeError(f"Unrecognized opcode: {fn.opcode}")
        self.handlers[fn.opcode](fn)
        return True

    def run(self) -> None:
        do_step = True
        while do_step:
            do_step = self.step()

    def io_push(self, v: int) -> None:
        self.stdin.append(v)
    
    def io_pop(self) -> int:
        v = self.stdout[0]
        self.stdout = self.stdout[1:]
        return v

    def read(self, addr: int, mode=ADDR_MODE_IMM) -> int:
        v = self.data[addr]
        if mode == ADDR_MODE_IMM:
            return v
        if mode == ADDR_MODE_POS:
            return self.data[v]
        raise RuntimeError(f"Unrecognized address mode: {mode}")

    def op_add(self, fn: Instruction) -> None:
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        c = self.read(self.ptr + 3)
        self.data[c] = a + b
        print(f"[{self.ptr:04}] ADD({a}, {b}) -> {self.data[c]} (@ {c})")
        self.ptr += 4

    def op_multiply(self, fn: Instruction) -> None:
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        c = self.read(self.ptr + 3)
        self.data[c] = a * b
        print(f"[{self.ptr:04}] MUL({a}, {b}) -> {self.data[c]} (@ {c})")
        self.ptr += 4

    def op_input(self, fn: Instruction) -> None:
        addr = self.read(self.ptr + 1)
        if self.stdin:
            v = self.stdin[0]
            self.stdin = self.stdin[1:]
        else:
            v = int(input(f"[{self.ptr:04}] Enter input: "))
        self.data[addr] = v
        print(f"[{self.ptr:04}] INPUT({addr}) -> {self.data[addr]}")
        self.ptr += 2

    def op_output(self, fn: Instruction) -> None:
        v = self.read(self.ptr + 1, fn.get_param(0))
        self.stdout.append(v)
        print(f"[{self.ptr:04}] OUTPUT() -> {v}")
        self.ptr += 2

    def op_jcc(self, fn: Instruction) -> None:
        """
        Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction
        pointer to the value from the second parameter. Otherwise, it does nothing.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        result = a != 0
        print(f"[{self.ptr:04}] JCC({a}, {b}) -> {result}")
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
        print(f"[{self.ptr:04}] JNC({a}, {b}) -> {result}")
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
        c = self.read(self.ptr + 3)
        self.data[c] = int(a < b)
        print(f"[{self.ptr:04}] LT({a}, {b}, {c}) -> {self.data[c]}")
        self.ptr += 4

    def op_eq(self, fn: Instruction) -> None:
        """
        Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in
        the position given by the third parameter. Otherwise, it stores 0.
        """
        a = self.read(self.ptr + 1, fn.get_param(0))
        b = self.read(self.ptr + 2, fn.get_param(1))
        c = self.read(self.ptr + 3)
        self.data[c] = int(a == b)
        print(f"[{self.ptr:04}] EQ({a}, {b}, {c}) -> {self.data[c]}")
        self.ptr += 4

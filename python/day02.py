#!/usr/bin/env python3
import unittest

from typing import List, Tuple

from common import open_fixture


NOUN = 1
VERB = 2

Data = List[int]


def run(data: Data) -> int:
    for ptr in range(0, len(data), 4):
        opcode = data[ptr]
        if opcode == 99:
            return data[0]
        if opcode == 1:
            (a, b, out) = data[ptr + 1], data[ptr + 2], data[ptr + 3]
            data[out] = data[a] + data[b]
            continue
        if opcode == 2:
            (a, b, out) = data[ptr + 1], data[ptr + 2], data[ptr + 3]
            data[out] = data[a] * data[b]
            continue
        raise RuntimeError(f"Invalid opcode at {ptr}: {opcode}")
    return data[0]


def find_input(data: Data, needle: int) -> int:
    for noun in range(100):
        for verb in range(100):
            data2 = data.copy()
            data2[NOUN] = noun
            data2[VERB] = verb
            if run(data2) == needle:
                return 100 * noun + verb
    raise RuntimeError("Not found")


class TestDay2(unittest.TestCase):
    def test_part1(self):
        with open_fixture("day02") as fp:
            data = [int(opcode) for opcode in fp.read().split(",")]
        data[1] = 12
        data[2] = 2
        self.assertEqual(run(data), 5110675)

    def test_part2(self):
        with open_fixture("day02") as fp:
            data = [int(opcode) for opcode in fp.read().split(",")]
        self.assertEqual(find_input(data, 19690720), 4847)

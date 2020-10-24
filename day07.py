#!/usr/bin/env python3

import unittest

from typing import Iterable, Sequence
from itertools import permutations

from lib.intcode import decode, IntcodeVM, Data


def thrust(data: Data, phase_sequence: Sequence[int]) -> int:
    vms = [IntcodeVM(data, [phase]) for phase in phase_sequence]
    E = vms[len(vms) - 1]
    v = 0
    i = 0
    while not E.halted:
        vm = vms[i]
        vm.io_push(v)
        while not vm.halted:
            vm.step()
            if vm.stdout:
                v = vm.io_pop()
                break
        i = (i + 1) % len(phase_sequence)
    return v


def compute_thrust(
    data: Data, phase_sequence: Iterable[int], loop: bool = False
) -> int:
    max_v = 0
    phases_permutations = permutations(phase_sequence)
    for phase_sequence in phases_permutations:
        v = thrust(data, phase_sequence)
        if v > max_v:
            max_v = v
    return max_v


class TestDay7(unittest.TestCase):
    def test_part1_example1(self):
        data = decode("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
        self.assertEqual(compute_thrust(data, range(5)), 43210)

    def test_part1_example2(self):
        data = decode(
            "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
        )
        self.assertEqual(compute_thrust(data, range(5)), 54321)

    def test_part1_example3(self):
        data = decode(
            "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,"
            "1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
        )
        self.assertEqual(compute_thrust(data, range(5)), 65210)

    def test_part1(self):
        with open("./day07.input", "r") as fp:
            data = decode(fp.readline())
        self.assertEqual(compute_thrust(data, range(5)), 13848)

    def test_part2_example1(self):
        data = decode(
            "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
        )
        self.assertEqual(compute_thrust(data, range(5, 10)), 139629729)

    def test_part2_example2(self):
        data = decode(
            "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,"
            "-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,"
            "53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
        )
        self.assertEqual(compute_thrust(data, range(5, 10)), 18216)

    def test_part2(self):
        with open("./day07.input", "r") as fp:
            data = decode(fp.readline())
        self.assertEqual(compute_thrust(data, range(5, 10)), 12932154)

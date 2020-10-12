#!/usr/bin/env python3

from typing import List
import unittest

from intcode import decode, Data, IntcodeVM


def load_vm(stdin: Data) -> Data:
    with open("./day5.input", "r") as fp:
        s = fp.readline()
    data = decode(s)
    vm = IntcodeVM(data, stdin)
    vm.run()
    return vm.stdout


class TestIntcodeVM(unittest.TestCase):
    def test_part1(self):
        stdout = load_vm([1])
        self.assertEqual(stdout[len(stdout) - 1], 7988899)
    
    def test_part2(self):
        stdout = load_vm([5])
        self.assertEqual(stdout[len(stdout) - 1], 13758663)

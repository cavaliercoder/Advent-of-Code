#!/usr/bin/env python3

from typing import List
import unittest

from intcode import decode, Data, IntcodeVM, run, load_and_run


class TestIntcodeVM(unittest.TestCase):
    def test_part1(self):
        stdout = load_and_run("./day5.input", stdin=[1])
        self.assertEqual(stdout[len(stdout) - 1], 7988899)

    def test_part2(self):
        stdout = load_and_run("./day5.input", stdin=[5])
        self.assertEqual(stdout[len(stdout) - 1], 13758663)

    def test_part2_example1(self):
        data = decode("3,9,8,9,10,9,4,9,99,-1,8")
        for i in range(10):
            stdout = run(data, [i])
            self.assertListEqual(stdout, [1 if i == 8 else 0])

    def test_part2_example2(self):
        data = decode("3,9,7,9,10,9,4,9,99,-1,8")
        for i in range(10):
            stdout = run(data, [i])
            self.assertListEqual(stdout, [1 if i < 8 else 0])

    def test_part2_example3(self):
        data = decode("3,3,1108,-1,8,3,4,3,99")
        for i in range(10):
            stdout = run(data, [i])
            self.assertListEqual(
                stdout, [1 if i == 8 else 0], f"Failed with input: {i}"
            )

    def test_part2_example4(self):
        data = decode("3,3,1107,-1,8,3,4,3,99")
        for i in range(10):
            stdout = run(data, [i])
            self.assertListEqual(stdout, [1 if i < 8 else 0])
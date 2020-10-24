#!/usr/bin/env python3

from typing import List
import unittest

from lib.intcode import decode, Data, IntcodeVM, run, load_and_run


class TestIntcodeVM(unittest.TestCase):
    def test_part1(self):
        stdout = load_and_run("./day05.input", stdin=[1])
        self.assertEqual(stdout[len(stdout) - 1], 7988899)

    def test_part2(self):
        stdout = load_and_run("./day05.input", stdin=[5])
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

    def test_part2_example5(self):
        data = decode("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
        for i in range(10):
            stdout = run(data, [i])
            self.assertListEqual(stdout, [0 if i == 0 else 1])

    def test_part2_example6(self):
        data = decode("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
        for i in range(10):
            stdout = run(data, [i])
            self.assertListEqual(stdout, [0 if i == 0 else 1])

    def test_part2_example7(self):
        data = decode(
            "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,"
            "4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
        )
        for i in range(10):
            expect = 999 if i < 8 else 1000 if i == 8 else 1001
            stdout = run(data, [i])
            self.assertListEqual(stdout, [expect])

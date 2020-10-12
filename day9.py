import unittest

from intcode import IntcodeVM, Data, decode


class TestDay9(unittest.TestCase):
    def test_part1_example1(self):
        data = decode("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
        vm = IntcodeVM(data)
        vm.run()
        self.assertListEqual(vm.stdout, data)

    def test_part1_example2(self):
        data = decode("1102,34915192,34915192,7,4,7,99,0")
        vm = IntcodeVM(data)
        vm.run()
        self.assertGreaterEqual(vm.stdout[0], 1000000000000000)

    def test_part1_example3(self):
        data = decode("104,1125899906842624,99")
        vm = IntcodeVM(data)
        vm.run()
        self.assertEqual(vm.stdout[0], 1125899906842624)

    def test_part1(self):
        with open("./day9.input", "r") as fp:
            data = decode(fp.readline())
        vm = IntcodeVM(data, [1])
        vm.run()
        self.assertListEqual(vm.stdout, [3335138414])

    def test_part2(self):
        with open("./day9.input", "r") as fp:
            data = decode(fp.readline())
        vm = IntcodeVM(data, [2])
        vm.run()
        self.assertListEqual(vm.stdout, [49122])

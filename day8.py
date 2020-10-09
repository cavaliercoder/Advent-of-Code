#!/usr/bin/env python3

import unittest


def part1(image: str, width: int, height: int) -> int:
    layer_size = width * height
    layers = []
    zeros = 0
    ones = 0
    twos = 0
    for i in range(len(image)):
        if i and i % layer_size == 0:
            layers.append((zeros, ones, twos))
            zeros = 0
            ones = 0
            twos = 0
        if image[i] == "0":
            zeros += 1
        if image[i] == "1":
            ones += 1
        if image[i] == "2":
            twos += 1
    
    layers = sorted(layers, key=lambda x: x[0])
    return layers[0][1] * layers[0][2]


def part2(image: str, width: int, height: int) -> None:
    charmap = {"0": " ", "1": "█", "2": "2"}
    layer_size = width * height
    buf = ["2"] * layer_size
    for i in range(len(image)):
        j = i % layer_size
        if buf[j] == "2":
            buf[j] = charmap[image[i]]
    for i in range(0, len(buf), width):
        print("".join(buf[i:i+width]))


class TestDay8(unittest.TestCase):
    def test_part1(self):
        with open("day8.input", "r") as fp:
            image = fp.readline()
        self.assertEqual(part1(image, 25, 6), 1560)
    
    def test_part2(self):
        with open("day8.input", "r") as fp:
            image = fp.readline()
        print()
        part2(image, 25, 6)



if __name__ == "__main__":
    unittest.main()

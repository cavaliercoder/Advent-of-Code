#!/usr/bin/env python3

import unittest


def part1(image: str, width: int, height: int) -> int:
    layer_size = width * height
    layers = []
    tally = [0, 0, 0]
    for i in range(len(image)):
        if i and i % layer_size == 0:
            layers.append(tally.copy())
            tally = [0, 0, 0]
        tally[int(image[i])] += 1
    layers = sorted(layers, key=lambda x: x[0])
    return layers[0][1] * layers[0][2]


def part2(image: str, width: int, height: int) -> None:
    charmap = [" ", "█", "?"]
    layer_size = width * height
    buf = ["?"] * layer_size
    for i in range(len(image)):
        j = i % layer_size
        if buf[j] == "?":
            buf[j] = charmap[int(image[i])]
    for i in range(0, len(buf), width):
        print("".join(buf[i:i+width]))


class TestDay8(unittest.TestCase):
    def test_part1(self):
        with open("day8.input", "r") as fp:
            image = fp.readline().rstrip()
        self.assertEqual(part1(image, 25, 6), 1560)
    
    def test_part2(self):
        with open("day8.input", "r") as fp:
            image = fp.readline().rstrip()
        print()
        part2(image, 25, 6)

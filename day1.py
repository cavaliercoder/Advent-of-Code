#!/usr/bin/env python3

from math import floor
from typing import Sequence
import unittest


def get_fuel_for_module(mass: float, recursive: bool = False) -> float:
    fuel = floor(mass / 3) - 2
    if recursive:
        if fuel <= 0:
            return 0
        return fuel + get_fuel_for_module(fuel)
    return fuel


def get_fuel_for_modules(masses: Sequence[int], recursive: bool = False) -> int:
    return sum([
        get_fuel_for_module(module_mass, recursive=recursive)
        for module_mass in masses
    ])


class TestDay1(unittest.TestCase):
    def test_part1(self):
        with open("./day1.input", "r") as fp:
            masses = [float(line.rstrip()) for line in fp.readlines()]
        fuel_required = get_fuel_for_modules(masses)
        self.assertEqual(fuel_required, 3235550)

    def test_part2(self):
        with open("./day1.input", "r") as fp:
            masses = [float(line.rstrip()) for line in fp.readlines()]
        fuel_required = get_fuel_for_modules(masses, recursive=True)
        self.assertEqual(fuel_required, 4850462)


if __name__ == '__main__':
    unittest.main()

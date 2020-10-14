#!/usr/bin/env python3

from math import floor
from typing import Sequence
import unittest


def get_fuel_for_module(mass: float, recursive: bool = False) -> float:
    fuel = floor(mass / 3) - 2
    if recursive:
        if fuel <= 0:
            return 0
        return fuel + get_fuel_for_module(fuel, recursive=recursive)
    return fuel


def get_fuel_for_modules(masses: Sequence[int], recursive: bool = False) -> float:
    return sum(
        [
            get_fuel_for_module(module_mass, recursive=recursive)
            for module_mass in masses
        ]
    )


class TestDay1(unittest.TestCase):
    def test_part1(self):
        with open("./day01.input", "r") as fp:
            masses = [float(line.rstrip()) for line in fp.readlines()]
        fuel_required = get_fuel_for_modules(masses)
        self.assertEqual(fuel_required, 3235550)

    def test_part2_example1(self):
        self.assertEqual(get_fuel_for_module(14, recursive=True), 2)
        self.assertEqual(get_fuel_for_module(1969, recursive=True), 966)
        self.assertEqual(get_fuel_for_module(100756, recursive=True), 50346)

    def test_part2(self):
        with open("./day01.input", "r") as fp:
            masses = [float(line.rstrip()) for line in fp.readlines()]
        fuel_required = get_fuel_for_modules(masses, recursive=True)
        self.assertEqual(fuel_required, 4850462)

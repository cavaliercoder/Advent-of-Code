#!/usr/bin/env python3

from typing import List

import unittest


class Password:
    def __init__(self, value: int) -> None:
        v = value
        m = 1
        data: List[int] = []
        while v:
            m *= 10
            n = v % m
            v -= n
            data.append(int(n / (m / 10)))
        data.reverse()
        self.data = data

    def __str__(self) -> str:
        return str(self.to_int())

    def to_int(self) -> int:
        v = 0
        for i in range(len(self.data)):
            v *= 10
            v += self.data[i]
        return v

    @property
    def is_valid(self) -> bool:
        # It is a six-digit number.
        if len(self.data) != 6:
            return False

        # Going from left to right, the digits never decrease
        for i in range(1, len(self.data)):
            if self.data[i] < self.data[i - 1]:
                return False

        # Two adjacent digits are the same
        for i in range(1, len(self.data)):
            if self.data[i] == self.data[i - 1]:
                return True
        return False

    @property
    def is_valid2(self) -> bool:
        if not self.is_valid:
            return False
        
        # the two adjacent matching digits are not part
        # of a larger group of matching digits
        track = self.data[0]
        dups = 0
        for i in range(1, len(self.data)):
            if self.data[i] == track:
                dups += 1
            else:
                if dups == 1:
                    return True
                dups = 0
                track = self.data[i]
        if dups == 1:
            return True
        return False

    def _increment(self, i: int) -> None:
        v = self.data[i] + 1
        if v > 9:
            self._increment(i - 1)
            return
        for n in range(i, len(self.data)):
            self.data[n] = v

    def increment(self) -> None:
        self._increment(len(self.data) - 1)


def count_passwords(start: int, end: int, part2: bool = False) -> int:
    n = 0
    pwd = Password(start)
    while True:
        if pwd.to_int() > end:
            return n
        if not part2:
            if pwd.is_valid:
                n += 1
        else:
            if pwd.is_valid2:
                n += 1
        pwd.increment()


class TestDay4(unittest.TestCase):
    def test_day4_part1(self):
        self.assertEqual(count_passwords(256310, 732736), 979)

    def test_day4_part2(self):
        self.assertEqual(count_passwords(256310, 732736, part2=True), 635)


if __name__ == "__main__":
    unittest.main()
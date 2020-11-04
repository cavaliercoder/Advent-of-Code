import unittest

from typing import Generator, List

from common import open_fixture


BASE_PATTERN = (0, 1, 0, -1)


def decode(s: str) -> List[int]:
    return [int(c) for c in s.strip()]


def pattern(position: int) -> Generator[int, None, None]:
    skip = True
    while True:
        for i in range(len(BASE_PATTERN)):
            for _ in range(position):
                if skip:
                    skip = False
                    continue
                yield BASE_PATTERN[i]


def fft(signal: List[int]) -> None:
    output = [0] * len(signal)
    for i in range(len(signal)):
        n = 0
        gen = pattern(i + 1)
        for j in range(len(signal)):
            n += signal[j] * next(gen)
        output[i] = abs(n) % 10
    for i in range(len(signal)):
        signal[i] = output[i]


class TestDay16(unittest.TestCase):
    def test_part1_example1(self):
        A = decode("12345678")
        tests = [
            decode("48226158"),
            decode("34040438"),
            decode("03415518"),
            decode("01029498"),
        ]
        for test in tests:
            fft(A)
            self.assertListEqual(A, test)

    def test_part1_example2(self):
        tests = [
            (decode("80871224585914546619083218645595"), decode("24176176")),
            (decode("19617804207202209144916044189917"), decode("73745418")),
            (decode("69317163492948606335995924319873"), decode("52432133")),
        ]
        for A, expect in tests:
            for _ in range(100):
                fft(A)
            self.assertListEqual(A[:8], expect)

    def test_part1(self):
        with open_fixture("day16") as fp:
            A = decode(fp.readline())
        for _ in range(100):
            fft(A)
        self.assertListEqual(A[:8], decode("96136976"))

    def test_part2(self):
        # TODO: exploit tge predictable pattern in the last half of the algorithm
        # Answer: 85,600,369
        # https://www.reddit.com/r/adventofcode/comments/ebf5cy/2019_day_16_part_2_understanding_how_to_come_up/fb4bvw4/
        pass

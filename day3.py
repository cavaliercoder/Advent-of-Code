#!/usr/bin/env python3

from dataclasses import dataclass
from typing import Optional, Sequence

import unittest


def abs(n: int) -> int:
    if n < 0:
        return -n
    return n


def in_plane(a: int, b: int, n: int) -> bool:
    if n < a and n < b:
        return False
    if n > a and n > b:
        return False
    return True


@dataclass
class Point:
    x: int
    y: int
    score: int = 0

    def __str__(self) -> str:
        return f"({self.x}, {self.y})"

    def manhattan_distance(self) -> int:
        return abs(self.x) + abs(self.y)


@dataclass
class Line:
    a: Point
    b: Point
    score: int = 0

    def __str__(self) -> str:
        return f"[{self.a}, {self.b}]"

    @property
    def is_horizontal(self) -> bool:
        return self.a.y == self.b.y

    @property
    def is_vertical(self) -> bool:
        return self.a.x == self.b.x

    @property
    def length(self) -> int:
        return abs((self.a.x - self.b.x) + (self.a.y - self.b.y))

    def intersection(self, other: "Line") -> Optional[Point]:
        if self.is_horizontal:
            if (
                other.is_vertical
                and in_plane(self.a.x, self.b.x, other.a.x)
                and in_plane(other.a.y, other.b.y, self.a.y)
            ):
                return Point(other.a.x, self.a.y)
        else:
            if (
                other.is_horizontal
                and in_plane(self.a.y, self.b.y, other.a.y)
                and in_plane(other.a.x, other.b.x, self.a.x)
            ):
                return Point(self.a.x, other.a.y)
        return None


Path = Sequence[Line]


def decode_path(line: str) -> Path:
    tokens = line.split(",")
    x, y, d = 0, 0, 0
    path: Path = []
    for token in tokens:
        (o, n) = token[0], int(token[1:])
        start = Point(x, y)
        if o == "R":
            x += n
        elif o == "D":
            y -= n
        elif o == "L":
            x -= n
        elif o == "U":
            y += n
        else:
            raise RuntimeError(f"Unknown token: {token}")
        line = Line(start, Point(x, y), d)
        path.append(line)
        d += line.length
    return path


def get_intersections(A: Path, B: Path) -> Sequence[Point]:
    intersections = []
    for line_a in A:
        for line_b in B:
            intersection = line_a.intersection(line_b)
            if intersection:
                # compute score
                intersection.score = (
                    line_a.score + Line(line_a.a, intersection).length
                    + line_b.score + Line(line_b.a, intersection).length
                )
                intersections.append(intersection)
    return intersections


def get_closest_by_manhattan(A: Path, B: Path) -> int:
    results = sorted(point.manhattan_distance() for point in get_intersections(A, B))
    for result in results:
        if result:
            return result
    return results[0]


def get_closest_by_distance(A: Path, B: Path) -> int:
    results = sorted(point.score for point in get_intersections(A, B))
    for result in results:
        if result:
            return result
    return results[0]


class TestDay3(unittest.TestCase):
    def test_part1_fixture1(self):
        A = decode_path("R75,D30,R83,U83,L12,D49,R71,U7,L72")
        B = decode_path("U62,R66,U55,R34,D71,R55,D58,R83")
        self.assertEqual(get_closest_by_manhattan(A, B), 159)

    def test_part1_fixture2(self):
        A = decode_path("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
        B = decode_path("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
        self.assertEqual(get_closest_by_manhattan(A, B), 135)

    def test_part1(self):
        with open("./day3.input", "r") as fp:
            lines = fp.readlines()
        A = decode_path(lines[0])
        B = decode_path(lines[1])
        self.assertEqual(get_closest_by_manhattan(A, B), 2129)

    def test_part2_fixture1(self):
        A = decode_path("R75,D30,R83,U83,L12,D49,R71,U7,L72")
        B = decode_path("U62,R66,U55,R34,D71,R55,D58,R83")
        self.assertEqual(get_closest_by_distance(A, B), 610)

    def test_part2_fixture2(self):
        A = decode_path("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
        B = decode_path("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
        self.assertEqual(get_closest_by_distance(A, B), 410)
    
    def test_part2(self):
        with open("./day3.input", "r") as fp:
            lines = fp.readlines()
        A = decode_path(lines[0])
        B = decode_path(lines[1])
        self.assertEqual(get_closest_by_distance(A, B), 134662)


if __name__ == "__main__":
    unittest.main()

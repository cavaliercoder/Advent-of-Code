#!/usr/bin/env python3

from dataclasses import dataclass
from typing import Dict, List, Optional, Sequence

import unittest


def parse_orbits(s: str) -> Dict[str, str]:
    orbits = {}
    for line in s.splitlines():
        (parent, child) = line.rstrip().split(")")
        orbits[child] = parent
    return orbits


def load_orbits(file: str) -> Dict[str, str]:
    with open(file, "r") as fp:
        s = fp.read()
    return parse_orbits(s)


def get_graph(orbits: Dict[str, str]) -> Dict[str, Sequence[str]]:
    graph = {}

    def add_edge(parent: str, child: str) -> None:
        if parent in graph:
            graph[parent].append(child)
        else:
            graph[parent] = [child]

    for child, parent in orbits.items():
        add_edge(parent, child)
        add_edge(child, parent)

    return graph


def count_orbits(orbits: Dict[str, str]) -> int:
    cache = {}

    def count_orbit(body: str) -> int:
        if body == "COM":
            return 0
        if body in cache:
            return cache[body]
        parent = orbits[body]
        v = cache[body] = count_orbit(parent) + 1
        return v

    return sum(count_orbit(body) for body in orbits)


def count_transfers(graph: Dict[str, Sequence[str]], needle = "SAN", start = "YOU") -> int:
    seen = set()
    distances = {start: 0}
    stack = [start]
    while True:
        V = stack.pop()
        if V in seen:
            continue
        if V == needle:
            return distances[V] - 2
        seen.add(V)
        stack.extend(graph[V])
        for edge in graph[V]:
            stack.append(edge)
            distances[edge] = distances[V] + 1


class TestDay6(unittest.TestCase):
    def test_part1(self):
        orbits = load_orbits("./day6.input")
        self.assertEqual(count_orbits(orbits), 312697)

    def test_part2(self):
        result = count_transfers(get_graph(load_orbits("./day6.input")))
        self.assertEqual(result, 466)

    def test_part2_example1(self):
        s = """COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
"""

        result = count_transfers(get_graph(parse_orbits(s)))
        self.assertEqual(result, 4)


if __name__ == "__main__":
    unittest.main()

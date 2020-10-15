import re
import unittest

from dataclasses import dataclass
from hashlib import sha256
from typing import Any, Sequence, Tuple


@dataclass
class Coord:
    PATTERN = re.compile(r"^<x=(-?[0-9]+), y=(-?[0-9]+), z=(-?[0-9]+)>$")

    x: int
    y: int
    z: int

    def __add__(self, other: Any) -> "Coord":
        if not isinstance(other, Coord):
            return NotImplemented
        return Coord(
            x=self.x + other.x,
            y=self.y + other.y,
            z=self.z + other.z,
        )

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, Coord):
            return NotImplemented
        return self.x == other.x and self.y == other.y and self.z == other.z

    def __str__(self) -> str:
        return f"<x={self.x}, y={self.y}, z={self.z}>"

    def copy(self) -> "Coord":
        return Coord(
            x=self.x,
            y=self.y,
            z=self.z,
        )

    def energy(self) -> int:
        return abs(self.x) + abs(self.y) + abs(self.z)

    @classmethod
    def from_str(cls, s: str) -> "Coord":
        match = cls.PATTERN.match(s)
        if not match:
            raise ValueError(f"Not a valid Coord: '{s}'")
        x, y, z = [int(v) for v in match.groups()]
        return cls(x=x, y=y, z=z)


ZERO = Coord(0, 0, 0)


class Moon:
    def __init__(self, position: Coord, velocity: Coord = ZERO) -> None:
        self.position = position.copy()
        self.velocity = velocity.copy()

    def __eq__(self, other: Any) -> bool:
        if not isinstance(other, Moon):
            return NotImplemented
        return self.position == other.position and self.velocity == other.velocity

    def __str__(self) -> str:
        return f"pos={self.position}, vel={self.velocity}"

    def energy(self) -> int:
        return self.position.energy() * self.velocity.energy()


def velocity_delta(a: int, b: int) -> Tuple[int, int]:
    """
    Compute the velocity delta of a single step between two bodies, along a single axis.
    """
    if a < b:
        return (1, -1)
    if a > b:
        return (-1, 1)
    return (0, 0)


def apply_gravity(a: Moon, b: Moon) -> None:
    """
    Adjust the velocity of two moons according to their position along each axis.
    """
    # x axis
    delta = velocity_delta(a.position.x, b.position.x)
    a.velocity.x += delta[0]
    b.velocity.x += delta[1]

    # y axis
    delta = velocity_delta(a.position.y, b.position.y)
    a.velocity.y += delta[0]
    b.velocity.y += delta[1]

    # z axis
    delta = velocity_delta(a.position.z, b.position.z)
    a.velocity.z += delta[0]
    b.velocity.z += delta[1]


def step(moons: Sequence[Moon]) -> None:
    # apply gravity and update velocities
    for i in range(len(moons)):
        for j in range(i, len(moons)):
            apply_gravity(moons[i], moons[j])

    # apply velocity and update positions
    for moon in moons:
        moon.position = moon.position + moon.velocity


class TestDay12(unittest.TestCase):
    def assertMoons(self, actual: Sequence[Moon], expect: Sequence[Moon]):
        for i in range(len(expect)):
            if actual[i] == expect[i]:
                continue
            raise AssertionError(f"   {actual[i]}\n" f"!= {expect[i]}")

    def test_part1_example1(self):
        moons = [
            Moon(position=Coord.from_str(s))
            for s in [
                "<x=-1, y=0, z=2>",
                "<x=2, y=-10, z=-7>",
                "<x=4, y=-8, z=8>",
                "<x=3, y=5, z=-1>",
            ]
        ]

        # step 1
        step(moons)
        self.assertMoons(
            moons,
            [
                Moon(position=Coord(x=2, y=-1, z=1), velocity=Coord(x=3, y=-1, z=-1)),
                Moon(position=Coord(x=3, y=-7, z=-4), velocity=Coord(x=1, y=3, z=3)),
                Moon(position=Coord(x=1, y=-7, z=5), velocity=Coord(x=-3, y=1, z=-3)),
                Moon(position=Coord(x=2, y=2, z=0), velocity=Coord(x=-1, y=-3, z=1)),
            ],
        )

        # step 10
        for _ in range(9):
            step(moons)
        self.assertMoons(
            moons,
            [
                Moon(position=Coord(x=2, y=1, z=-3), velocity=Coord(x=-3, y=-2, z=1)),
                Moon(position=Coord(x=1, y=-8, z=0), velocity=Coord(x=-1, y=1, z=3)),
                Moon(position=Coord(x=3, y=-6, z=1), velocity=Coord(x=3, y=2, z=-3)),
                Moon(position=Coord(x=2, y=0, z=4), velocity=Coord(x=1, y=-1, z=-1)),
            ],
        )

    def get_moons(self) -> Sequence[Moon]:
        moons = []
        with open("./day12.input", "r") as fp:
            while True:
                s = fp.readline()
                if not s:
                    break
                moons.append(Moon(position=Coord.from_str(s.strip())))
        return moons

    def test_part1(self):
        moons = self.get_moons()
        for _ in range(1000):
            step(moons)
        energy = sum([moon.energy() for moon in moons])
        self.assertEqual(energy, 12351)

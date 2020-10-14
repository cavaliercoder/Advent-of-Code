import unittest

from enum import Enum
from intcode import IntcodeVM, Data, decode
from typing import Dict, Tuple


Position = Tuple[int, int]


class Color(Enum):
    BLACK = 0
    WHITE = 1


class Orientation(Enum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3

    def turn_left(self) -> "Orientation":
        return Orientation((4 + self.value - 1) % 4)

    def turn_right(self) -> "Orientation":
        return Orientation((4 + self.value + 1) % 4)


class Grid:
    def __init__(self) -> None:
        self.data: Dict[int, Dict[int, Color]] = {}
        self.min_x = 0
        self.max_x = 0
        self.min_y = 0
        self.max_y = 0

    def __getitem__(self, position: Position) -> Color:
        x, y = position
        if x in self.data:
            if y in self.data[x]:
                return self.data[x][y]
        return Color.BLACK

    def __setitem__(self, position: Position, color: Color) -> None:
        x, y = position
        if x in self.data:
            self.data[x][y] = color
        else:
            self.data[x] = {y: color}
        if x < self.min_x:
            self.min_x = x
        if x > self.max_x:
            self.max_x = x
        if y < self.min_y:
            self.min_y = y
        if y > self.max_y:
            self.max_y = y

    def __len__(self) -> int:
        return sum([len(plane) for plane in self.data.values()])

    def __str__(self) -> str:
        charmap = {
            Color.BLACK: " ",
            Color.WHITE: "█",
        }
        lines = []
        for y in range(self.height):
            line = [
                charmap[self[(self.min_x + x, self.min_y + y)]]
                for x in range(self.width)
            ]
            lines.append("".join(line) + "\n")
        return "".join(lines)

    @property
    def width(self) -> int:
        return self.max_x - self.min_x + 1

    @property
    def height(self) -> int:
        return self.max_y - self.min_y + 1


class Robot:
    def __init__(self, program: Data) -> None:
        self.vm = IntcodeVM(program)
        self.grid = Grid()
        self.orientation = Orientation.UP
        self.position = (0, 0)

    def step(self) -> bool:
        # get current color
        self.vm.io_push(self.grid[self.position].value)

        # run the program until output is availble
        while len(self.vm.stdout) < 2:
            if not self.vm.step():
                return False

        # paint
        color = Color(self.vm.io_pop())
        self.grid[self.position] = color

        # turn
        direction = self.vm.io_pop()
        if direction:
            self.orientation = self.orientation.turn_right()
        else:
            self.orientation = self.orientation.turn_left()

        # move
        if self.orientation == Orientation.UP:
            self.position = (self.position[0], self.position[1] - 1)
        if self.orientation == Orientation.RIGHT:
            self.position = (self.position[0] + 1, self.position[1])
        if self.orientation == Orientation.DOWN:
            self.position = (self.position[0], self.position[1] + 1)
        if self.orientation == Orientation.LEFT:
            self.position = (self.position[0] - 1, self.position[1])

        return not self.vm.halted

    def run(self) -> None:
        while self.step():
            pass


class TestDay11(unittest.TestCase):
    def test_part1(self):
        with open("./day11.input", "r") as fp:
            data = decode(fp.readline())
        robot = Robot(data)
        robot.run()
        self.assertEqual(len(robot.grid), 1747)

    def test_part2(self):
        expect = (
            " ████  ██   ██  ███  █  █ █  █ █    ███    \n"
            "    █ █  █ █  █ █  █ █  █ █ █  █    █  █   \n"
            "   █  █    █    █  █ ████ ██   █    ███    \n"
            "  █   █    █ ██ ███  █  █ █ █  █    █  █   \n"
            " █    █  █ █  █ █ █  █  █ █ █  █    █  █   \n"
            " ████  ██   ███ █  █ █  █ █  █ ████ ███    \n"
        )
        with open("./day11.input", "r") as fp:
            data = decode(fp.readline())
        robot = Robot(data)
        robot.grid[(0, 0)] = Color.WHITE
        robot.run()
        actual = str(robot.grid)
        self.assertEqual(actual, expect)

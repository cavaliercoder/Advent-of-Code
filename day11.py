import unittest

from enum import Enum
from intcode import IntcodeVM, Data, decode
from typing import Dict, Tuple


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


class Robot:
    def __init__(self, program: Data) -> None:
        self.vm = IntcodeVM(program)
        self.grid = {0: {0: Color.BLACK}}
        self.orientation = Orientation.UP
        self.position = (0, 0)
        self.color = Color.BLACK

    def get_color(self, position: Tuple[int, int]) -> Color:
        x, y = position
        if x in self.grid:
            if y in self.grid[x]:
                return self.grid[x][y]
        return Color.BLACK

    def paint(self, color: Color) -> None:
        x, y = self.position
        if x in self.grid:
            self.grid[x][y] = color
        else:
            self.grid[x] = {y: color}
        self.color = color

    def step(self) -> bool:
        # get current color
        self.vm.io_push(self.get_color(self.position).value)

        # run the program until output is availble
        while len(self.vm.stdout) < 2:
            if not self.vm.step():
                # print("  -> End")
                return False

        # paint
        color = Color(self.vm.io_pop())
        self.paint(color)

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

    def render(self) -> str:
        # find bounds
        min_x = 0
        min_y = 0
        max_x = 0
        max_y = 0
        for x in self.grid:
            if x < min_x:
                min_x = x
            if x > max_x:
                max_x = x
            for y in self.grid[x]:
                if y < min_y:
                    min_y = y
                if y > max_y:
                    max_y = y
        w = max_x - min_x + 1
        h = max_y - min_y + 1

        # print
        charmap = {
            Color.BLACK: " ",
            Color.WHITE: "█",
        }
        lines = []
        for y in range(h):
            line = [charmap[self.get_color((min_x + x, min_y + y))] for x in range(w)]
            lines.append("".join(line) + "\n")
        return "".join(lines)


class TestDay11(unittest.TestCase):
    def test_part1(self):
        with open("./day11.input", "r") as fp:
            data = decode(fp.readline())
        robot = Robot(data)
        robot.run()
        n = sum([len(plane) for plane in robot.grid.values()])
        self.assertEqual(n, 1747)

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
        robot.paint(Color.WHITE)
        robot.run()
        actual = robot.render()
        self.assertEqual(actual, expect)

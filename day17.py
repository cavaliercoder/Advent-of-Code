import logging
import sys
import unittest

from io import StringIO
from time import sleep
from typing import Any, IO, List, Iterator, Tuple

from lib.curses import CursesScreen
from lib.intcode import IntcodeVM, load


GRID_WIDTH = 45
GRID_HEIGHT = 41

# Out of laziness, I figured out the solution by hand.
# If I were going to do it programmatically, then maybe:
#  - parse the grid to a graph data structure
#  - compute all possible paths through the graph
#  - Computing rolling hash of all subpaths
#  - Look for repeating hashes
# Way too much code for today, so...
SOLUTION = [
    "A,B,A,C,A,B,A,C,B,C",  # main()
    "R,4,L,12,L,8,R,4",  # A()
    "L,8,R,10,R,10,R,6",  # B()
    "R,4,R,10,L,12",  # C()
]


class Ascii:
    def __init__(
        self,
        wake_up: bool = False,
    ):
        self.vm = IntcodeVM(load("./day17.input"))
        if wake_up:
            self.vm.data[0] = 2

    def putc(self, c: int) -> None:
        """
        Block until the character is written to the vm input buffer.
        """
        assert not self.vm.stdin
        assert not self.vm.stdout
        self.vm.io_push(c)
        while self.vm.step() and self.vm.stdin:
            assert not self.vm.stdout

    def write(self, s: str) -> None:
        """
        Block until the whole string is written to the vm input buffer.
        """
        for c in s:
            self.putc(ord(c))

    def getc(self) -> str:
        """
        Block until a character is read from the vm output buffer.
        """
        assert not self.vm.stdin
        assert not self.vm.stdout
        while self.vm.step():
            if self.vm.stdout:
                return chr(self.vm.io_pop())
        return ""

    def readline(self) -> str:
        """
        Block until a line is read from the vm output buffer.
        """
        with StringIO() as s:
            while True:
                c = self.getc()
                if not c:
                    break
                s.write(c)
                if c == "\n":
                    break
            return s.getvalue()

    def readgrid(self) -> str:
        """
        Block until a whole grid is read from the vm output buffer.
        """
        with StringIO() as s:
            for _ in range((GRID_WIDTH + 1) * GRID_HEIGHT):
                c = self.getc()
                s.write(c)
                if not c:
                    break
            assert not self.vm.stdout
            return s.getvalue()

    def get_dust_count(self) -> int:
        """
        Run the program to completion and return the sum of dust collected.
        """
        while self.vm.step():
            continue
        return self.vm.stdout[len(self.vm.stdout) - 1]


def compute_intersections(
    s: str, width: int = GRID_WIDTH, height: int = GRID_HEIGHT
) -> int:
    """
    Compute the sum of alignment parameters of all intersections in the grid.
    """

    def getc(pos: Tuple[int, int]) -> str:
        x, y = pos
        if x < 0 or x >= width or y < 0 or y >= height:
            return "."
        return s[(y * (width + 1)) + x]

    def is_intersection(pos: Tuple[int, int]) -> bool:
        if getc(pos) != "#":
            return False
        neighbors = [
            getc((pos[0], pos[1] - 1)),
            getc((pos[0] + 1, pos[1])),
            getc((pos[0], pos[1] + 1)),
            getc((pos[0] - 1, pos[1])),
        ]
        return sum([1 for c in neighbors if c == "#"]) > 2

    intersections = [
        (x, y) for x in range(width) for y in range(height) if is_intersection((x, y))
    ]
    return sum([(pos[0] * pos[1]) for pos in intersections])


class TestDay17(unittest.TestCase):
    def test_part1(self):
        s = Ascii().readgrid()
        self.assertEqual(compute_intersections(s), 3936)

    def test_part2(self):
        # print initial grid
        bot = Ascii(wake_up=True)
        logging.debug(bot.readgrid())
        logging.debug(bot.readline())

        # send arguments
        args = SOLUTION + ["n"]
        for i in range(len(args)):
            logging.debug(bot.readline())
            logging.debug(f"> {args[i]}")
            bot.write(args[i] + "\n")

        # print final grid
        logging.debug(bot.readgrid())
        bot.readline()
        bot.readline()

        # compute dust collected
        self.assertEqual(bot.get_dust_count(), 785733)


if __name__ == "__main__":
    bot = Ascii(wake_up=True)
    args = SOLUTION + ["y"]
    with CursesScreen() as stdscr:
        # print initial grid
        stdscr.addstr(0, 0, bot.readgrid())
        bot.readline()
        stdscr.refresh()

        # send arguments
        for i in range(len(args)):
            arg = args[i] + "\n"
            bot.readline()  # discard prompt
            bot.write(arg)

        # play feed
        while not bot.vm.halted:
            bot.readline()  # discard \n
            stdscr.addstr(0, 0, bot.readgrid())
            stdscr.refresh()
        stdscr.addstr(41, 0, "HALTED (Ctrl-C to exit)")
        stdscr.refresh()
        while True:
            sleep(1)

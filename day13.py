import curses
import unittest

from enum import Enum
from typing import Any, Dict, Optional, List, Tuple, IO
from time import sleep
from sys import stdout

from lib.intcode import IntcodeVM, Data, load
from lib.curses import CursesScreen

Position = Tuple[int, int]  # (x, y)


class Tile(Enum):
    EMPTY = 0
    WALL = 1
    BLOCK = 2
    PADDLE = 3
    BALL = 4

    def __str__(self) -> str:
        s = [" ", "█", "X", "=", "⌾"]
        return s[int(self.value)]


class ArcadeCabinet:
    def __init__(self, data: Data) -> None:
        self.vm = IntcodeVM(data, trap_input=self._trap_input)
        self.score = 0
        self.frame = 0
        self.init_block_count = 0
        self.last_good_frame = 0
        self.last_good_state = data.copy()
        self.last_miss_x = 0
        self.last_paddle_x = 0
        self.failed = False
        self.moves: Optional[List[int]] = None
        self.stdscr: Any = None

    def step(self) -> bool:
        if not self.vm.step():
            return False
        if len(self.vm.stdout) == 3:
            self._trap_draw_tile()
        return True

    def _trap_draw_tile(self) -> None:
        x = self.vm.io_pop()
        y = self.vm.io_pop()
        if x == -1 and y == 0:
            # update score
            self.score = self.vm.io_pop()
            return

        # paint a tile
        tile = Tile(self.vm.io_pop())
        if tile == Tile.BLOCK:
            # track count for part 1
            self.init_block_count += 1

        if tile == Tile.PADDLE:
            self.last_paddle_x = x

        if tile == Tile.BALL:
            # did the ball hit the paddle
            if y == 20 and x in [
                self.last_paddle_x - 1,
                self.last_paddle_x,
                self.last_paddle_x + 1,
            ]:
                self.last_good_frame = self.frame
                self.last_good_state = self.vm.data.copy()

            # did the ball go past the paddle
            if y == 21:
                self.failed = True
                self.last_miss_x = x

        # update curses window
        if self.stdscr is None:
            return
        self.stdscr.addstr(y, x, str(tile))

    def _trap_input(self) -> int:
        """
        Called whenever the vm is starved for input, between each iteration of the program state.
        """
        self.frame += 1

        v = 0  # neutral paddle movement
        if self.moves:
            v = self.moves[0]
            self.moves = self.moves[1:]

        stdscr = self.stdscr
        if stdscr is None:
            return v

        # paint curses window
        stdscr.addstr(24, 0, f"Score: {self.score} Frame: {self.frame}")
        stdscr.refresh()
        sleep(0.01)
        return v

    def run(self, stdscr=None, moves: Optional[List[int]] = None) -> bool:
        # if stdscr:
        self.stdscr = stdscr
        self.moves = moves
        while not self.vm.halted:
            self.step()
        return not self.failed


def learn_and_score(data: Data, stdscr: Any = None) -> int:
    """
    Learn to beat the game and return the final score.
    """
    cab = ArcadeCabinet(data)
    cab.vm.data[0] = 2  # insert coin!
    state = cab.vm.data.copy()
    moves: List[int] = []
    while True:
        cab = ArcadeCabinet(state)
        if cab.run(moves=moves, stdscr=stdscr):
            return cab.score

        # compute moves required from last good state
        diff_x = cab.last_paddle_x - cab.last_miss_x
        if diff_x > 0:
            moves = [-1] * diff_x
        else:
            moves = [1] * -diff_x

        # reset to last good state
        state = cab.last_good_state


class TestDay13(unittest.TestCase):
    def test_part1(self):
        cab = ArcadeCabinet(load("./day13.input"))
        cab.run()
        self.assertEqual(cab.init_block_count, 420)

    def test_part2(self):
        score = learn_and_score(load("./day13-frame10588.input"))
        self.assertEqual(score, 21651)


if __name__ == "__main__":
    with CursesScreen() as stdscr:
        score = learn_and_score(load("./day13.input"), stdscr=stdscr)
    print(f"Score: {score}")

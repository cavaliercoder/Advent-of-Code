import curses
import os

from typing import Any, IO


def fixture(s: str) -> str:
    """
    Get the relative path to a puzzle input file.
    """
    return os.path.join("..", "..", "inputs", "2019", s) + ".txt"


def open_fixture(s: str) -> IO[Any]:
    """
    Open a puzzle input file.
    """
    return open(fixture(s), "r")


class CursesScreen:
    """
    Setup and tear down a Cursors window.
    """

    def __enter__(self):
        self.stdscr = curses.initscr()
        curses.noecho()
        curses.cbreak()
        curses.curs_set(False)
        self.stdscr.nodelay(True)
        self.stdscr.keypad(1)
        return self.stdscr

    def __exit__(self, type, value, traceback):
        self.stdscr.keypad(0)
        curses.echo()
        curses.nocbreak()
        curses.endwin()

import curses


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

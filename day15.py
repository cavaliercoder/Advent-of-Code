import unittest

from enum import Enum
from intcode import IntcodeVM, load
from sys import stdout
from typing import Any, Dict, Optional, List, Tuple, IO


class Direction(Enum):
    NORTH = 1
    SOUTH = 2
    WEST = 3
    EAST = 4

    @classmethod
    def all(cls) -> Tuple["Direction", "Direction", "Direction", "Direction"]:
        return (
            Direction.NORTH,
            Direction.SOUTH,
            Direction.WEST,
            Direction.EAST,
        )


class Status(Enum):
    WALL = 0
    OKAY = 1
    OXYGENATED = 2

    @property
    def passable(self) -> bool:
        return self == Status.OKAY or self == Status.OXYGENATED


class Position:
    def __init__(self, x: int, y: int) -> None:
        self.x = x
        self.y = y

    def __repr__(self) -> str:
        return f"({self.x}, {self.y})"

    def __bool__(self) -> bool:
        return self.x != 0 or self.y != 0

    def move(self, direction: Direction) -> "Position":
        if direction == Direction.NORTH:
            return Position(x=self.x, y=self.y + 1)
        if direction == Direction.SOUTH:
            return Position(x=self.x, y=self.y - 1)
        if direction == Direction.WEST:
            return Position(x=self.x - 1, y=self.y)
        if direction == Direction.EAST:
            return Position(x=self.x + 1, y=self.y)
        raise RuntimeError(f"Bad direction: {direction}")

    def neighbors(self) -> Tuple["Position", "Position", "Position", "Position"]:
        return (
            Position(x=self.x, y=self.y + 1),
            Position(x=self.x, y=self.y - 1),
            Position(x=self.x - 1, y=self.y),
            Position(x=self.x + 1, y=self.y),
        )


class Cell:
    def __init__(self, pos: Position, distance: int, status: Status) -> None:
        self.pos = pos
        self.distance = distance
        self.status = status


class Grid:
    def __init__(self) -> None:
        self.cells: Dict[int, Dict[int, Cell]] = {
            0: {0: Cell(pos=Position(0, 0), distance=0, status=Status.OKAY)}
        }
        self.min_x = 0
        self.max_x = 0
        self.min_y = 0
        self.max_y = 0

    def __contains__(self, pos: Position) -> bool:
        return pos.y in self.cells and pos.x in self.cells[pos.y]

    def __getitem__(self, pos: Position) -> Cell:
        if pos.y not in self.cells:
            raise KeyError(str(pos))
        if pos.x not in self.cells[pos.y]:
            raise KeyError(str(pos))
        return self.cells[pos.y][pos.x]

    def __setitem__(self, pos: Position, cell: Cell) -> None:
        # update bounds
        if pos.x < self.min_x:
            self.min_x = pos.x
        if pos.x > self.max_x:
            self.max_x = pos.x
        if pos.y < self.min_y:
            self.min_y = pos.y
        if pos.y > self.max_y:
            self.max_y = pos.y

        # store
        if pos.y in self.cells:
            self.cells[pos.y][pos.x] = cell
            return
        self.cells[pos.y] = {pos.x: cell}

    def distance_to(self, pos: Position) -> int:
        if pos in self:
            # distance is already known
            return self[pos].distance

        # get all passable neighbors
        neighbors = [
            self[neighbor]
            for neighbor in pos.neighbors()
            if neighbor in self and self[neighbor].status.passable
        ]
        if not neighbors:
            raise ValueError(f"No path to {pos} {self.cells}")

        # compute shortest distance
        distance = neighbors[0].distance
        for neighbor in neighbors:
            if neighbor.distance < distance:
                distance = neighbor.distance
        return distance + 1

    def set_status(self, pos: Position, status: Status) -> None:
        if pos in self:
            self[pos].status = status
            return

        # compute distance and store
        distance = self.distance_to(pos)
        self[pos] = Cell(
            pos=pos,
            distance=distance,
            status=status,
        )

    def oxygenate(self, start: Position) -> int:
        t = 0
        stack = [start]
        next_stack = []
        while stack:
            while stack:
                # oxygenate and find unoxygenated neighbors
                pos = stack.pop()
                self.set_status(pos, Status.OXYGENATED)
                neighbors = [
                    neighbor
                    for neighbor in pos.neighbors()
                    if neighbor in self and self[neighbor].status == Status.OKAY
                ]
                next_stack.extend(neighbors)

            # swap stack
            stack.extend(next_stack)
            next_stack = []
            t += 1
        return t - 1

    def print(self, f: IO = stdout) -> None:
        for y in range(self.min_y, self.max_y + 1):
            for x in range(self.min_x, self.max_x + 1):
                pos = Position(x, y)
                if not pos:
                    f.write("^")
                    continue
                if not pos in self:
                    f.write("?")
                    continue
                cell = self[pos]
                if cell.status == Status.WALL:
                    f.write("█")
                    continue
                if cell.status == Status.OKAY:
                    f.write(f" ")
                    continue
                if cell.status == Status.OXYGENATED:
                    f.write("$")
                    continue
            f.write("\n")


class Droid:
    def __init__(self) -> None:
        self.vm = IntcodeVM(load("./day15.input"))
        self.grid = Grid()
        self.pos = Position(0, 0)
        self.oxygen_system_pos = Position(0, 0)

    @property
    def cell(self) -> Cell:
        return self.grid[self.pos]

    def move(self, direction: Direction) -> bool:
        self.vm.io_push(direction.value)
        while not self.vm.stdout:
            self.vm.step()
        status = Status(self.vm.io_pop())
        pos = self.pos.move(direction)
        self.grid.set_status(pos, status)
        if status.passable:
            self.pos = pos
            if status == Status.OXYGENATED:
                self.oxygen_system_pos = pos
            return True
        return False

    def explore(self) -> bool:
        # first explore unexplored cells
        for direction in Direction.all():
            npos = self.pos.move(direction)
            if npos not in self.grid:
                if self.move(direction):
                    return True

        # invariant: all neighbors are known

        # quit if home again
        if not self.pos:
            return False

        # move back towards home
        options = {}
        for direction in Direction.all():
            npos = self.pos.move(direction)
            cell = self.grid[npos]
            if cell.status == Status.OKAY:
                options[cell.distance] = direction
                continue
        assert options
        best = sorted(options.keys())[0]
        assert best < self.cell.distance
        direction = options[best]
        assert self.move(direction)
        return True


class TestDay15(unittest.TestCase):
    def test_part1(self):
        droid = Droid()
        while droid.explore():
            continue
        self.assertEqual(droid.grid[droid.oxygen_system_pos].distance, 262)

    def test_part2(self):
        droid = Droid()
        while droid.explore():
            continue
        self.assertEqual(droid.grid.oxygenate(droid.oxygen_system_pos), 314)

package day10

import (
	"strings"
	"testing"
)

func assertVaporize(t *testing.T, grid *Grid, station Coord, n int, expect Coord) bool {
	actual, err := Part2(grid, station, n)
	if err != nil {
		t.Error(err)
		return false
	}
	if actual == expect {
		return true
	}
	t.Errorf("expected %v, got %v", expect, actual)
	return false
}

func TestDay10_Part2_Example1(t *testing.T) {
	s := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`
	grid := testGrid(t, strings.NewReader(s))
	assertVaporize(t, grid, Coord{11, 13}, 1, Coord{11, 12})
	assertVaporize(t, grid, Coord{11, 13}, 2, Coord{12, 1})
	assertVaporize(t, grid, Coord{11, 13}, 3, Coord{12, 2})
	assertVaporize(t, grid, Coord{11, 13}, 10, Coord{12, 8})
	assertVaporize(t, grid, Coord{11, 13}, 20, Coord{16, 0})
	assertVaporize(t, grid, Coord{11, 13}, 50, Coord{16, 9})
	assertVaporize(t, grid, Coord{11, 13}, 100, Coord{10, 16})
	assertVaporize(t, grid, Coord{11, 13}, 199, Coord{9, 6})
	assertVaporize(t, grid, Coord{11, 13}, 200, Coord{8, 2})
	assertVaporize(t, grid, Coord{11, 13}, 201, Coord{10, 9})
	assertVaporize(t, grid, Coord{11, 13}, 299, Coord{11, 1})
}

func TestDay10_Part2(t *testing.T) {
	grid := testGridFromInput(t)
	station, _ := Part1(grid)
	assertVaporize(t, grid, station, 200, Coord{26, 28})
}

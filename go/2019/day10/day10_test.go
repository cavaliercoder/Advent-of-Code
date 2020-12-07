package day10

import (
	"io"
	"strings"
	"testing"

	. "aoc"
)

func testGrid(t *testing.T, r io.Reader) *Grid {
	grid, err := ReadGrid(r)
	if err != nil {
		t.Fatal(err)
	}
	return grid
}

func testGridFromInput(t *testing.T) *Grid {
	grid, err := OpenGrid(Fixture("day10"))
	if err != nil {
		t.Fatal(err)
	}
	return grid
}

func assertVaporize(t *testing.T, grid *Grid, station Pos, n int, expect Pos) bool {
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

func assertBestStation(t *testing.T, grid *Grid, p Pos, n int) bool {
	actualP, actualN := Part1(grid)
	okP := AssertPos(t, p, actualP, "Bad best station")
	okN := AssertInt(t, n, actualN, "Bad best station")
	ok := okP && okN
	if ok {
		return true
	}
	return false
}

func TestGrid(t *testing.T) {
	grid := testGridFromInput(t)
	AssertInt(t, grid.Width, 41, "Bad grid width")
	AssertInt(t, grid.Height, 41, "Bad grid height")
	AssertGridCell(t, grid, NewPos(0, 0), '.', "Bad grid cell")
	AssertGridCell(t, grid, NewPos(1, 0), '#', "Bad grid cell")
	AssertGridCell(t, grid, NewPos(40, 40), '.', "Bad grid cell")
	for i := 0; i < len(grid.Data); i++ {
		p := grid.Pos(i)
		b := grid.Get(p)
		if b != grid.Data[i] {
			t.Fatalf("expected '%c', got '%c' at i: %d, Pos: %v", grid.Data[i], b, i, p)
		}
	}
}

func TestDay10_Part1_Example1(t *testing.T) {
	s := `.#..#
.....
#####
....#
...##
`
	grid := testGrid(t, strings.NewReader(s))
	assertBestStation(t, grid, NewPos(3, 4), 8)
}

func TestDay10_Part1_Example3(t *testing.T) {
	s := `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`
	grid := testGrid(t, strings.NewReader(s))
	assertBestStation(t, grid, NewPos(5, 8), 33)
}

func TestDay10_Part1_Example4(t *testing.T) {
	s := `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
`
	grid := testGrid(t, strings.NewReader(s))
	assertBestStation(t, grid, NewPos(1, 2), 35)
}

func TestDay10_Part1_Example5(t *testing.T) {
	s := `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
`
	grid := testGrid(t, strings.NewReader(s))
	assertBestStation(t, grid, NewPos(6, 3), 41)
}

func TestDay10_Part1_Example6(t *testing.T) {
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
	assertBestStation(t, grid, NewPos(11, 13), 210)
}

func TestDay10_Part1(t *testing.T) {
	grid := testGridFromInput(t)
	assertBestStation(t, grid, NewPos(28, 29), 340)
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
	assertVaporize(t, grid, NewPos(11, 13), 1, NewPos(11, 12))
	assertVaporize(t, grid, NewPos(11, 13), 2, NewPos(12, 1))
	assertVaporize(t, grid, NewPos(11, 13), 3, NewPos(12, 2))
	assertVaporize(t, grid, NewPos(11, 13), 10, NewPos(12, 8))
	assertVaporize(t, grid, NewPos(11, 13), 20, NewPos(16, 0))
	assertVaporize(t, grid, NewPos(11, 13), 50, NewPos(16, 9))
	assertVaporize(t, grid, NewPos(11, 13), 100, NewPos(10, 16))
	assertVaporize(t, grid, NewPos(11, 13), 199, NewPos(9, 6))
	assertVaporize(t, grid, NewPos(11, 13), 200, NewPos(8, 2))
	assertVaporize(t, grid, NewPos(11, 13), 201, NewPos(10, 9))
	assertVaporize(t, grid, NewPos(11, 13), 299, NewPos(11, 1))
}

func TestDay10_Part2(t *testing.T) {
	grid := testGridFromInput(t)
	station, _ := Part1(grid)
	assertVaporize(t, grid, station, 200, NewPos(26, 28))
}

package day10

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func testGrid(t *testing.T, r io.Reader) *Grid {
	grid, err := ReadGrid(r)
	if err != nil {
		t.Fatal(err)
	}
	return grid
}

func testGridFromInput(t *testing.T) *Grid {
	grid, err := OpenGrid("./day10.input")
	if err != nil {
		t.Fatal(err)
	}
	return grid
}

func assertCell(t *testing.T, grid *Grid, p Coord, v byte) {
	b, err := grid.Get(p)
	if err != nil {
		panic(err)
	}
	if b != v {
		panic(fmt.Sprintf("expected '%c', got '%c' at %v", v, b, p))
	}
}

func assertCoord(t *testing.T, actual Coord, expected Coord) bool {
	if expected == actual {
		return true
	}
	t.Errorf("expected: %v, got: %v", expected, actual)
	return false
}

func assertInt(t *testing.T, actual int, expected int) bool {
	if actual == expected {
		return true
	}
	t.Errorf("expected: %v, got: %v", expected, actual)
	return false
}

func assertBestStation(t *testing.T, grid *Grid, p Coord, n int) bool {
	actualP, actualN := Part1(grid)
	okP := assertCoord(t, actualP, p)
	okN := assertInt(t, actualN, n)
	ok := okP && okN
	if ok {
		return true
	}
	return false
}

func TestGrid(t *testing.T) {
	grid := testGridFromInput(t)
	assertInt(t, grid.width, 41)
	assertInt(t, grid.height, 41)
	assertCell(t, grid, Coord{0, 0}, '.')
	assertCell(t, grid, Coord{1, 0}, '#')
	assertCell(t, grid, Coord{40, 40}, '.')
	for i := 0; i < len(grid.data); i++ {
		p := grid.CoordOf(i)
		b, _ := grid.Get(p)
		if b != grid.data[i] {
			t.Fatalf("expected '%c', got '%c' at i: %d, coord: %v", grid.data[i], b, i, p)
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
	assertBestStation(t, grid, Coord{3, 4}, 8)
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
	assertBestStation(t, grid, Coord{5, 8}, 33)
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
	assertBestStation(t, grid, Coord{1, 2}, 35)
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
	assertBestStation(t, grid, Coord{6, 3}, 41)
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
	assertBestStation(t, grid, Coord{11, 13}, 210)
}

func TestDay10_Part1(t *testing.T) {
	grid := testGridFromInput(t)
	assertBestStation(t, grid, Coord{28, 29}, 340)
}

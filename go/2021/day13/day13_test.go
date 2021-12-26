package day13

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
	"aoc/internal/geo"
)

var answer = `####..##..#..#..##..#..#.###..####..##.
#....#..#.#.#..#..#.#.#..#..#....#.#..#
###..#....##...#....##...###....#..#...
#....#.##.#.#..#....#.#..#..#..#...#.##
#....#..#.#.#..#..#.#.#..#..#.#....#..#
#.....###.#..#..##..#..#.###..####..###
`

func openFixture(t *testing.T) (grid map[geo.Pos]struct{}, folds []geo.Pos) {
	grid = make(map[geo.Pos]struct{}, 32)
	folds = make([]geo.Pos, 0, 32)
	fixture.ScanStrings(t, 2021, 13, func(s string) (err error) {
		if s == "" {
			return
		}
		var x, y int
		if strings.HasPrefix(s, "fold along x=") {
			x, err = strconv.Atoi(s[13:])
			if err != nil {
				return
			}
			folds = append(folds, geo.Pos{X: x})
			return
		}
		if strings.HasPrefix(s, "fold along y=") {
			y, err = strconv.Atoi(s[13:])
			if err != nil {
				return
			}
			folds = append(folds, geo.Pos{Y: y})
			return
		}
		var p geo.Pos
		p, err = geo.ParsePos(s)
		if err != nil {
			return
		}
		grid[p] = struct{}{}
		return
	})
	return
}

func TestPart1(t *testing.T) {
	grid, folds := openFixture(t)
	grid = Fold(folds[0], grid)
	assert.Int(t, 687, len(grid), "bad dot count")
}

func TestPart2(t *testing.T) {
	grid, folds := openFixture(t)
	for _, fold := range folds {
		grid = Fold(fold, grid)
	}
	b := new(bytes.Buffer)
	PrintGrid(b, grid)
	if assert.String(t, answer, b.String(), "bad paper folds") {
		t.Log("\n" + b.String())
	}
}

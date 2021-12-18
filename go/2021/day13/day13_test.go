package day13

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"testing"

	. "aoc2021"
)

var answer = `####..##..#..#..##..#..#.###..####..##.
#....#..#.#.#..#..#.#.#..#..#....#.#..#
###..#....##...#....##...###....#..#...
#....#.##.#.#..#....#.#..#..#..#...#.##
#....#..#.#.#..#..#.#.#..#..#.#....#..#
#.....###.#..#..##..#..#.###..####..###
`

func mustOpenFixture(name string) (grid map[Pos]struct{}, folds []Pos) {
	f := MustOpenFixture(name)
	defer f.Close()
	grid = make(map[Pos]struct{}, 32)
	folds = make([]Pos, 0, 32)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		if strings.HasPrefix(s, "fold along x=") {
			x, err := strconv.Atoi(s[13:])
			if err != nil {
				panic(err)
			}
			folds = append(folds, Pos{X: x})
			continue
		}
		if strings.HasPrefix(s, "fold along y=") {
			y, err := strconv.Atoi(s[13:])
			if err != nil {
				panic(err)
			}
			folds = append(folds, Pos{Y: y})
			continue
		}
		p, err := ParsePos(s)
		if err != nil {
			panic(err)
		}
		grid[p] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func TestPart1(t *testing.T) {
	grid, folds := mustOpenFixture("day13")
	grid = Fold(folds[0], grid)
	AssertInt(t, 687, len(grid), "bad dot count")
}

func TestPart2(t *testing.T) {
	grid, folds := mustOpenFixture("day13")
	for _, fold := range folds {
		grid = Fold(fold, grid)
	}
	b := new(bytes.Buffer)
	PrintGrid(b, grid)
	if AssertString(t, answer, b.String(), "bad paper folds") {
		t.Log("\n" + b.String())
	}
}

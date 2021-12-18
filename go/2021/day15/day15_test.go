package day15

import (
	"testing"

	. "aoc2021"
)

func mustOpenFixture(name string) *Grid {
	f := MustOpenFixture(name)
	defer f.Close()
	g, err := ReadGrid(f)
	if err != nil {
		panic(err)
	}
	for i, c := range g.Data {
		g.Data[i] = c - '0'
	}
	return g
}

func TestPart1(t *testing.T) {
	g := mustOpenFixture("day15")
	AssertInt(t, 621, ShortestPath(g, 0, g.Len()-1), "bad shortest path")
}

func TestPart2(t *testing.T) {
	g := mustOpenFixture("day15")
	g = GrowGrid(g)
	AssertInt(t, 2904, ShortestPath(g, 0, g.Len()-1), "bad shortest path")
}

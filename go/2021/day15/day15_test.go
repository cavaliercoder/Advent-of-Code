package day15

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
	"aoc/internal/geo"
)

func openFixture(t *testing.T) *geo.Grid {
	g := fixture.Grid(t, 2021, 15)
	for i, c := range g.Data {
		g.Data[i] = c - '0'
	}
	return g
}

func TestPart1(t *testing.T) {
	g := openFixture(t)
	assert.Int(t, 621, ShortestPath(g, 0, g.Len()-1), "bad shortest path")
}

func TestPart2(t *testing.T) {
	g := openFixture(t)
	g = GrowGrid(g)
	assert.Int(t, 2904, ShortestPath(g, 0, g.Len()-1), "bad shortest path")
}

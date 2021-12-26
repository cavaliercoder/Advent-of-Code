package day15

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func TestPart1(t *testing.T) {
	g := fixture.Grid(t, 2021, 15).Normalize('0')
	assert.Int(t, 621, ShortestPath(g, 0, g.Len()-1), "bad shortest path")
}

func TestPart2(t *testing.T) {
	g := fixture.Grid(t, 2021, 15).Normalize('0')
	g = GrowGrid(g)
	assert.Int(t, 2904, ShortestPath(g, 0, g.Len()-1), "bad shortest path")
}

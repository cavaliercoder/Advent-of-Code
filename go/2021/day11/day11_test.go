package day11

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
	"aoc/internal/geo"
)

func openFixture(t *testing.T) *geo.Grid {
	g := fixture.Grid(t, 2021, 11)
	for i, c := range g.Data {
		g.Data[i] = c - '0'
	}
	return g
}

func TestPart1(t *testing.T) {
	flashes := 0
	g := openFixture(t)
	for i := 0; i < 100; i++ {
		flashes += Step(g)
	}
	assert.Int(t, 1755, flashes, "bad flash count")
}

func TestPart2(t *testing.T) {
	g := openFixture(t)
	steps := GetSyncSteps(g)
	assert.Int(t, 212, steps, "bad step count")
}

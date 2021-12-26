package day25

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func TestPart1(t *testing.T) {
	g := fixture.Grid(t, 2021, 25)
	assert.Int(t, 305, MaxSteps(g), "bad step count")
}

package day01

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func TestPart1(t *testing.T) {
	a := fixture.Ints(t, 2021, 1)
	assert.Int(t, 1462, CountIncreases(a), "Bad")
}

func TestPart2(t *testing.T) {
	a := fixture.Ints(t, 2021, 1)
	assert.Int(t, 1497, CountIncreasesSliding(a), "Bad")
}

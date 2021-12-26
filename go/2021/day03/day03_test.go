package day03

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) []int {
	a := make([]int, 0, 64)
	fixture.ScanBytes(t, 2021, 3, func(b []byte) error {
		a = append(a, Parse(b))
		return nil
	})
	return a
}

func TestPart1(t *testing.T) {
	a := openFixture(t)
	assert.Int(
		t,
		3969000,
		PowerConsumptionRate(a...),
		"bad power consumption rate",
	)
}

func TestPart2(t *testing.T) {
	a := openFixture(t)
	assert.Int(
		t,
		4267809,
		LifeSupportRating(a...),
		"bad life support rating",
	)
}

package day09

import (
	"aoc/internal/geo"
	"sort"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) *geo.Grid {
	g := fixture.Grid(t, 2021, 9)
	for i, b := range g.Data {
		g.Data[i] = b - '0'
	}
	return g
}

func TestPart1(t *testing.T) {
	assert.Int(t, 436, SumRisk(openFixture(t)), "bad risk sum")
}

func TestPart2(t *testing.T) {
	a := GetBasinSizes(openFixture(t))
	sort.Ints(a)
	sum := a[len(a)-3] * a[len(a)-2] * a[len(a)-1]
	assert.Int(t, 1317792, sum, "bad basin product")
}

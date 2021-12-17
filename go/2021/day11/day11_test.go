package day11

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
	flashes := 0
	g := mustOpenFixture("day11")
	for i := 0; i < 100; i++ {
		flashes += Step(g)
	}
	AssertInt(t, 1755, flashes, "bad flash count")
}

func TestPart2(t *testing.T) {
	g := mustOpenFixture("day11")
	steps := GetSyncSteps(g)
	AssertInt(t, 212, steps, "bad step count")
}

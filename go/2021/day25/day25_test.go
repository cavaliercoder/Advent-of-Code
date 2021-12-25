package day25

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
	return g
}

func TestPart1(t *testing.T) {
	g := mustOpenFixture("day25")
	AssertInt(t, 305, MaxSteps(g), "bad step count")
}

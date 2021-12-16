package day09

import (
	"sort"
	"testing"

	. "aoc2021"
)

func mustOpenFixture(name string) *Grid {
	f, err := OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	g, err := ReadGrid(f)
	if err != nil {
		panic(err)
	}
	for i, b := range g.Data {
		g.Data[i] = b - '0'
	}
	return g
}

func TestPart1(t *testing.T) {
	AssertInt(t, 436, SumRisk(mustOpenFixture("day09")), "bad risk sum")
}

func TestPart2(t *testing.T) {
	a := GetBasinSizes(mustOpenFixture("day09"))
	sort.Ints(a)
	sum := a[len(a)-3] * a[len(a)-2] * a[len(a)-1]
	AssertInt(t, 1317792, sum, "bad basin product")
}

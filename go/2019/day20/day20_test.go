package day20

import (
	"testing"

	. "aoc"
)

func TestPart1Example1(t *testing.T) {
	g, err := OpenGrid(Fixture("day20-example1"))
	if err != nil {
		panic(err)
	}
	AssertInt(t, 23, ShortestPath(g), "bad shortest distance")
}

func TestPart1Example2(t *testing.T) {
	g, err := OpenGrid(Fixture("day20-example2"))
	if err != nil {
		panic(err)
	}
	AssertInt(t, 58, ShortestPath(g), "bad shortest distance")
}

func TestPart1(t *testing.T) {
	g, err := OpenGrid(Fixture("day20"))
	if err != nil {
		panic(err)
	}
	AssertInt(t, 590, ShortestPath(g), "bad shortest distance")
}

func TestPart2Example1(t *testing.T) {
	g, err := OpenGrid(Fixture("day20-example3"))
	if err != nil {
		panic(err)
	}
	AssertInt(t, 396, ShortestPathRecursive(g), "bad shortest distance")
}

func TestPart2(t *testing.T) {
	g, err := OpenGrid(Fixture("day20"))
	if err != nil {
		panic(err)
	}
	AssertInt(t, 7180, ShortestPathRecursive(g), "bad shortest distance")
}

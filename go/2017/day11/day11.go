package day11

import (
	"fmt"
	"strings"
)

// maxabs returns the larger absolute value of a or b.
func maxabs(a, b int) int {
	if a < 0 {
		a = 0 - a
	}
	if b < 0 {
		b = 0 - b
	}
	if a < b {
		return b
	}
	return a
}

// step returns the new axial coordinates after a step is made in the given
// direction from the given axial coordinates.
func step(direction string, x, y int) (int, int) {
	switch strings.Trim(direction, "\n") {
	case "n":
		y--
	case "ne":
		x++
		y--
	case "se":
		x++
	case "s":
		y++
	case "sw":
		x--
		y++
	case "nw":
		x--
	default:
		panic(fmt.Sprintf("invalid direction: '%q'", direction))
	}
	return x, y
}

// ShortestPath is my solution to Part One. It computes the final location of
// a "program" after it has taken the given steps (of form "n,ne,se,s,sw,nw")
// and returns the shortest path back to the beginning in a hexagonal graph.
func ShortestPath(s string) int {
	var x, y int
	steps := strings.Split(s, ",")
	for _, d := range steps {
		x, y = step(d, x, y)
	}
	return maxabs(x, y)
}

// MaxShortestPath is my solution to Part Two. It computes the final location of
// a "program" after it has taken the given steps (of form "n,ne,se,s,sw,nw")
// and returns the maximum value of the shortest path back to the beginning of
// the hexagonal graph for each step taken.
func MaxShortestPath(s string) int {
	var x, y, m int
	steps := strings.Split(s, ",")
	for _, d := range steps {
		x, y = step(d, x, y)
		if n := maxabs(x, y); n > m {
			m = n
		}
	}
	return m
}

package day24

import (
	"bytes"
	"testing"

	. "aoc"
)

var example1 = []byte(`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
`)

func TestPartOne(t *testing.T) {
	tests := []struct {
		Input  []byte
		Expect int
	}{
		{example1, 31},
		{MustReadFixture("day24"), 1906},
	}
	for _, test := range tests {
		g := ParseGraph(bytes.NewReader(test.Input))
		actual := g.MaxBridgeStrength()
		if actual != test.Expect {
			t.Fatalf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  []byte
		Expect int
	}{
		{example1, 19},
		{MustReadFixture("day24"), 1824},
	}
	for _, test := range tests {
		g := ParseGraph(bytes.NewReader(test.Input))
		actual := g.LongestBridgeStrength()
		if actual != test.Expect {
			t.Fatalf("expected %d, got %d", test.Expect, actual)
		}
	}
}

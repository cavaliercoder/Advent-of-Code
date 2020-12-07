package day22

import (
	"bytes"
	"testing"

	. "aoc"
)

var example1 = []byte(`..#
#..
...
`)

func TestPartOne(t *testing.T) {
	tests := []struct {
		Input  []byte
		Steps  int
		Expect int
	}{
		{example1, 70, 41},
		{example1, 10000, 5587},
		{MustReadFixture("day22"), 10000, 5538},
	}
	for _, test := range tests {
		g := ParseGrid(bytes.NewReader(test.Input))
		actual := g.Infect(test.Steps)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  []byte
		Steps  int
		Expect int
	}{
		{example1, 100, 26},
		{example1, 10000000, 2511944},
		{MustReadFixture("day22"), 10000000, 2511090},
	}
	for _, test := range tests {
		g := ParseGrid(bytes.NewReader(test.Input))
		actual := g.InfectEvolved(test.Steps)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

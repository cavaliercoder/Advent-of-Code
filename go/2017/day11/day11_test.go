package day11

import (
	"testing"

	. "aoc"
)

func TestPartOne(t *testing.T) {
	tests := map[string]int{
		string(MustReadFixture("day11")): 834,
		"ne,ne,ne":                       3,
		"ne,ne,sw,sw":                    0,
		"ne,ne,s,s":                      2,
		"se,sw,se,sw,sw":                 3,
	}
	for test, expect := range tests {
		actual := ShortestPath(test)
		if actual != expect {
			t.Errorf("expected %d, got %d for %s", expect, actual, test)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := map[string]int{
		string(MustReadFixture("day11")): 1569,
		"ne,ne,ne":                       3,
		"ne,ne,sw,sw":                    2,
		"ne,ne,s,s":                      2,
		"se,sw,se,sw,sw":                 3,
	}
	for test, expect := range tests {
		actual := MaxShortestPath(test)
		if actual != expect {
			t.Errorf("expected %d, got %d for %s", expect, actual, test)
		}
	}
}

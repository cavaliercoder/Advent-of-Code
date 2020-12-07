package day12

import (
	"strings"
	"testing"

	. "aoc"
)

func TestPartOne(t *testing.T) {
	tests := map[string]int{
		`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`: 6,
		string(MustReadFixture("day12")): 380,
	}
	for test, expect := range tests {
		G := ParseGraph(strings.NewReader(test), 2000)
		actual := G.CountReachable(0)
		if actual != expect {
			t.Errorf("expected %d, got %d", expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  string
		Size   int
		Expect int
	}{
		{
			Input: `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`,
			Size:   7,
			Expect: 2,
		},
		{
			Input:  string(MustReadFixture("day12")),
			Size:   2000,
			Expect: 181,
		},
	}
	for _, test := range tests {
		G := ParseGraph(strings.NewReader(test.Input), test.Size)
		actual := G.CountGroups()
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

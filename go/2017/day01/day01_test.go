package day01

import (
	"io/ioutil"
	"testing"

	. "aoc"
)

func fixture(name string) []int {
	f, err := OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	A := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		A[i] = int(b[i] - byte('0'))
	}
	return A
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect int
	}{
		{[]int{1, 1, 2, 2}, 3},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
		{fixture("day01"), 1158},
	}
	for _, test := range tests {
		actual := SumPairs(test.Input)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect int
	}{
		{[]int{1, 2, 1, 2}, 6},
		{[]int{1, 2, 2, 1}, 0},
		{[]int{1, 2, 3, 4, 2, 5}, 4},
		{[]int{1, 2, 3, 1, 2, 3}, 12},
		{[]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
		{fixture("day01"), 1132},
	}
	for _, test := range tests {
		actual := SumSplitPairs(test.Input)
		if actual != test.Expect {
			t.Errorf("expected %d, got %d", test.Expect, actual)
		}
	}
}

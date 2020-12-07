package day01

import (
	"bufio"
	"strconv"
	"testing"

	. "aoc"
)

func ParseDeltas(name string) []int {
	f := MustOpenFixture("day01")
	defer f.Close()

	deltas := make([]int, 0, 8)
	s := bufio.NewScanner(f)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		deltas = append(deltas, n)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return deltas
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
		{ParseDeltas("fixture1"), 484},
	}
	for _, test := range tests {
		actual := ComputeFrequency(test.input)
		if actual != test.expect {
			t.Errorf(
				"expected: %v, got: %v, for input: %v",
				test.expect,
				actual,
				test.input)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{+1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
		{ParseDeltas("fixture1"), 367},
	}
	for _, test := range tests {
		actual := ComputeFirstDuplicateFrequency(test.input)
		if actual != test.expect {
			t.Errorf(
				"expected: %v, got: %v, for input: %v",
				test.expect,
				actual,
				test.input)
		}
	}
}

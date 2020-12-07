package day06

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	test := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	expect := 4074
	if actual := getLoopIndex(test); actual != expect {
		t.Errorf("expected %v, got %v", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	test := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	expect := 2793
	if actual := getLoopSize(test); actual != expect {
		t.Errorf("expected %v, got %v", expect, actual)
	}
}

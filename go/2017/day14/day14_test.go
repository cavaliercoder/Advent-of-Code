package day14

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	var tests = map[string]int{
		"flqrgnkx": 8108,
		"wenycdww": 8226,
	}
	for test, expect := range tests {
		actual := NewGrid(128, test).BitCount()
		if actual != expect {
			t.Errorf("expected %d, got %d", expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	var tests = map[string]int{
		"flqrgnkx": 1242,
		"wenycdww": 1128,
	}
	for test, expect := range tests {
		actual := NewGrid(128, test).GroupCount()
		if actual != expect {
			t.Errorf("expected %d, got %d", expect, actual)
		}
	}
}

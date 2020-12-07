package day10

import "testing"

func TestPartOne(t *testing.T) {
	test := []int{46, 41, 212, 83, 1, 255, 157, 65, 139, 52, 39, 254, 2, 86, 0, 204}
	expect := 52070
	actual := KnotHash(test)
	if actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	tests := map[string]string{
		"":         "a2582a3a0e66e6e86e3812dcb672a272",
		"AoC 2017": "33efeb34ea91902bb2f59c9920caa6cd",
		"1,2,3":    "3efbe78a8d82f29979031a4aa0b16a9d",
		"1,2,4":    "63960835bcdc130f0b66d7ff4f6a5a8e",
		"46,41,212,83,1,255,157,65,139,52,39,254,2,86,0,204": "7f94112db4e32e19cf6502073c66f9bb",
	}
	for test, expect := range tests {
		actual := KnotHashAdvanced([]byte(test))
		if actual != expect {
			t.Errorf("expected '%s', got '%s'", expect, actual)
		}
	}
}

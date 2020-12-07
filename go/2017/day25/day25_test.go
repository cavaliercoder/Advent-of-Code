package day25

import "testing"

func TestPartOne(t *testing.T) {
	expect := 2846
	actual := Checksum()
	if actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}

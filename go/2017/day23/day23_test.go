package day23

import (
	"bytes"
	"testing"

	. "aoc"
)

func TestPartOne(t *testing.T) {
	expect := 6241
	p := NewProcess(bytes.NewReader(MustReadFixture("day23")))
	p.Run()
	actual := p.mul
	if actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expect := 909
	actual := Optimized(1)
	if actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}

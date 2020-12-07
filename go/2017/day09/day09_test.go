package day09

import (
	"testing"

	. "aoc"
)

func assertInt(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("expected %d, got %d", a, b)
	}
}

func TestPartOne(t *testing.T) {
	p := NewParser(MustReadFixture("day09"))
	p.Parse()
	assertInt(t, 10050, p.score)
}
func TestPartTwo(t *testing.T) {
	p := NewParser(MustReadFixture("day09"))
	p.Parse()
	assertInt(t, 4482, p.garbage)
}

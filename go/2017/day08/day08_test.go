package day08

import (
	"bufio"
	"testing"

	. "aoc"
)

func TestPartOne(t *testing.T) {
	f := MustOpenFixture("day08")
	defer f.Close()
	expect := 5143
	r := make(RegisterFile, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ins, err := ParseInstruction(scanner.Text())
		if err != nil {
			t.Fatalf("Error parsing '%s': %v", scanner.Text(), err)
			continue
		}
		r.Mutate(ins)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading input: %v", err)
	}
	if actual := r.Max(); actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}

func TestPartTwo(t *testing.T) {
	f := MustOpenFixture("day08")
	defer f.Close()
	expect := 6209
	actual := ^(1 << 32)
	r := make(RegisterFile, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ins, err := ParseInstruction(scanner.Text())
		if err != nil {
			t.Fatalf("Error parsing '%s': %v", scanner.Text(), err)
			continue
		}
		r.Mutate(ins)
		n := r.Max()
		if n > actual {
			actual = n
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading input: %v", err)
	}
	if actual != expect {
		t.Errorf("expected %d, got %d", expect, actual)
	}
}

package day10

import (
	"bufio"
	"sort"
	"testing"

	. "aoc2021"
)

func mustOpenFixture(name string) [][]byte {
	f := MustOpenFixture(name)
	defer f.Close()
	a := make([][]byte, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := make([]byte, len(scanner.Bytes()))
		copy(b, scanner.Bytes())
		a = append(a, b)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return a
}

func TestPart1(t *testing.T) {
	score := 0
	a := mustOpenFixture("day10")
	for _, b := range a {
		score += CheckSyntax(b)
	}
	AssertInt(t, 294195, score, "bad syntax score")
}

func TestPart2(t *testing.T) {
	a := mustOpenFixture("day10")
	scores := make([]int, 0, len(a))
	for _, b := range a {
		if score := Autocomplete(b); score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	AssertInt(t, 3490802734, scores[len(scores)/2], "bad autocomplete score")
}

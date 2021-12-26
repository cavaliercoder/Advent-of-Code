package day10

import (
	"sort"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func TestPart1(t *testing.T) {
	score := 0
	fixture.ScanBytes(t, 2021, 10, func(b []byte) error {
		score += CheckSyntax(b)
		return nil
	})
	assert.Int(t, 294195, score, "bad syntax score")
}

func TestPart2(t *testing.T) {
	scores := make([]int, 0, 64)
	fixture.ScanBytes(t, 2021, 10, func(b []byte) error {
		if score := Autocomplete(b); score > 0 {
			scores = append(scores, score)
		}
		return nil
	})
	sort.Ints(scores)
	assert.Int(t, 3490802734, scores[len(scores)/2], "bad autocomplete score")
}

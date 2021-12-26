package day21

import (
	"testing"

	"aoc/internal/assert"
)

func TestPart1(t *testing.T) {
	assert.Int(t, 739785, Play(4, 8), "bad final score")
	assert.Int(t, 888735, Play(4, 6), "bad final score")
}

func TestPart2(t *testing.T) {
	assert.Int64(t, 444356092776315, PlayDirac(4, 8), "bad universe count")
	assert.Int64(t, 647608359455719, PlayDirac(4, 6), "bad universe count")
}

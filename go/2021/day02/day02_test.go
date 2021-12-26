package day02

import (
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func TestPart1(t *testing.T) {
	commands := fixture.Strings(t, 2021, 2)
	depth, hpos, err := PilotSub(commands)
	if err != nil {
		t.Fatal(err)
	}
	assert.Int(t, 1693300, depth*hpos, "bad submarine position")
}

func TestPart2(t *testing.T) {
	commands := fixture.Strings(t, 2021, 2)
	depth, hpos, err := PilotSubWithAim(commands)
	if err != nil {
		t.Fatal(err)
	}
	assert.Int(t, 1857958050, depth*hpos, "bad submarine position")
}

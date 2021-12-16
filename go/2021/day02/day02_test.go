package day02

import (
	"bufio"
	"testing"

	"aoc2021"
)

func mustOpenFixture(name string) []string {
	f, err := aoc2021.OpenFixture("day02")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	a := make([]string, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return a
}

func TestPart1(t *testing.T) {
	commands := mustOpenFixture("day02")
	depth, hpos, err := PilotSub(commands)
	if err != nil {
		t.Fatal(err)
	}
	aoc2021.AssertInt(t, 1693300, depth*hpos, "bad submarine position")
}

func TestPart2(t *testing.T) {
	commands := mustOpenFixture("day02")
	depth, hpos, err := PilotSubWithAim(commands)
	if err != nil {
		t.Fatal(err)
	}
	aoc2021.AssertInt(t, 1857958050, depth*hpos, "bad submarine position")
}

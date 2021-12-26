package day06

import (
	"strconv"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) []int {
	parts := strings.Split(strings.TrimSpace(fixture.String(t, 2021, 6)), ",")
	a := make([]int, 0, len(parts))
	for _, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			t.Fatalf("bad fish: %s", s)
		}
		a = append(a, int(n))
	}
	return a
}

func TestPart1(t *testing.T) {
	fishes := openFixture(t)
	assert.Int(t, 354564, GenerateFish(80, fishes), "bad fish count")
}

func TestPart2(t *testing.T) {
	fishes := openFixture(t)
	assert.Int(
		t,
		1609058859115,
		GenerateFish(256, fishes),
		"bad fish count",
	)
}

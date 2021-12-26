package day07

import (
	"strconv"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) []int {
	parts := strings.Split(strings.TrimSpace(fixture.String(t, 2021, 7)), ",")
	a := make([]int, 0, len(parts))
	for _, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			t.Fatalf("bad x-position: %s", s)
		}
		a = append(a, int(n))
	}
	return a
}

func TestPart1(t *testing.T) {
	assert.Int(
		t,
		37,
		AlignCrabmarines(16, 1, 2, 0, 4, 2, 7, 1, 2, 14),
		"bad fuel count",
	)
	positions := openFixture(t)
	assert.Int(
		t,
		336721,
		AlignCrabmarines(positions...),
		"bad fuel count",
	)
}

func TestPart2(t *testing.T) {
	assert.Int(
		t,
		168,
		AlignCrabmarinesProper(16, 1, 2, 0, 4, 2, 7, 1, 2, 14),
		"bad fuel count",
	)
	positions := openFixture(t)
	assert.Int(
		t,
		91638945,
		AlignCrabmarinesProper(positions...),
		"bad fuel count",
	)
}

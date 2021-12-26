package day05

import (
	"fmt"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func parseVent(s string) (*Vent, error) {
	var v Vent
	n, err := fmt.Sscanf(s, "%d,%d -> %d,%d", &v.A.X, &v.A.Y, &v.B.X, &v.B.Y)
	if err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, fmt.Errorf("bad vent: %s", s)
	}
	return &v, nil
}

func openFixture(t *testing.T) []*Vent {
	vents := make([]*Vent, 0, 64)
	fixture.ScanStrings(t, 2021, 5, func(s string) error {
		if s == "" {
			return nil
		}
		vent, err := parseVent(s)
		if err != nil {
			return err
		}
		vents = append(vents, vent)
		return nil
	})
	return vents
}

func TestPart1(t *testing.T) {
	vents := openFixture(t)
	assert.Int(
		t,
		6397,
		CountIntersects(false, vents...),
		"bad intersect count",
	)
}

func TestPart2(t *testing.T) {
	vents := openFixture(t)
	assert.Int(
		t,
		22335,
		CountIntersects(true, vents...),
		"bad intersect count",
	)
}

package day05

import (
	"bufio"
	"fmt"
	"testing"

	"aoc2021"
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

func mustOpenFixture(name string) []*Vent {
	f, err := aoc2021.OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	vents := make([]*Vent, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}
		vent, err := parseVent(s)
		if err != nil {
			panic(err)
		}
		vents = append(vents, vent)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return vents
}

func TestPart1(t *testing.T) {
	vents := mustOpenFixture("day05")
	aoc2021.AssertInt(
		t,
		6397,
		CountIntersects(false, vents...),
		"bad intersect count",
	)
}

func TestPart2(t *testing.T) {
	vents := mustOpenFixture("day05")
	aoc2021.AssertInt(
		t,
		22335,
		CountIntersects(true, vents...),
		"bad intersect count",
	)
}

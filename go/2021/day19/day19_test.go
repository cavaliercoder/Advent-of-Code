package day19

import (
	"fmt"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
	"aoc/internal/geo3d"
)

func openFixture(t *testing.T) []*Scanner {
	a := make([]*Scanner, 0, 32)
	var s *Scanner
	fixture.ScanStrings(t, 2021, 19, func(line string) error {
		if len(line) == 0 {
			return nil
		}
		if strings.HasPrefix(line, "--- scanner ") {
			s = NewScanner(geo3d.Pos{})
			a = append(a, s)
			return nil
		}
		beacon, err := geo3d.ParsePos(line)
		if err != nil {
			return fmt.Errorf("Bad beacon: %s", line)
		}
		s.Add(beacon)
		return nil
	})
	return a
}

func TestDay19(t *testing.T) {
	a := openFixture(t)
	field, ok := Merge(a...)
	if assert.Bool(t, true, ok, "Merge failed") {
		assert.Int(t, 432, len(field.Beacons), "Bad beacon count")
		assert.Int(t, 14414, field.MaxManhattan(), "Bad manhattan distance")
	}
}

package day19

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) []*Scanner {
	a := make([]*Scanner, 0, 32)
	var s *Scanner
	fixture.ScanStrings(t, 2021, 19, func(line string) error {
		if len(line) == 0 {
			return nil
		}
		if strings.HasPrefix(line, "--- scanner ") {
			s = NewScanner(Coord{})
			a = append(a, s)
			return nil
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return fmt.Errorf("bad beacon: %s", line)
		}
		var err error
		beacon := Coord{}
		beacon.X, err = strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("bad beacon X-coord: %s", line)
		}
		beacon.Y, err = strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("bad beacon Y-coord: %s", line)
		}
		beacon.Z, err = strconv.Atoi(parts[2])
		if err != nil {
			return fmt.Errorf("bad beacon Z-coord: %s", line)
		}
		s.Add(beacon)
		return nil
	})
	return a
}

func TestDay19(t *testing.T) {
	a := openFixture(t)
	field, ok := Merge(a...)
	if assert.Bool(t, true, ok, "merge failed") {
		assert.Int(t, 432, len(field.Beacons), "bad beacon count")
		assert.Int(t, 14414, field.MaxManhattan(), "bad manhattan distance")
	}
}

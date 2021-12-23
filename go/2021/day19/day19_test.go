package day19

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	. "aoc2021"
)

func readFixture(r io.Reader) ([]*Scanner, error) {
	a := make([]*Scanner, 0, 32)
	var s *Scanner
	lines := bufio.NewScanner(r)
	for lines.Scan() {
		line := lines.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "--- scanner ") {
			s = NewScanner(Coord{})
			a = append(a, s)
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("bad beacon: %s", line)
		}
		var err error
		beacon := Coord{}
		beacon.X, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("bad beacon X-coord: %s", line)
		}
		beacon.Y, err = strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("bad beacon Y-coord: %s", line)
		}
		beacon.Z, err = strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("bad beacon Z-coord: %s", line)
		}
		s.Add(beacon)
	}
	if err := lines.Err(); err != nil {
		return nil, err
	}
	return a, nil
}

func mustOpenFixture(name string) []*Scanner {
	f := MustOpenFixture(name)
	defer f.Close()
	field, err := readFixture(f)
	if err != nil {
		panic(err)
	}
	return field
}

func TestExample(t *testing.T) {
	a := mustOpenFixture("day19-example")
	field, ok := Merge(a...)
	if AssertBool(t, true, ok, "merge failed") {
		AssertInt(t, 79, len(field.Beacons), "bad beacon count")
		AssertInt(t, 3621, field.MaxManhattan(), "bad manhattan distance")
	}
}

func TestDay19(t *testing.T) {
	a := mustOpenFixture("day19")
	field, ok := Merge(a...)
	if AssertBool(t, true, ok, "merge failed") {
		AssertInt(t, 432, len(field.Beacons), "bad beacon count")
		AssertInt(t, 14414, field.MaxManhattan(), "bad manhattan distance")
	}
}

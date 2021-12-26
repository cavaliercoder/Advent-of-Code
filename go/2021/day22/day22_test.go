package day22

import (
	"fmt"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
	"aoc/internal/geo3d"
)

func parseOp(s string) (op Op, err error) {
	if strings.HasPrefix(s, "on ") {
		op.On = true
		s = s[3:]
	} else if strings.HasPrefix(s, "off ") {
		s = s[4:]
	} else {
		return Op{}, fmt.Errorf("invalid operation: %s", s)
	}
	op.Cube, err = parseCube(s)
	return
}

func parseCube(s string) (geo3d.Cube, error) {
	var c geo3d.Cube
	n, _ := fmt.Sscanf(
		s,
		"x=%d..%d,y=%d..%d,z=%d..%d",
		&c.A.X,
		&c.B.X,
		&c.A.Y,
		&c.B.Y,
		&c.A.Z,
		&c.B.Z,
	)
	if n != 6 {
		return geo3d.Cube{}, fmt.Errorf("invalid operation: %s", s)
	}
	c.B.X++
	c.B.Y++
	c.B.Z++
	return c, nil
}

func openFixture(t *testing.T) []Op {
	a := make([]Op, 0, 64)
	fixture.ScanStrings(t, 2021, 22, func(s string) error {
		op, err := parseOp(s)
		if err != nil {
			return err
		}
		a = append(a, op)
		return nil
	})
	return a
}

func TestPart1(t *testing.T) {
	ops := openFixture(t)
	r := NewReactor()
	assert.Int(t, 612714, r.Init(ops...), "bad init count")
}

func TestPart2(t *testing.T) {
	ops := openFixture(t)
	r := NewReactor()
	assert.Int(t, 1311612259117092, r.Reboot(ops...), "bad reboot count")
}

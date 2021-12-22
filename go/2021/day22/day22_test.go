package day22

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	. "aoc2021"
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

func parseCube(s string) (Cube, error) {
	var c Cube
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
		return Cube{}, fmt.Errorf("invalid operation: %s", s)
	}
	c.B.X++
	c.B.Y++
	c.B.Z++
	return c, nil
}

func mustOpenFixture(name string) []Op {
	f := MustOpenFixture(name)
	defer f.Close()
	a := make([]Op, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		op, err := parseOp(scanner.Text())
		if err != nil {
			panic(err)
		}
		a = append(a, op)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return a
}

func TestPart1(t *testing.T) {
	ops := mustOpenFixture("day22")
	r := NewReactor()
	AssertInt(t, 612714, r.Init(ops...), "bad init count")
}

func TestPart2(t *testing.T) {
	ops := mustOpenFixture("day22")
	r := NewReactor()
	AssertInt(t, 1311612259117092, r.Reboot(ops...), "bad reboot count")
}

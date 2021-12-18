package day14

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"

	. "aoc2021"
)

func mustReadFixture(r io.Reader) (tmpl *Template, rules map[Pair]byte) {
	scanner := bufio.NewScanner(r)
	if scanner.Scan() {
		tmpl = NewTemplate(scanner.Bytes()...)
	}
	rules = make(map[Pair]byte)
	scanner.Scan()
	for scanner.Scan() {
		b := scanner.Bytes()
		a := bytes.Split(b, []byte(" -> "))
		if len(a) != 2 {
			panic(fmt.Sprintf("bad pair insertion: %s", b))
		}
		if len(a[0]) != 2 {
			panic(fmt.Sprintf("bad pair insertion: %s", b))
		}
		if len(a[1]) != 1 {
			panic(fmt.Sprintf("bad pair insertion: %s", b))
		}
		rules[[2]byte{a[0][0], a[0][1]}] = a[1][0]
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func mustOpenFixture(name string) (tmpl *Template, rules map[Pair]byte) {
	f := MustOpenFixture(name)
	defer f.Close()
	return mustReadFixture(f)
}

func TestPart1(t *testing.T) {
	tmpl, rules := mustOpenFixture("day14")
	for i := 0; i < 10; i++ {
		tmpl.Step(rules)
	}
	AssertInt(t, 2170, tmpl.Hash(), "bad polymer hash")
}

func TestPart2(t *testing.T) {
	tmpl, rules := mustOpenFixture("day14")
	for i := 0; i < 40; i++ {
		tmpl.Step(rules)
	}
	AssertInt(t, 2422444761283, tmpl.Hash(), "bad polymer hash")
}

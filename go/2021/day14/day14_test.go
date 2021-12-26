package day14

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func readFixture(
	r io.Reader,
) (
	tmpl *Template,
	rules map[Pair]byte,
	err error,
) {
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
			err = fmt.Errorf("bad pair insertion: %s", b)
			return
		}
		if len(a[0]) != 2 {
			err = fmt.Errorf("bad pair insertion: %s", b)
			return
		}
		if len(a[1]) != 1 {
			err = fmt.Errorf("bad pair insertion: %s", b)
			return
		}
		rules[[2]byte{a[0][0], a[0][1]}] = a[1][0]
	}
	err = scanner.Err()
	return
}

func openFixture(t *testing.T) (tmpl *Template, rules map[Pair]byte) {
	f := fixture.Open(t, 2021, 14)
	defer f.Close()
	var err error
	tmpl, rules, err = readFixture(f)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestPart1(t *testing.T) {
	tmpl, rules := openFixture(t)
	for i := 0; i < 10; i++ {
		tmpl.Step(rules)
	}
	assert.Int(t, 2170, tmpl.Hash(), "bad polymer hash")
}

func TestPart2(t *testing.T) {
	tmpl, rules := openFixture(t)
	for i := 0; i < 40; i++ {
		tmpl.Step(rules)
	}
	assert.Int(t, 2422444761283, tmpl.Hash(), "bad polymer hash")
}

package main

import (
	. "aoc/2019/common"
	"aoc/2019/intcode"
	"testing"
)

func TestPart1(t *testing.T) {
	solution := []string{
		// if A or B or C are a hole
		"NOT A T",
		"OR T J",
		"NOT B T",
		"OR T J",
		"NOT C T",
		"OR T J",

		// and we can land on D, jump!
		"AND D J",
		"WALK",
	}
	data, err := intcode.OpenData(Fixture("day21"))
	if err != nil {
		t.Fatal(err)
	}
	v, err := intcode.RunASCII(data, solution...)
	if err != nil {
		t.Fatal((err))
	}
	AssertInt(t, 19350258, v, "bad hull damage")
}

func TestPart2(t *testing.T) {
	solution := []string{
		// if A or B or C are a hole
		"NOT A T",
		"OR T J",
		"NOT B T",
		"OR T J",
		"NOT C T",
		"OR T J",

		// and after we jump, we can move to E or jump to H
		"NOT J T",
		"OR E T",
		"OR H T",
		"AND T J",

		// and we can land on D, jump!
		"AND D J",
		"RUN",
	}
	data, err := intcode.OpenData(Fixture("day21"))
	if err != nil {
		t.Fatal(err)
	}
	v, err := intcode.RunASCII(data, solution...)
	if err != nil {
		t.Fatal((err))
	}
	AssertInt(t, 1142627861, v, "bad hull damage")
}

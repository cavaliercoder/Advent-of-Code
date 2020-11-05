package main

import (
	"testing"

	"aoc"
	"aoc/intcode"
)

var solution = []string{
	// need: prime number + fixed point + whirled peas + antenna
	"east",
	"take whirled peas",
	"north",
	"west",
	"south",
	"take antenna",
	"north",
	"east",
	"south",
	"east",
	"north",
	"take prime number",
	"south",
	"west",
	"west",
	"north",
	"take fixed point",
	"north",
	"east",
	"south",
}

func TestPart1(t *testing.T) {
	data, err := intcode.OpenData(aoc.Fixture("day25"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = intcode.RunASCII(data, solution...)
	if err != nil {
		t.Fatal(err)
	}
}

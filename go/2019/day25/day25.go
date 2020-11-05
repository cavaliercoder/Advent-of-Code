package main

import (
	"aoc"
	"aoc/intcode"
)

func main() {
	data, err := intcode.OpenData(aoc.Fixture("day25"))
	if err != nil {
		panic(err)
	}
	_, err = intcode.RunASCII(data)
	if err != nil {
		panic(err)
	}
}

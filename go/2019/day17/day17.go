package main

import (
	"log"

	. "aoc"
	"aoc/intcode"
)

func main() {
	data, err := intcode.OpenData(Fixture("day17"))
	if err != nil {
		log.Fatal(err)
	}
	data[0] = 2
	if _, err := intcode.RunASCII(data); err != nil {
		panic(err)
	}
}

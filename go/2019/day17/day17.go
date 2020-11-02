package main

import (
	"log"

	"aoc/2019/common"
	"aoc/2019/intcode"
)

func main() {
	data, err := intcode.OpenData(common.Fixture("day17"))
	if err != nil {
		log.Fatal(err)
	}
	data[0] = 2
	if _, err := intcode.RunASCII(data); err != nil {
		panic(err)
	}
}

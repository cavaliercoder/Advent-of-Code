package main

import (
	. "aoc/2019/common"
	"aoc/2019/intcode"
)

func main() {
	data, err := intcode.OpenData(Fixture("day21"))
	if err != nil {
		panic(err)
	}
	intcode.RunASCII(data)
}

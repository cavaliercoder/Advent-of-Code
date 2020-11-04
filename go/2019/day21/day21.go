package day21

import (
	. "aoc"
	"aoc/intcode"
)

func main() {
	data, err := intcode.OpenData(Fixture("day21"))
	if err != nil {
		panic(err)
	}
	intcode.RunASCII(data)
}

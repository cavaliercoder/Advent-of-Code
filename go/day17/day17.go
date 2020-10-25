package main

import (
	"fmt"
	"log"
	"time"

	"../common"
	"../intcode"
)

var solutions = []string{
	"A,B,A,C,A,B,A,C,B,C", // main()
	"R,4,L,12,L,8,R,4",    // A()
	"L,8,R,10,R,10,R,6",   // B()
	"R,4,R,10,L,12",       // C()
}

func printGrid(r intcode.BytesIO) {
	for i := 0; i < 41; i++ {
		s, err := r.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Printf(s)
	}
}

func printLine(r intcode.BytesIO) {
	s, err := r.ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Print(s)
}

func writeString(r intcode.BytesIO, s string) {
	_, err := r.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := intcode.OpenData(common.Fixture("day17"))
	if err != nil {
		log.Fatal(err)
	}
	data[0] = 2
	vm := intcode.New(data)
	r := intcode.NewBytesIO(vm)

	// show init grid
	printGrid(r)
	printLine(r)

	// send args
	for i := 0; i < len(solutions); i++ {
		printLine(r)
		fmt.Printf("> %s\n", solutions[i])
		writeString(r, solutions[i]+"\n")
	}

	// enable feed
	printLine(r)
	fmt.Printf("> y\n")
	writeString(r, "y\n")

	// run feed
	for i := 0; ; i++ {
		if i > 0 {
			time.Sleep(time.Millisecond * 50)
			fmt.Print("\033[42A")
		}
		printLine(r)
		printGrid(r)
	}
}

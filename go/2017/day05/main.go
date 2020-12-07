/*
This is my solution for Day 5 of the Advent of Code 2017.

Input is read - one per line - from standard input. The number of jumps required
to exit the instruction list is printed standard out.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing integer '%v': %v\n", scanner.Text(), err)
			os.Exit(1)
		}
		A = append(A, n)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", countStepsOne(A))
}

func countStepsOne(A []int) int {
	var i, n, c int
	for {
		if i < 0 || i >= len(A) {
			return c
		}
		n = i + A[i]
		A[i]++
		i = n
		c++
	}
	panic("this should never happen")
}

func countStepsTwo(A []int) int {
	var i, n, c int
	for {
		if i < 0 || i >= len(A) {
			return c
		}
		n = i + A[i]
		if A[i] < 3 {
			A[i]++
		} else {
			A[i]--
		}
		i = n
		c++
	}
	panic("this should never happen")
}

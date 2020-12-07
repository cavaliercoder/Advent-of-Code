/*
This is my solution for Day 2 of the Advent of Code 2017.

Input arrays are read - one per line - from standard input. The spreadsheet
checksum is printed to standard output.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "\t")
		A := make([]int, len(words))
		for i := 0; i < len(words); i++ {
			v, err := strconv.Atoi(words[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error parsing integer '%v': %v\n", words[i], err)
				os.Exit(1)
			}
			A[i] = v
		}
		sum += findEvenQuotient(A)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", sum)
}

// computeRange is the solution to Part One and returns the difference between
// the highest and lowest integer values in an array of integers.
//
// The function returns in O(n) time.
func computeRange(A []int) int {
	min := A[0]
	max := A[0]
	for i := 1; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	return max - min
}

// findEvenQuotient is the solution for Part Two and compares every possible
// quotient in an array of integers and returns the first quotient that is a
// whole number. It is assumed any given array has one, and only one whole
// number quotient.
//
// This function returns in O(n2) time, as every possible pair of integers must
// be evaluated. A minor optimisation is applied in that both orderings of all
// pairs are evaluated in a single pass and are therefore only ever compared
// once.
func findEvenQuotient(A []int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			a := A[i]
			b := A[j]
			if a < b {
				a, b = b, a
			}
			if a%b == 0 {
				return a / b
			}
		}
	}
	panic("this should never happen")
}

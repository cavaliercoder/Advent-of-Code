/*
Package day06 is my solution for Day 6 of the Advent of Code 2017.
*/
package day06

import (
	"bytes"
)

// getLoopIndex is my solution to Part One and returns the number reallocation
// cycles required before a previously encountered state is encountered.
func getLoopIndex(A []int) int {
	count := 0
	seen := map[string]bool{}
	for {
		key := hashKey(A)
		if seen[key] {
			return count
		}
		seen[key] = true
		balance(A)
		count++
	}
}

// getLoopIndex is my solution to Part Two and returns the number reallocation
// cycles inside the infinite loop.
func getLoopSize(A []int) int {
	count := 0
	seen := map[string]int{}
	for {
		key := hashKey(A)
		if ix, ok := seen[key]; ok {
			return count - ix
		}
		seen[key] = count
		balance(A)
		count++
	}
}

// balance redistributes the blocks from the largest memory bank across all
// memory banks. This happens in linear time over the length of the input
// array (number of memory banks).
func balance(A []int) {
	var ix, max int
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			ix = i
			max = A[i]
		}
	}
	A[ix] = 0
	r := max % len(A)
	d := (max - r) / len(A)
	for i := 0; i < len(A); i++ {
		ix++
		if i < r {
			A[ix%len(A)] += d + 1
		} else {
			A[ix%len(A)] += d
		}
	}
}

// hashKey generates a hash string representing the state of the given array.
// The hash can be used in a hash map to identify duplicate states in constant
// time.
//
// I could treat the input as a byte array and reallocate as a string, but I
// anticipate there may be inputs that quickly overflow a byte.
func hashKey(A []int) string {
	b := &bytes.Buffer{}
	for i := 0; i < len(A); i++ {
		b.WriteRune(rune(A[i]))
	}
	return b.String()
}

package day10

import (
	"encoding/hex"
)

const (
	n = 256
)

// KnotHash is my solution to part one.
func KnotHash(input []int) int {
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i
	}
	var pos, skip, a, b int
	for _, l := range input {
		for i := 0; i < l/2; i++ {
			a = (pos + i) % n
			b = (pos + l - i - 1) % n
			A[a], A[b] = A[b], A[a]
		}
		pos = (pos + l + skip) % n
		skip++
	}
	return A[0] * A[1]
}

// KnotHashAdvanced is my solution to part two.
func KnotHashAdvanced(input []byte) string {
	input = append(input, 17, 31, 73, 47, 23)
	A := make([]byte, n)
	for i := 0; i < n; i++ {
		A[i] = byte(i)
	}
	var pos, skip, a, b int
	for r := 0; r < 64; r++ {
		for _, l := range input {
			for i := byte(0); i < l/2; i++ {
				a = (pos + int(i)) % n
				b = (pos + int(l-i) - 1) % n
				A[a], A[b] = A[b], A[a]
			}
			pos = (pos + int(l) + skip) % n
			skip++
		}
	}
	h := make([]byte, 16)
	for i := 0; i < 16; i++ {
		var b byte
		for j := 0; j < 16; j++ {
			b ^= A[(16*i)+j]
		}
		h[i] = b
	}
	return hex.EncodeToString(h)
}

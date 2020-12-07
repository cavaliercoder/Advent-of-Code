package day01

// SumPairs returns the sum of all integers in the given array whose subsequent
// integers are equal. The array is treated as a circular array, meaning the
// subsequent integer for the last integer, is the first integer.
//
// Time complexity of this function if O(n).
func SumPairs(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		next := (i + 1) % len(A)
		if A[i] == A[next] {
			sum += A[i]
		}
	}
	return sum
}

// SumSplitPairs returns the sum of all integers in the given array whose
// counterpart integers are equal. Two integers are considered a counterpart if
// the distance between their indices is half the length of the array. The array
// is treated as a circular array and assumed to be even in length.
//
// Time complexity of this function if O(n).
func SumSplitPairs(A []int) int {
	if len(A)%2 != 0 {
		panic("array length is not even")
	}
	sum := 0
	for i := 0; i < len(A); i++ {
		next := (i + len(A)/2) % len(A)
		if A[i] == A[next] {
			sum += A[i]
		}
	}
	return sum
}

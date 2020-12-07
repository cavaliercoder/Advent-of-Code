package day01

// ComputeFrequency returns the resulting frequency, starting from zero after
// all given changes have been applied.
func ComputeFrequency(deltas []int) (v int) {
	for _, n := range deltas {
		v += n
	}
	return
}

// ComputeFirstDuplicateFrequency applies each given change in a loop and
// returns the first resulting frequency to be observed twice.
func ComputeFirstDuplicateFrequency(deltas []int) (v int) {
	seen := make(map[int]struct{}, 64)
	seen[0] = struct{}{}
	for i := 0; ; i = (i + 1) % len(deltas) {
		v += deltas[i]
		if _, ok := seen[v]; ok {
			return v
		}
		seen[v] = struct{}{}
	}
}

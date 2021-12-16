package day06

func GenerateFish(days int, fishes []int) int {
	// count fish at each possible stage
	var a [9]int
	for i := 0; i < len(a); i++ {
		for _, fish := range fishes {
			if fish == i {
				a[i]++
			}
		}
	}

	// for each day, spawn fish from stage 0 to stage day+7 % 9
	for i := 0; i < days; i++ {
		a[(i+7)%9] += a[i%9]
	}

	// sum
	n := 0
	for i := 0; i < len(a); i++ {
		n += a[i]
	}
	return n
}

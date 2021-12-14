package day01

func CountIncreases(a []int) int {
	n := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			n++
		}
	}
	return n
}

func CountIncreasesSliding(a []int) int {
	if len(a) < 4 {
		return 0
	}
	n := 0
	l := a[0] + a[1] + a[2]
	for i := 3; i < len(a); i++ {
		r := l + a[i] - a[i-3]
		if r > l {
			n++
		}
		l = r
	}
	return n
}

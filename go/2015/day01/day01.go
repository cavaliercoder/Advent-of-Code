package day01

func Elevate(b []byte) int {
	v := 0
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '(':
			v++
		case ')':
			v--
		}
	}
	return v
}

func GoToBasement(b []byte) int {
	v := 0
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '(':
			v++
		case ')':
			v--
		}
		if v == -1 {
			return i + 1
		}
	}
	return v
}

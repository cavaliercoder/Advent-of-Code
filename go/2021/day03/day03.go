package day03

const Mask = 0x0FFF

func Parse(b []byte) int {
	if len(b) != 12 {
		panic("bad input")
	}
	n := 0
	for _, c := range b {
		n <<= 1
		if c == '1' {
			n |= 1
		}
	}
	return n
}

func setBit(v, i int) int {
	if i < 0 || i > 11 {
		panic("bit out of range")
	}
	return v | (1 << (11 - i))
}

func bitIsSet(v, i int) bool {
	if i < 0 || i > 11 {
		panic("bit out of range")
	}
	return v&(1<<(11-i)) != 0
}

func Gamma(a ...int) int {
	var g int
	var tally [12]int
	for _, k := range a {
		for i := 0; i < len(tally); i++ {
			mask := 1 << (len(tally) - 1 - i)
			if k&mask != 0 {
				tally[i]++
			}
		}
	}
	for i, n := range tally {
		if n > len(a)/2 {
			g = setBit(g, i)
		}
	}
	return g
}

func PowerConsumptionRate(a ...int) int {
	g := Gamma(a...)
	return g * (^g & Mask)
}

func OxygenGeneratorRating(a ...int) int {
	var mask, result int
	for i := 0; i < 12; i++ {
		var ones, count, last int
		for _, k := range a {
			if k&mask != result {
				continue
			}
			last = k
			count++
			if bitIsSet(k, i) {
				ones++
			}
		}
		if count == 1 {
			return last
		}
		if ones >= count-ones {
			result = setBit(result, i)
		}
		mask = setBit(mask, i)
	}
	return result
}

func CO2ScrubberRating(a ...int) int {
	var mask, result int
	for i := 0; i < 12; i++ {
		var ones, count, last int
		for _, k := range a {
			if k&mask != result {
				continue
			}
			last = k
			count++
			if bitIsSet(k, i) {
				ones++
			}
		}
		if count == 1 {
			return last
		}
		if ones < count-ones {
			result = setBit(result, i)
		}
		mask = setBit(mask, i)
	}
	return result
}

func LifeSupportRating(a ...int) int {
	o := OxygenGeneratorRating(a...)
	co2 := CO2ScrubberRating(a...)
	return o * co2
}

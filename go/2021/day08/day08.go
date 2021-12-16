package day08

import "fmt"

type Display byte

func ParseDisplay(b []byte) (Display, error) {
	var v Display
	for _, c := range b {
		c = c - 'a'
		if c >= 7 {
			return 0, fmt.Errorf("bad display: %s", b)
		}
		v |= 1 << c
	}
	return v, nil
}

func (d Display) SegmentCount() (n int) {
	for d != 0 {
		n++
		d &= d - 1
	}
	return
}

func Rewire(a [10]Display) [10]Display {
	var known [10]Display
	for i := 0; i < len(a); i++ {
		switch a[i].SegmentCount() {
		case 2:
			known[1] = a[i]
		case 3:
			known[7] = a[i]
		case 4:
			known[4] = a[i]
		case 7:
			known[8] = a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i].SegmentCount() != 6 {
			continue
		}
		if a[i]&known[1] != known[1] {
			known[6] = a[i]
		} else if a[i]&known[4] == known[4] {
			known[9] = a[i]
		} else {
			known[0] = a[i]
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i].SegmentCount() != 5 {
			continue
		}
		if a[i]&known[1] == known[1] {
			known[3] = a[i]
		} else if a[i]&known[9] == a[i] {
			known[5] = a[i]
		} else {
			known[2] = a[i]
		}
	}
	return known
}

func Decode(cipher [10]Display, input ...Display) (output int) {
	m := make(map[Display]int, len(cipher))
	for i := 0; i < len(cipher); i++ {
		m[cipher[i]] = i
	}
	for _, c := range input {
		output *= 10
		output += m[c]
	}
	return
}

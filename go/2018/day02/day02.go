package day02

// ChecksumID returns true for hasDups if any one or more characters appears
// exactly twice in the given input, and true for hasTrips if any appears
// exactly three times.
func ChecksumID(id []byte) (hasDups, hasTrips bool) {
	m := [26]byte{}
	dups, trips := 0, 0
	for _, b := range id {
		b -= 'a'
		m[b]++
		switch m[b] {
		case 2:
			dups++
		case 3:
			trips++
			dups--
		case 4:
			trips--
		}
	}
	return dups > 0, trips > 0
}

// ChecksumIDList computes the checksum for the given list of IDs.
func ChecksumIDList(ids [][]byte) int {
	dups, trips := 0, 0
	for _, id := range ids {
		hasDups, hasTrips := ChecksumID(id)
		if hasDups {
			dups++
		}
		if hasTrips {
			trips++
		}
	}
	return dups * trips
}

// CompareIDs returns a new ID which removes all characters that are not equal
// in the given input IDs.
func CompareIDs(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("length mismatch")
	}
	v := make([]byte, 0, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			v = append(v, a[i])
		}
	}
	return v
}

// FindPrototypeIDs returns the output of CompareIDs for the first two ID pairs
// in the given ID list that differ by a single character. This currently runs
// in Θ(n²).
func FindPrototypeIDs(ids [][]byte) []byte {
	for i := 0; i < len(ids)-1; i++ {
		for j := i + 1; j < len(ids); j++ {
			v := CompareIDs(ids[i], ids[j])

			// correct pair will be off by a single character
			if len(v) == len(ids[i])-1 {
				return v
			}
		}
	}
	return nil
}

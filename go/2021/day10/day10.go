package day10

var chunks = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var errorValues = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompleteValues = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func CheckSyntax(b []byte) (score int) {
	stack := make([]byte, 0, len(b))
	for _, c := range b {
		if closer, ok := chunks[c]; ok {
			stack = append(stack, closer)
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
			continue
		}
		score = errorValues[c]
		break
	}
	return
}

func Autocomplete(b []byte) (score int) {
	stack := make([]byte, 0, len(b))
	for _, c := range b {
		if closer, ok := chunks[c]; ok {
			stack = append(stack, closer)
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
			continue
		}
		return 0
	}
	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		score *= 5
		score += autocompleteValues[c]
	}
	return score
}

package day04

type Board struct {
	data [25]int
}

func (c *Board) Call(n int) bool {
	for i, nn := range c.data {
		if n == nn {
			c.data[i] = -1
			break
		}
	}
	for i := 0; i < 5; i++ {
		if c.checkRow(i) || c.checkColumn(i) {
			return true
		}
	}
	return false
}

func (c *Board) checkRow(n int) bool {
	offset := n * 5
	for i := 0; i < 5; i++ {
		if c.data[offset+i] >= 0 {
			return false
		}
	}
	return true
}

func (c *Board) checkColumn(n int) bool {
	offset := n
	for i := 0; i < 25; i += 5 {
		if c.data[offset+i] >= 0 {
			return false
		}
	}
	return true
}

func (c *Board) Score() int {
	n := 0
	for i := 0; i < len(c.data); i++ {
		if c.data[i] >= 0 {
			n += c.data[i]
		}
	}
	return n
}

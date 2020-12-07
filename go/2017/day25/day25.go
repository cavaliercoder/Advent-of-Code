package day25

const (
	checksumAfter = 12994925
)

// Node represents a single slot on the infinite Tape drive.
type Node struct {
	v    int
	l, r *Node
}

// Left returns the Node to the left of the given Node.
func (n *Node) Left() *Node {
	if n.l == nil {
		n.l = &Node{r: n}
	}
	return n.l
}

// Right returns the Node to the right of the given Node.
func (n *Node) Right() *Node {
	if n.r == nil {
		n.r = &Node{l: n}
	}
	return n.r
}

// A StateFunc is a function that given a Node, might mutate the state of the
// Node and returns the next Node and StateFunc to call.
type StateFunc func(*Node) (*Node, StateFunc)

func StateA(n *Node) (*Node, StateFunc) {
	if n.v == 0 {
		n.v = 1
		return n.Right(), StateB
	}
	n.v = 0
	return n.Left(), StateF
}

func StateB(n *Node) (*Node, StateFunc) {
	if n.v == 0 {
		return n.Right(), StateC
	}
	n.v = 0
	return n.Right(), StateD
}

func StateC(n *Node) (*Node, StateFunc) {
	if n.v == 0 {
		n.v = 1
		return n.Left(), StateD
	}
	return n.Right(), StateE
}

func StateD(n *Node) (*Node, StateFunc) {
	if n.v == 0 {
		return n.Left(), StateE
	}
	n.v = 0
	return n.Left(), StateD
}

func StateE(n *Node) (*Node, StateFunc) {
	if n.v == 0 {
		return n.Right(), StateA
	}
	return n.Right(), StateC
}

func StateF(n *Node) (*Node, StateFunc) {
	if n.v == 0 {
		n.v = 1
		return n.Left(), StateA
	}
	return n.Right(), StateA
}

// Checksum is my solution to Part One. It applies applies the above StateFuncs
// and counts the number of non-zero slots after checksumAfter iterations.
func Checksum() int {
	s := StateA
	n := &Node{}
	for i := 0; i < checksumAfter; i++ {
		n, s = s(n)
	}
	count := n.v
	for m := n.l; m != nil; m = m.l {
		count += m.v
	}
	for m := n.r; m != nil; m = m.r {
		count += m.v
	}
	return count
}

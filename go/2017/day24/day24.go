package day24

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	ID           int
	PortA, PortB int
	Edges        []*Node
}

// CanConnect returns true if the two given Nodes can connect on either port.
func (n *Node) CanConnect(a *Node) bool {
	return n.PortA == a.PortA ||
		n.PortA == a.PortB ||
		n.PortB == a.PortA ||
		n.PortB == a.PortB
}

// FindMaxStrength returns the maximum strength bridge that can be formed
// starting from the given Node. Consumed nodes are tracked (and never
// re-traversed) in the seen bitmap. The ingress port of the given node is
// excluded from evaluation as it has also been consumed by the previous node.
func (n *Node) FindMaxStrength(seen uint64, ingress int) (max int) {
	if seen&(1<<uint64(n.ID)) != 0 {
		return 0
	}
	seen |= 1 << uint64(n.ID)
	egress := n.PortA
	if ingress == n.PortA {
		egress = n.PortB
	}
	for i := 0; i < len(n.Edges); i++ {
		if n.Edges[i].PortA == egress || n.Edges[i].PortB == egress {
			m := n.Edges[i].FindMaxStrength(seen, egress)
			if m > max {
				max = m
			}
		}
	}
	max += n.PortA + n.PortB
	return
}

// FindMaxLength returns the length and strength of the longest bridge that can
// be formed starting from the given Node. Consumed nodes are tracked (and never
// re-traversed) in the seen bitmap. The ingress port of the given node is
// excluded from evaluation as it has also been consumed by the previous node.
func (n *Node) FindMaxLength(seen uint64, ingress int) (length, strength int) {
	if seen&(1<<uint64(n.ID)) != 0 {
		return
	}
	seen |= 1 << uint64(n.ID)
	egress := n.PortA
	if ingress == n.PortA {
		egress = n.PortB
	}
	for i := 0; i < len(n.Edges); i++ {
		if n.Edges[i].PortA == egress || n.Edges[i].PortB == egress {
			l, s := n.Edges[i].FindMaxLength(seen, egress)
			if l == length && s > strength {
				strength = s
			} else if l > length {
				length = l
				strength = s
			}
		}
	}
	length++
	strength += n.PortA + n.PortB
	return
}

// Graph is an adjacency list representation of a graph of Nodes and their
// possible edges.
type Graph []*Node

// ParseGraph parses the challenge input from the given reader into a Graph of
// Nodes.
func ParseGraph(r io.Reader) Graph {
	g := make(Graph, 0)
	s := bufio.NewScanner(r)
	i := 0
	for s.Scan() {
		tkns := strings.Split(s.Text(), "/")
		a, _ := strconv.Atoi(tkns[0])
		b, _ := strconv.Atoi(tkns[1])
		g = append(g, &Node{ID: i, PortA: a, PortB: b})
		i++
	}
	for i := 0; i < len(g)-1; i++ {
		for j := i + 1; j < len(g); j++ {
			if g[i].CanConnect(g[j]) {
				g[i].Edges = append(g[i].Edges, g[j])
				g[j].Edges = append(g[j].Edges, g[i])
			}
		}
	}
	return g
}

// MaxBridgeStrength is my solution to Part One and returns the maximum strength
// bridge that can be constructed from the available components.
func (g Graph) MaxBridgeStrength() (max int) {
	for i := 0; i < len(g); i++ {
		if g[i].PortA == 0 || g[i].PortB == 0 {
			s := g[i].FindMaxStrength(0, 0)
			if s > max {
				max = s
			}
		}
	}
	return
}

// LongestBridgeStrength is my solution to Part Two and returns the strength of
// the longest bridge that can be constructed from the available components.
func (g Graph) LongestBridgeStrength() int {
	length, strength := 0, 0
	for i := 0; i < len(g); i++ {
		if g[i].PortA == 0 || g[i].PortB == 0 {
			l, s := g[i].FindMaxLength(0, 0)
			if l == length && s > strength {
				strength = s
			} else if l > length {
				length = l
				strength = s
			}
		}
	}
	return strength
}

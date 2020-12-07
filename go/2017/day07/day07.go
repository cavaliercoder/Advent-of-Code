/*
Package day07 is my solution for Day 7 of the Advent of Code 2017.
*/
package day07

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	Name        string
	Weight      int
	SubWeight   int
	Children    []*Node
	ParentCount int
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: make([]*Node, 0),
	}
}

// Graph is a map of Nodes, indexed by Node name.
type Graph map[string]*Node

// getBottomProgram is my solution to Part One. It searches the graph in linear
// time for the bottom node, identified by having a ParentCount of zero.
func getBottomProgram(G Graph) *Node {
	for _, node := range G {
		if node.ParentCount == 0 {
			return node
		}
	}
	return nil
}

// getCorrectProgramWeight is my solution to Part Two. It searches the graph for
// the node whose weight is incorrect and returns the desired weight to
// rebalance the "Program Tower".
func getCorrectProgramWeight(G Graph) int {
	root := getBottomProgram(G)
	computeSubWeights(root)
	return bfs(root, 0)
}

// computeSubWeights recurses through the graph from the given root node and
// computes and stores the sum weight of each node (i.e. the weight of the node
// plus all of its ancestors).
func computeSubWeights(V *Node) int {
	V.SubWeight = V.Weight
	for _, child := range V.Children {
		V.SubWeight += computeSubWeights(child)
	}
	return V.SubWeight
}

// bfs performs a breadth-first search, looking for the node whose subweight is
// unbalanced from its siblings. The desired weight to correct the issue is
// returned.
func bfs(V *Node, want int) int {
	a := make(map[int]int, 0)
	b := make(map[int]*Node, 0)
	for _, child := range V.Children {
		a[child.SubWeight]++
		b[child.SubWeight] = child
	}
	if len(a) < 2 {
		return V.Weight - (V.SubWeight - want)
	}
	var next *Node
	for weight, count := range a {
		if count == 1 {
			next = b[weight]
		} else {
			want = weight
		}
	}
	return bfs(next, want)
}

// parseGraph reads an adjacency list format graph from the given reader and
// parses it into a Graph type.
func parseGraph(r io.Reader) (Graph, error) {
	G := Graph{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		tkns := strings.Split(scanner.Text(), " ")
		if len(tkns) < 2 {
			return nil, fmt.Errorf("invalid node: %s", scanner.Text())
		}

		node := G[tkns[0]]
		if node == nil {
			node = NewNode(tkns[0])
			G[tkns[0]] = node
		}

		weight := strings.Trim(tkns[1], "()")
		node.Weight, _ = strconv.Atoi(weight)

		for i := 3; i < len(tkns); i++ {
			name := strings.TrimRight(tkns[i], ",")
			child := G[name]
			if child == nil {
				child = NewNode(name)
				G[name] = child
			}
			child.ParentCount++
			node.Children = append(node.Children, child)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return G, nil
}

package day12

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	Id    int
	Edges []int
	Seen  bool
}

type Graph []*Node

// ParseGraph reads a graph with the given node count from the given reader.
func ParseGraph(r io.Reader, n int) Graph {
	G := make(Graph, n)
	for i := 0; i < n; i++ {
		G[i] = &Node{
			Id:    i,
			Edges: make([]int, 0),
		}
	}
	s := bufio.NewScanner(r)
	for s.Scan() {
		tkns := strings.SplitN(s.Text(), " ", 3)
		i, err := strconv.Atoi(tkns[0])
		if err != nil {
			panic(err)
		}
		v := G[i]
		tkns = strings.Split(tkns[2], ", ")
		for _, tkn := range tkns {
			e, err := strconv.Atoi(tkn)
			if err != nil {
				panic(err)
			}
			v.Edges = append(v.Edges, e)
			G[e].Edges = append(G[e].Edges, v.Id)
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return G
}

// CountReachable is my solution to Part One and returns the number of nodes
// reachable from the given node in a graph.
func (G Graph) CountReachable(v int) int {
	if G[v].Seen {
		return 0
	}
	G[v].Seen = true
	c := 1
	for _, e := range G[v].Edges {
		c += G.CountReachable(e)
	}
	return c
}

// CountGroups is my solution to Part Two and returns the number of subgraphs in
// the graph.
func (G Graph) CountGroups() int {
	c := 0
	for i := 0; i < len(G); i++ {
		if G.CountReachable(i) > 0 {
			c++
		}
	}
	return c
}

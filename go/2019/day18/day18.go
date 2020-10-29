package main

import (
	"bytes"
	"fmt"

	. "aoc/2019/common"
)

// KeyMask represents a set of keys or doors.
type KeyMask uint32

// NewKeyMask returns a new KeyMask containing only the given key.
func NewKeyMask(key Tile) KeyMask {
	if key >= 'a' && key <= 'z' {
		return 1 << (key - 'a')
	}
	if key >= 'A' && key <= 'Z' {
		return 1 << (key - 'A')
	}
	panic(fmt.Sprintf("not a key or door: %c", key))
}

// Add returns a new KeyMask which contains a copy of this KeyMask plus the
// given key.
func (k KeyMask) Add(key Tile) KeyMask {
	return k | NewKeyMask(key)
}

// Contains returns true if the given key or door is contained in this KeyMask.
func (k KeyMask) Contains(key Tile) bool {
	return NewKeyMask(key)&k != 0
}

// Unlock returns true if the keys in this KeyMask can unlock all doors in the
// given KeyMask.
func (k KeyMask) Unlock(doors KeyMask) bool {
	return doors&^k == 0
}

func (k KeyMask) String() string {
	b := bytes.NewBuffer([]byte{'{'})
	for i := 0; i < 26; i++ {
		if (k>>i)&0x01 == 0 {
			continue
		}
		if b.Len() > 1 {
			b.Write([]byte{',', ' '})
		}
		b.WriteByte('A' + byte(i))
	}
	b.WriteByte('}')
	return b.String()
}

// GetAllKeys returns a KeyMask representing every key in a graph
func GetAllKeys(g *Grid) KeyMask {
	result := KeyMask(0)
	for i := 0; i < len(g.Data); i++ {
		tile := Tile(g.Data[i])
		if tile.IsKey() {
			result = result.Add(tile)
		}
	}
	return result
}

type Tile byte

func (b Tile) IsWall() bool { return b == '#' }
func (b Tile) IsKey() bool  { return b >= 'a' && b <= 'z' }
func (b Tile) IsDoor() bool { return b >= 'A' && b <= 'Z' }

// Edge describes the path to another key.
type Edge struct {
	Pos      Pos
	Tile     Tile
	Mask     KeyMask
	Distance int
}

// GetEdges computes the distance and KeyMask required for all keys in the
// grid from the given start position.
func GetEdges(g *Grid, start Pos) map[Tile]Edge {
	result := make(map[Tile]Edge)
	startEdge := Edge{Pos: start, Tile: Tile(g.Get(start))}
	seen := make(map[Pos]Edge)
	seen[startEdge.Pos] = startEdge
	queue := make([]Edge, 1)
	queue[0] = startEdge
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, npos := range cell.Pos.URDL() {
			if !g.Contains(npos) {
				continue
			}
			if _, ok := seen[npos]; ok {
				continue
			}
			ncell := Edge{
				Pos:      npos,
				Tile:     Tile(g.Get(npos)),
				Mask:     cell.Mask,
				Distance: cell.Distance + 1,
			}
			seen[npos] = ncell
			if ncell.Tile.IsDoor() {
				ncell.Mask = ncell.Mask.Add(ncell.Tile)
			}
			if ncell.Tile.IsKey() {
				result[ncell.Tile] = ncell
			}
			if !ncell.Tile.IsWall() {
				queue = append(queue, ncell)
			}
		}
	}
	return result
}

// GetAllEdges returns all edges betweens between all keys, mapped by source and
// destination key.
//
// Invariant: Start positions must not be reachable from eachother.
func GetAllEdges(g *Grid, start ...Pos) map[Tile]map[Tile]Edge {
	result := make(map[Tile]map[Tile]Edge)

	for _, pos := range start {
		srcTile := Tile(g.Get(pos))
		result[srcTile] = make(map[Tile]Edge)

		// find all edges from start
		edges := GetEdges(g, pos)
		for dstTile, edge := range edges {
			result[srcTile][dstTile] = edge
		}
		// find all other possible paths
		for _, edge := range edges {
			result[edge.Tile] = GetEdges(g, edge.Pos)
		}
	}
	return result
}

// Split the grid into four subgrids at the start position.
func Split(g *Grid) {
	starts := FindStarts(g)
	if len(starts) != 1 {
		panic("grid must have only 1 '@' tile")
	}
	start := starts[0]
	g.Set(start.Add(Pos{X: -1, Y: -1}), '@')
	g.Set(start.Add(Pos{X: 0, Y: -1}), '#')
	g.Set(start.Add(Pos{X: 1, Y: -1}), '@')
	g.Set(start.Add(Pos{X: -1, Y: 0}), '#')
	g.Set(start, '#')
	g.Set(start.Add(Pos{X: 1, Y: 0}), '#')
	g.Set(start.Add(Pos{X: -1, Y: 1}), '@')
	g.Set(start.Add(Pos{X: 0, Y: 1}), '#')
	g.Set(start.Add(Pos{X: 1, Y: 1}), '@')
}

// FindStarts returns the position of all '@' entry points in a grid.
func FindStarts(g *Grid) []Pos {
	v := make([]Pos, 0)
	for i, a := range g.Data {
		if a == '@' {
			v = append(v, g.Pos(i))
		}
	}
	return v
}

// ShortestPath returns the distance of shortest possible path to collect all
// keys in the given grid.
func ShortestPath(g *Grid) int {
	type State struct {
		Bots    [4]Tile
		KeyMask KeyMask
	}

	// rewrite start positions as bots 1..N
	starts := FindStarts(g)
	initState := State{}
	for i := range starts {
		g.Set(starts[i], '1'+byte(i))
		initState.Bots[i] = Tile('1' + byte(i))
	}
	paths := GetAllEdges(g, starts...)
	// fmt.Printf("%s\n", initState.Bots)

	finalKeyMask := GetAllKeys(g)
	finalDistance := -1
	distances := make(map[State]int)
	queue := make([]State, 0, 256)
	queue = append(queue, initState)
	for len(queue) > 0 {
		srcState := queue[0]
		queue = queue[1:]

		// try move a bot
		for i, srcTile := range srcState.Bots {
			if srcTile == 0 {
				continue // no such bot
			}
			botMoved := false
			for destTile, dstEdge := range paths[srcTile] {
				// Have we already seen this key?
				if srcState.KeyMask.Contains(destTile) {
					continue
				}

				// Is this key accessible?
				if !srcState.KeyMask.Unlock(dstEdge.Mask) {
					continue
				}

				// compute next state
				dstState := State{
					Bots:    srcState.Bots,
					KeyMask: srcState.KeyMask.Add(destTile),
				}
				dstState.Bots[i] = destTile

				// compute distance to the next state
				distance := distances[srcState] + dstEdge.Distance
				if knownDistance, ok := distances[dstState]; ok {
					if knownDistance < distance {
						continue
					}
				}
				distances[dstState] = distance
				if dstState.KeyMask == finalKeyMask {
					if finalDistance < 0 || distance < finalDistance {
						finalDistance = distance
					}
					continue
				}
				queue = append(queue, dstState)
				botMoved = true
			}
			if botMoved {
				break
			}
		}
	}

	return finalDistance
}

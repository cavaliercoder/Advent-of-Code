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
func GetAllEdges(g *Grid, start Pos) map[Tile]map[Tile]Edge {
	startTile := Tile(g.Get(start))
	startPaths := GetEdges(g, start)
	result := make(map[Tile]map[Tile]Edge)
	result[startTile] = startPaths
	for _, cell := range startPaths {
		result[cell.Tile] = GetEdges(g, cell.Pos)
	}
	return result
}

// ShortestPath returns the distance of shortest possible path to collect all
// keys in the given grid.
func ShortestPath(g *Grid) int {
	type State struct {
		Tile    Tile
		KeyMask KeyMask
	}
	pos := g.Pos(g.Find('@'))
	paths := GetAllEdges(g, pos)
	finalKeyMask := GetAllKeys(g)
	finalDistance := 0
	distances := make(map[State]int)
	queue := make([]State, 0, 256)
	queue = append(queue, State{
		Tile:    Tile('@'),
		KeyMask: KeyMask(0),
	})
	for len(queue) > 0 {
		srcState := queue[0]
		queue = queue[1:]
		for destKey, edge := range paths[srcState.Tile] {
			if srcState.KeyMask.Contains(destKey) {
				continue
			}
			if !srcState.KeyMask.Unlock(edge.Mask) {
				continue
			}
			dstState := State{
				Tile:    destKey,
				KeyMask: srcState.KeyMask.Add(destKey),
			}
			distance := distances[srcState] + edge.Distance
			if knownDistance, ok := distances[dstState]; ok {
				if knownDistance < distance {
					continue
				}
			}
			distances[dstState] = distance
			if dstState.KeyMask == finalKeyMask {
				if finalDistance == 0 || distance < finalDistance {
					finalDistance = distance
				}
				continue
			}
			queue = append(queue, dstState)
		}
	}

	return finalDistance
}

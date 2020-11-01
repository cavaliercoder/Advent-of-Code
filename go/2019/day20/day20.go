package main

import (
	. "aoc/2019/common"
	"fmt"
)

type PortalID [2]byte

type Portal struct {
	ID    PortalID
	Pos   Pos
	Inner bool
}

func (p Portal) String() string {
	inner := ":I"
	if !p.Inner {
		inner = ":O"
	}
	return fmt.Sprintf("%s%s(%d, %d)", p.ID, inner, p.Pos.X, p.Pos.Y)
}

func isCapital(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func DecodePortal(g *Grid, pos Pos) (p Portal, ok bool) {
	b := g.GetWithDefault(pos, ' ')
	if !isCapital(b) {
		return
	}
	u := g.GetWithDefault(Pos{X: pos.X, Y: pos.Y - 1}, ' ')
	r := g.GetWithDefault(Pos{X: pos.X + 1, Y: pos.Y}, ' ')
	d := g.GetWithDefault(Pos{X: pos.X, Y: pos.Y + 1}, ' ')
	l := g.GetWithDefault(Pos{X: pos.X - 1, Y: pos.Y}, ' ')
	if isCapital(u) && d == '.' {
		return Portal{ID: [2]byte{u, b}, Pos: Pos{X: pos.X, Y: pos.Y + 1}, Inner: pos.Y > g.Height/2}, true
	}
	if isCapital(d) && u == '.' {
		return Portal{ID: [2]byte{b, d}, Pos: Pos{X: pos.X, Y: pos.Y - 1}, Inner: pos.Y < g.Height/2}, true
	}
	if isCapital(l) && r == '.' {
		return Portal{ID: [2]byte{l, b}, Pos: Pos{X: pos.X + 1, Y: pos.Y}, Inner: pos.X > g.Width/2}, true
	}
	if isCapital(r) && l == '.' {
		return Portal{ID: [2]byte{b, r}, Pos: Pos{X: pos.X - 1, Y: pos.Y}, Inner: pos.X < g.Width/2}, true
	}
	return
}

func ParsePortals(g *Grid) []Portal {
	portals := make([]Portal, 0, 8)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			pos := Pos{X: x, Y: y}
			portal, ok := DecodePortal(g, pos)
			if !ok {
				continue
			}
			portals = append(portals, portal)
		}
	}
	return portals
}

func FindEdgesFromPortal(g *Grid, portal Portal) map[Portal]int {
	result := make(map[Portal]int)
	seen := make(map[Pos]int)
	seen[portal.Pos] = 0
	queue := make([]Pos, 1, 64)
	queue[0] = portal.Pos
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		distance := seen[pos]
		neighbors := pos.URDL()
		for _, npos := range neighbors {
			// is there already a shorter path here?
			ndistance, ok := seen[npos]
			if ok && ndistance <= distance+1 {
				continue // already a short path here
			}
			seen[npos] = distance + 1

			// did we find a portal?
			if nportal, ok := DecodePortal(g, npos); ok {
				if nportal == portal {
					continue // self
				}
				result[nportal] = distance // current position is the portal, not neighbor
				continue
			}

			// is the position impassable?
			b := g.GetWithDefault(npos, ' ')
			if b != '.' {
				continue
			}

			// enqueue next position
			queue = append(queue, npos)
		}
	}
	return result
}

func FindEdges(g *Grid, portals []Portal) map[Portal]map[Portal]int {
	result := make(map[Portal]map[Portal]int)
	for _, portal := range portals {
		edges := FindEdgesFromPortal(g, portal)

		// create an edge for the portal's counterpart
		for _, nportal := range portals {
			if nportal.ID == portal.ID && nportal.Pos != portal.Pos {
				edges[nportal] = 1
				break
			}
		}
		result[portal] = edges
	}
	return result
}

func ShortestPath(g *Grid) int {
	// get all portals and edges between them
	portals := ParsePortals(g)
	edges := FindEdges(g, portals)

	// identify portals AA and ZZ
	var AA, ZZ Portal
	for _, portal := range portals {
		if portal.ID == (PortalID{'A', 'A'}) {
			AA = portal
		}
		if portal.ID == (PortalID{'Z', 'Z'}) {
			ZZ = portal
		}
	}

	// compute shortest path
	seen := make(map[Portal]int)
	seen[AA] = 0
	queue := make([]Portal, 1, 64)
	queue[0] = AA
	for len(queue) > 0 {
		src := queue[0]
		queue = queue[1:]
		srcDistance := seen[src]
		for dst, dstDistance := range edges[src] {
			totalDistance := srcDistance + dstDistance
			seenDistance, ok := seen[dst]
			if ok && seenDistance <= totalDistance {
				continue
			}
			seen[dst] = totalDistance
			queue = append(queue, dst)
		}
	}
	return seen[ZZ]
}

func ShortestPathRecursive(g *Grid) int {
	// get all portals and edges between them
	portals := ParsePortals(g)
	edges := FindEdges(g, portals)

	// identify start and end states
	type State struct {
		Portal Portal
		Level  int
	}
	var AA, ZZ State
	for _, portal := range portals {
		if portal.ID == (PortalID{'A', 'A'}) {
			AA = State{Portal: portal, Level: 0}
		}
		if portal.ID == (PortalID{'Z', 'Z'}) {
			ZZ = State{Portal: portal, Level: 0}
		}
	}

	// compute shortest path
	queue := make([]State, 1, 64)
	queue[0] = AA
	seen := make(map[State]int)
	seen[AA] = 0
	shortestDistance := -1
	for len(queue) > 0 {
		src := queue[0]
		queue = queue[1:]
		srcDistance := seen[src]
		for dstPortal, dstDistance := range edges[src.Portal] {
			// block portals according to level
			if src.Level == 0 {
				// at level 0, all outer portals are blocked except self and ZZ
				if !dstPortal.Inner && dstPortal != ZZ.Portal && dstPortal.ID != src.Portal.ID {
					continue
				}
			} else {
				// at other levels, only AA and ZZ are blocked
				if dstPortal == AA.Portal || dstPortal == ZZ.Portal {
					continue
				}
			}

			// compute destination level
			level := src.Level
			if src.Portal.ID == dstPortal.ID {
				if src.Portal.Inner && !dstPortal.Inner {
					// inner to outer
					level++
				}
				if !src.Portal.Inner && dstPortal.Inner {
					// outer to inner
					level--
				}
			}

			// compute distance to next state
			dst := State{Portal: dstPortal, Level: level}
			totalDistance := srcDistance + dstDistance
			if shortestDistance >= 0 && totalDistance > shortestDistance {
				// there is a shorter path to ZZ
				continue
			}

			// ensure this is the shortest path to this destination
			seenDistance, ok := seen[dst]
			if ok && seenDistance <= totalDistance {
				// there is a shorter path to this state
				continue
			}
			seen[dst] = totalDistance

			// note shortest distance if dst == ZZ at level 0
			if dst == ZZ {
				if shortestDistance < 0 || shortestDistance > totalDistance {
					shortestDistance = totalDistance
				}
				continue
			}
			queue = append(queue, dst)
		}
	}
	return shortestDistance
}

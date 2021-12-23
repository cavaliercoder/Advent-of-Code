package day19

import (
	"fmt"

	. "aoc2021"
)

// Tranform is a 3D rotation matrix.
type Transform [9]int

// Rotations contains all possible 3D rotations with increments of 90 degrees.
// Courtesy: https://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
var Rotations = []Transform{
	{1, 0, 0, 0, 1, 0, 0, 0, 1}, // identity
	{0, 0, 1, 0, 1, 0, -1, 0, 0},
	{-1, 0, 0, 0, 1, 0, 0, 0, -1},
	{0, 0, -1, 0, 1, 0, 1, 0, 0},

	{0, -1, 0, 1, 0, 0, 0, 0, 1},
	{0, 0, 1, 1, 0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0, 0, 0, 0, -1},
	{0, 0, -1, 1, 0, 0, 0, -1, 0},

	{0, 1, 0, -1, 0, 0, 0, 0, 1},
	{0, 0, 1, -1, 0, 0, 0, -1, 0},
	{0, -1, 0, -1, 0, 0, 0, 0, -1},
	{0, 0, -1, -1, 0, 0, 0, 1, 0},

	{1, 0, 0, 0, 0, -1, 0, 1, 0},
	{0, 1, 0, 0, 0, -1, -1, 0, 0},
	{-1, 0, 0, 0, 0, -1, 0, -1, 0},
	{0, -1, 0, 0, 0, -1, 1, 0, 0},

	{1, 0, 0, 0, -1, 0, 0, 0, -1},
	{0, 0, -1, 0, -1, 0, -1, 0, 0},
	{-1, 0, 0, 0, -1, 0, 0, 0, 1},
	{0, 0, 1, 0, -1, 0, 1, 0, 0},

	{1, 0, 0, 0, 0, 1, 0, -1, 0},
	{0, -1, 0, 0, 0, 1, -1, 0, 0},
	{-1, 0, 0, 0, 0, 1, 0, 1, 0},
	{0, 1, 0, 0, 0, 1, 1, 0, 0},
}

type Coord struct {
	X, Y, Z int
}

func (c Coord) Add(other Coord) Coord {
	v := c
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
	return v
}

func (c Coord) Sub(other Coord) Coord {
	v := c
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
	return v
}

func (c Coord) Transform(t Transform) Coord {
	v := c
	v.X = c.X*t[0] + c.Y*t[1] + c.Z*t[2]
	v.Y = c.X*t[3] + c.Y*t[4] + c.Z*t[5]
	v.Z = c.X*t[6] + c.Y*t[7] + c.Z*t[8]
	return v
}

func (c Coord) Manhattan() int {
	return Abs(c.X) + Abs(c.Y) + Abs(c.Z)
}

func (c Coord) String() string {
	return fmt.Sprintf("<%d,%d,%d>", c.X, c.Y, c.Z)
}

type Scanner struct {
	Position Coord
	Beacons  map[Coord]struct{}

	rotations []*Scanner
	moves     map[Coord]*Scanner
}

func NewScanner(position Coord, beacons ...Coord) *Scanner {
	v := &Scanner{
		Position: position,
		Beacons:  make(map[Coord]struct{}, 32),
		moves:    make(map[Coord]*Scanner),
	}
	v.Add(beacons...)
	return v
}

func (c *Scanner) Add(beacons ...Coord) {
	for _, obj := range beacons {
		c.Beacons[obj] = struct{}{}
	}
}

func (c *Scanner) IsInRange(p Coord) bool {
	if p.X < c.Position.X-1000 ||
		p.Y < c.Position.Y-1000 ||
		p.Z < c.Position.Z-1000 {
		return false
	}
	if p.X > c.Position.X+1000 ||
		p.Y > c.Position.Y+1000 ||
		p.Z > c.Position.Z+1000 {
		return false
	}
	return true
}

func (c *Scanner) HasBeacon(p Coord) bool {
	_, ok := c.Beacons[p]
	return ok
}

func (c *Scanner) Transform(t Transform) *Scanner {
	v := NewScanner(c.Position.Transform(t))
	for c := range c.Beacons {
		v.Add(c.Transform(t))
	}
	return v
}

func (c *Scanner) Rotations() []*Scanner {
	if c.rotations != nil {
		return c.rotations
	}
	c.rotations = make([]*Scanner, len(Rotations))
	for i, t := range Rotations {
		c.rotations[i] = c.Transform(t)
	}
	return c.rotations
}

func (c *Scanner) Move(offset Coord) *Scanner {
	if v, ok := c.moves[offset]; ok {
		return v
	}
	v := NewScanner(c.Position.Add(offset))
	for obj := range c.Beacons {
		v.Add(obj.Add(offset))
	}
	c.moves[offset] = v
	return v
}

type Field struct {
	Scanners []*Scanner
	Beacons  map[Coord]struct{}
}

func NewField(scanners ...*Scanner) *Field {
	v := &Field{
		Scanners: make([]*Scanner, 0, 64),
		Beacons:  make(map[Coord]struct{}, 4096),
	}
	v.Add(scanners...)
	return v
}

func (c *Field) Add(scanners ...*Scanner) {
	c.Scanners = append(c.Scanners, scanners...)
	for _, s := range scanners {
		for obj := range s.Beacons {
			c.Beacons[obj] = struct{}{}
		}
	}
}

func (c *Field) HasBeacon(beacon Coord) bool {
	_, ok := c.Beacons[beacon]
	return ok
}

func (c *Field) CanFit(s *Scanner) bool {
	n := 0
	for beacon := range s.Beacons {
		if c.HasBeacon(beacon) {
			n++
			if n == 12 {
				break
			}
		}
	}
	if n < 12 {
		return false
	}
	for beacon := range c.Beacons {
		if !s.IsInRange(beacon) {
			continue
		}
		if !s.HasBeacon(beacon) {
			return false
		}
	}
	return true
}

func (c *Field) MaxManhattan() int {
	max := 0
	for i := 0; i < len(c.Scanners)-1; i++ {
		for j := 1; j < len(c.Scanners); j++ {
			n := c.Scanners[i].Position.Sub(c.Scanners[j].Position).Manhattan()
			if n > max {
				max = n
			}
		}
	}
	return max
}

func Merge(scanners ...*Scanner) (field *Field, ok bool) {
	// TODO: I'm very slow...
	field = NewField()
	if len(scanners) == 0 {
		ok = true
		return
	}
	field.Add(scanners[0])
	scanners[0] = nil
	merges, needMerges := 1, len(scanners)
	for merges < needMerges {
		merged := false
		for i, scanner := range scanners {
			if scanner == nil {
				continue
			}
			r := tryRotations(field, scanner)
			if r == nil {
				continue
			}
			field.Add(r)
			scanners[i] = nil
			merged = true
			merges++
			fmt.Printf("merged %d/%d...\n", merges, needMerges)
		}
		if !merged {
			break
		}
	}
	ok = merges == needMerges
	return
}

func tryRotations(field *Field, s *Scanner) *Scanner {
	for _, r := range s.Rotations() {
		if v := tryAlignments(field, r); v != nil {
			return v
		}
	}
	return nil
}

func tryAlignments(field *Field, s *Scanner) *Scanner {
	for a := range s.Beacons {
		for b := range field.Beacons {
			offset := b.Sub(a)
			s2 := s.Move(offset)
			if field.CanFit(s2) {
				return s2
			}
		}
	}
	return nil
}

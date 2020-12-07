package day20

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
)

// Plane represents the position, velocity and acceleration of a particle, along
// a single plane.
type Plane struct {
	p, v, a int
}

// Step computes the value of a plane (including position, velocity and
// acceleration) at time i, starting from the current values.
func (c Plane) Step(i int) Plane {
	c.p = c.p + i*c.v + c.a*(i*(i+1)/2)
	c.v = c.v + c.a*(i*(i+1)/2)
	return c
}

// FindIntersects return a list of times values at which the given trajectories
// intersect. Intersections might occur once, twice or not at all.
func FindIntersects(p0, p1 Plane) []int {
	// If p₀(t) = p₁(t), then:
	// 0 = t²(½(a₀ - a₁)) + t(½(a₀ - a₁) + v₀ - v₁) + p₀ - p₁
	//
	// Express as quadratic formula: `ax² + bx + c`
	// a = ½(a₀ - a₁)
	// b = ½(a₀ - a₁) + v₀ - v₁
	// c = p₀ - p₁
	a := float64(p0.a-p1.a) / 2
	b := a + float64(p0.v-p1.v)
	c := float64(p0.p - p1.p)

	// solve linear collision
	if a == 0.0 {
		t := -c / b
		return []int{int(t)}
	}

	// check intersects using discriminant form:
	//
	//    Δ = b² - 4ac
	//
	d := b*b - 4*a*c
	if d < 0 {
		return []int{}
	}

	// solve quadtratic intersects using:
	//
	//     t = -b ± √(b² - 4ac) / 2a
	//
	v := []int{}
	t0 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	t1 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	if t0 >= 0 && t0 == math.Trunc(t0) {
		v = append(v, int(t0))
	}
	if t1 >= 0 && t1 == math.Trunc(t1) {
		v = append(v, int(t1))
	}
	return v
}

// A Particle is a point in 3-dimensional space with properties position,
// velocity and acceleration.
type Particle struct {
	x, y, z Plane
}

// ParseParticles read particles from an io.Reader and expects one particle per
// line.
func ParseParticles(r io.Reader) []Particle {
	particles := make([]Particle, 0)
	s := bufio.NewScanner(r)
	for s.Scan() {
		p := Particle{}
		fmt.Sscanf(s.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.x.p, &p.y.p, &p.z.p,
			&p.x.v, &p.y.v, &p.z.v,
			&p.x.a, &p.y.a, &p.z.a,
		)
		particles = append(particles, p)
	}
	return particles
}

// Acceleration returns the acceleration of a Particle away from [0, 0, 0],
// calculated using the "Manhattan Distance" formula.
func (p Particle) Acceleration() int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	return abs(p.x.a) + abs(p.y.a) + abs(p.z.a)
}

// IntersectsAfter returns true if the given particles intersect on all planes
// at time t.
func IntersectsAfter(p0, p1 Particle, t int) bool {
	p0 = p0.Step(t)
	p1 = p1.Step(t)
	return p0.x.p == p1.x.p &&
		p0.y.p == p1.y.p &&
		p0.z.p == p1.z.p
}

// Step computes the values of all planes of a Particle (including position,
// velocity and acceleration) at time i, starting from the current values.
func (p Particle) Step(i int) Particle {
	p.x = p.x.Step(i)
	p.y = p.y.Step(i)
	p.z = p.z.Step(i)
	return p
}

// NearestParticle is my solution to Part One and returns the index of the
// Particle which accelerates the least, and therefore stays closest to 0,0,0
// over time. The function completes in O(n) linear time, where n is the number
// of Particles.
func NearestParticle(p []Particle) int {
	ans := 0
	d := p[0].Acceleration()
	for i := 1; i < len(p); i++ {
		v := p[i].Acceleration()
		if v < d {
			d = v
			ans = i
		}
	}
	return ans
}

// A collision represents a collision between two particles at time t and stores
// the index of the two particles, a, b.
type collision struct {
	a, b, t int
}

// collisions implements sort.Interface for a slice of collisions.
type collisions []collision

func (c collisions) Len() int           { return len(c) }
func (c collisions) Less(i, j int) bool { return c[i].t < c[1].t }
func (c collisions) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

// CountSurvivors is my solution for Part Two and completes in O(n²) where n is
// the number of particles. The evaluation of each Particle pair completes in
// constant time using the quadratic formula.
func CountSurvivors(p []Particle) int {
	// enumerate possible particle collisions - O(n²)
	c := make(collisions, 0)
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			t := append(FindIntersects(p[i].x, p[j].x))
			// BUG: what if a Particle accelerates on a single plane other than x?
			for x := 0; x < len(t); x++ {
				if IntersectsAfter(p[i], p[j], t[x]) {
					c = append(c, collision{i, j, t[x]})
				}
			}
		}
	}

	// eliminate particles in chronological order - O(n log n)
	sort.Sort(c)
	m := make([]int, len(p))
	for i := 0; i < len(c); i++ {
		// BUG: what if one of the particles was annihilated earlier?
		if m[c[i].a] == 0 || m[c[i].b] == 0 {
			m[c[i].a] = 1
			m[c[i].b] = 1
		}
	}

	// count remaining particles
	count := len(p)
	for i := 0; i < len(p); i++ {
		count -= m[i]
	}
	return count
}

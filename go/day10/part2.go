package day10

import (
	"errors"
	"math"
	"sort"
)

type byDistance struct {
	coords []Coord
	origin Coord
}

func (c *byDistance) Len() int { return len(c.coords) }
func (c *byDistance) Less(i, j int) bool {
	return c.coords[i].Subtract(c.origin).Distance() < c.coords[j].Subtract(c.origin).Distance()
}
func (c *byDistance) Swap(i, j int) { c.coords[i], c.coords[j] = c.coords[j], c.coords[i] }

// ByDistance returns a sort.Interface than will sort Coords by their distance
// from the given origin.
func ByDistance(coords []Coord, origin Coord) sort.Interface {
	return &byDistance{coords: coords, origin: origin}
}

// asteroidsByPlane groups all asteroids in the grid by their angle from the
// given origin. Each group is sorted by distance from the origin.
func asteroidsByPlane(grid *Grid, origin Coord) map[float64][]Coord {
	m := make(map[float64][]Coord)
	for i := 0; i < len(grid.data); i++ {
		if grid.data[i] != CellAsteroid {
			continue
		}
		asteroid := grid.CoordOf(i)
		offset := asteroid.Subtract(origin)
		if offset.IsZero() {
			continue // skip origin
		}
		plane := offset.Degrees()
		if asteroids, ok := m[plane]; ok {
			m[plane] = append(asteroids, asteroid)
		} else {
			asteroids = make([]Coord, 1)
			asteroids[0] = asteroid
			m[plane] = asteroids
		}
	}

	// sort each plane by distance ascending
	for _, asteroids := range m {
		sort.Sort(ByDistance(asteroids, origin))
	}
	return m
}

type degreeFunc func(n float64) float64

// mapPlanes changes the map key of all Coords to the output of f(key).
// Useful to rotate or invert the value of all planes, etc.
func mapPlanes(m map[float64][]Coord, f degreeFunc) map[float64][]Coord {
	v := make(map[float64][]Coord, len(m))
	for plane, coords := range m {
		plane = math.Mod(360+f(plane), 360)
		v[plane] = coords
	}
	return v
}

func sortedKeys(m map[float64][]Coord) []float64 {
	v := make([]float64, 0, len(m))
	for key := range m {
		v = append(v, key)
	}
	sort.Float64s(v)
	return v
}

// Part2 starts the laser at 90° and cycles it clockwise until n asteroids have
// been destroyed. The coordinates of the Nth asteroid are returned.
func Part2(grid *Grid, station Coord, n int) (Coord, error) {
	// group asteroids by plane and rotate 90° to meet the laser at index 0.
	m := asteroidsByPlane(grid, station)
	m = mapPlanes(m, func(n float64) float64 { return n + 90 })

	// sort planes so 0° is at index 0
	planes := sortedKeys(m)

	// track asteroid count so we don't loop forever is n > asteroidCount
	asteroidCount := grid.Count(CellAsteroid)

	// start shootin'
	for i := 0; asteroidCount > 0; i++ {
		i = i % len(planes)
		plane := planes[i]
		asteroids := m[plane]
		if len(asteroids) == 0 {
			// this plane is full destroyed
			continue
		}

		// kill the closest asteroid on this plane
		killed := asteroids[0]
		m[plane] = asteroids[1:]
		asteroidCount--
		n--
		if n == 0 {
			return killed, nil
		}
	}
	return Coord{}, errors.New("insufficient asteroids")
}

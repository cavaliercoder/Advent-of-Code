package day10

func observableAsteroids(grid *Grid, p Coord) int {
	m := make(map[float64]Coord)
	for i := 0; i < len(grid.data); i++ {
		if grid.data[i] != CellAsteroid {
			continue
		}
		q := grid.CoordOf(i)
		offset := q.Subtract(p)
		if offset.IsZero() {
			continue
		}
		Θ := offset.Degrees()
		m[Θ] = q
	}
	return len(m)
}

// Part1 returns the best number of observable asteroids from any possible
// station.
func Part1(grid *Grid) (Coord, int) {
	maxN := 0
	maxP := Coord{}
	for i := 0; i < len(grid.data); i++ {
		if grid.data[i] != CellAsteroid {
			continue
		}
		p := grid.CoordOf(i)
		n := observableAsteroids(grid, p)
		if n > maxN {
			maxN = n
			maxP = p
		}
	}
	return maxP, maxN
}

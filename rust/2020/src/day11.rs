#[cfg(test)]
mod tests {
    use crate::grid::{Grid, Point};

    fn occupied_adjacent(state: &Grid, p: Point) -> i32 {
        let mut n: i32 = 0;
        for p2 in Point::all_directions().map(|dir| dir + p) {
            match state.get(p2) {
                Some(v) => {
                    if v == b'#' {
                        n += 1;
                    }
                }
                None => (),
            }
        }
        n
    }

    fn occupied_visible(state: &Grid, p: Point) -> i32 {
        let mut n: i32 = 0;
        for dir in Point::all_directions() {
            let mut p2 = p.clone();
            loop {
                p2 = p2 + dir;
                match state.get(p2) {
                    Some(v) => {
                        match v {
                            b'L' => {
                                break;
                            } // empty seat seen
                            b'#' => {
                                // occupied seat seen
                                n += 1;
                                break;
                            }
                            _ => {} // keep looking
                        }
                    }
                    None => {
                        break;
                    } // no more in this direction
                }
            }
        }
        n
    }

    fn next_state(
        state: &Grid,
        neighbor_threshold: i32,
        neighbor_func: fn(&Grid, Point) -> i32,
    ) -> Option<Grid> {
        let mut next = Grid::new(state.width, state.height, b'?');
        let mut mutated = false;
        for (p, v) in state.iter() {
            next[p] = match v {
                b'L' => {
                    if neighbor_func(state, p) == 0 {
                        mutated = true;
                        b'#'
                    } else {
                        b'L'
                    }
                }
                b'#' => {
                    if neighbor_func(state, p) >= neighbor_threshold {
                        mutated = true;
                        b'L'
                    } else {
                        b'#'
                    }
                }
                _ => v,
            }
        }
        if mutated {
            Some(next)
        } else {
            None
        }
    }

    #[test]
    fn test_part1() {
        let mut grid = Grid::from_fixture("day11");
        loop {
            match next_state(&grid, 4, occupied_adjacent) {
                Some(next) => {
                    grid = next;
                }
                None => {
                    break;
                }
            }
        }
        assert_eq!(grid.count(b'#'), 2441);
    }

    #[test]
    fn test_part2() {
        let mut grid = Grid::from_fixture("day11");
        loop {
            match next_state(&grid, 5, occupied_visible) {
                Some(next) => {
                    grid = next;
                }
                None => {
                    break;
                }
            }
        }
        assert_eq!(grid.count(b'#'), 2190);
    }
}

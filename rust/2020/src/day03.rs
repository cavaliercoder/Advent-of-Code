#[cfg(test)]
mod tests {
  use crate::grid::{Grid, Point};

  fn wrap_point(grid: &Grid, p: Point) -> Point {
    Point {
      x: p.x % grid.width,
      y: p.y,
    }
  }

  fn count_trees(grid: &Grid, slope: Point) -> i32 {
    let mut p = Point::zero();
    let mut n: i32 = 0;
    while grid.contains(p) {
      if grid[p] == b'#' {
        n += 1;
      }
      p = wrap_point(&grid, p + slope);
    }
    n
  }

  #[test]
  fn test_part1() {
    let grid = Grid::from_fixture("day03");
    assert_eq!(count_trees(&grid, Point::new(3, 1)), 254);
  }

  #[test]
  fn test_part2() {
    let grid = Grid::from_fixture("day03");
    let mut n = count_trees(&grid, Point::new(1, 1));
    n *= count_trees(&grid, Point::new(3, 1));
    n *= count_trees(&grid, Point::new(5, 1));
    n *= count_trees(&grid, Point::new(7, 1));
    n *= count_trees(&grid, Point::new(1, 2));
    assert_eq!(n, 1666768320);
  }
}

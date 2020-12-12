use std::fmt;
use std::fmt::Write;
use std::iter::Iterator;
use std::ops::{Add, AddAssign, Index, IndexMut, Mul, Sub};

use crate::fixtures::Fixture;

#[derive(Copy, Clone, Debug)]
pub struct Point {
  pub x: i32,
  pub y: i32,
}

pub static NORTH: Point = Point { x: 0, y: -1 };
pub static NORTH_EAST: Point = Point { x: 1, y: -1 };
pub static NORTH_WEST: Point = Point { x: -1, y: -1 };
pub static SOUTH: Point = Point { x: 0, y: 1 };
pub static SOUTH_EAST: Point = Point { x: 1, y: 1 };
pub static SOUTH_WEST: Point = Point { x: -1, y: 1 };
pub static EAST: Point = Point { x: 1, y: 0 };
pub static WEST: Point = Point { x: -1, y: 0 };

static ALL_DIRECTIONS: [Point; 8] = [
  EAST, NORTH_EAST, NORTH, NORTH_WEST, WEST, SOUTH_WEST, SOUTH, SOUTH_EAST,
];

impl Point {
  pub fn new(x: i32, y: i32) -> Point {
    Point { x: x, y: y }
  }

  /// An iterator visiting all points adjacent to (0, 0).
  ///
  /// Includes diagonals. Order is counter-clockwise starting in the east.
  pub fn all_directions() -> impl Iterator<Item = Point> {
    ALL_DIRECTIONS.iter().cloned()
  }

  pub fn zero() -> Point {
    Self::new(0, 0)
  }

  pub fn manhattan_distance(&self) -> i32 {
    self.x.abs() + self.y.abs()
  }

  /// Rotate around origin, counter-clockwise from the positive x-axis.
  ///
  /// Precision is lost converting non-integral values to i32s.
  pub fn rotate(&self, degrees: f64) -> Point {
    let x = self.x as f64;
    let y = self.y as f64;
    let theta = degrees.to_radians();
    let x2 = (x * theta.cos()) - (y * theta.sin());
    let y2 = (y * theta.cos()) + (x * theta.sin());
    Point::new(x2.round() as i32, y2.round() as i32)
  }
}

impl PartialEq for Point {
  fn eq(&self, other: &Self) -> bool {
    self.x == other.x && self.y == other.y
  }
}

impl Add for Point {
  type Output = Self;
  fn add(self, other: Self) -> Self {
    Self {
      x: self.x + other.x,
      y: self.y + other.y,
    }
  }
}

impl AddAssign for Point {
  fn add_assign(&mut self, other: Self) {
    *self = Self {
      x: self.x + other.x,
      y: self.y + other.y,
    };
  }
}

impl Sub for Point {
  type Output = Self;
  fn sub(self, other: Self) -> Self::Output {
    Self {
      x: self.x - other.x,
      y: self.y - other.y,
    }
  }
}

impl Mul<i32> for Point {
  type Output = Self;

  fn mul(self, rhs: i32) -> Self::Output {
    Self {
      x: self.x * rhs,
      y: self.y * rhs,
    }
  }
}

impl fmt::Display for Point {
  fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
    write!(f, "({}, {})", self.x, self.y)
  }
}

#[derive(Clone)]
pub struct Grid {
  pub width: i32,
  pub height: i32,
  data: Vec<u8>,
}

impl Grid {
  /// Create a new grid filled with the seed value.
  pub fn new(width: i32, height: i32, seed: u8) -> Self {
    Grid {
      width,
      height,
      data: vec![seed; (width * height) as usize],
    }
  }

  /// Read a grid from a test fixture.
  pub fn from_fixture(name: &str) -> Self {
    Grid::from(Fixture::open(name))
  }

  /// An iterator visiting all Points in the grid, scanning the X axis first and
  /// the the Y axis, in order. The iterator element type is (Point, &'a u8).
  pub fn iter(&self) -> Iter {
    Iter { grid: self, i: 0 }
  }

  /// Returns `true` if the specified point is within the bounds of the grid.
  pub fn contains(&self, p: Point) -> bool {
    self.index_of(p).is_some()
  }

  /// Returns a the value corresponding to the specified point in the grid.
  pub fn get(&self, p: Point) -> Option<u8> {
    match self.index_of(p) {
      Some(i) => Some(self.data[i]),
      None => None,
    }
  }

  /// Returns the count of all points in the grid with the specified value.
  pub fn count(&self, v: u8) -> i32 {
    let mut n: i32 = 0;
    for a in self.data.iter() {
      if *a == v {
        n += 1;
      }
    }
    n
  }

  /// Converts a Point to an index in the grids underlying data structure.
  /// Returns None if the point is out of bounds.
  fn index_of(&self, p: Point) -> Option<usize> {
    if p.x < 0 || p.x >= self.width || p.y < 0 || p.y >= self.height {
      return None;
    }
    Some(((p.y * self.width) + p.x) as usize)
  }
}

impl From<&[u8]> for Grid {
  fn from(bytes: &[u8]) -> Self {
    let mut grid = Self {
      data: Vec::new(),
      width: 0,
      height: 0,
    };
    let mut i = 0;
    for b in bytes {
      if *b == b'\n' {
        if grid.width == 0 {
          grid.width = i;
        }
        grid.height += 1;
      } else {
        grid.data.push(*b);
        i += 1;
      }
    }
    grid
  }
}

impl From<Fixture> for Grid {
  fn from(fixture: Fixture) -> Self {
    Grid::from(&fixture.data[..])
  }
}

impl Index<Point> for Grid {
  type Output = u8;

  fn index(&self, p: Point) -> &Self::Output {
    match self.index_of(p) {
      Some(i) => &self.data[i],
      None => {
        panic!(format!("position out of bounds: {}", p));
      }
    }
  }
}

impl IndexMut<Point> for Grid {
  fn index_mut(&mut self, p: Point) -> &mut Self::Output {
    match self.index_of(p) {
      Some(i) => &mut self.data[i],
      None => {
        panic!(format!("position out of bounds: {}", p));
      }
    }
  }
}

impl fmt::Display for Grid {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    for i in 0..self.data.len() {
      f.write_char(self.data[i] as char)?;
      if i > 0 && (i + 1) % (self.width as usize) == 0 {
        f.write_char('\n')?;
      }
    }
    Ok(())
  }
}

pub struct Iter<'a> {
  grid: &'a Grid,
  i: i32,
}

impl Iterator for Iter<'_> {
  type Item = (Point, u8);

  fn next(&mut self) -> Option<Self::Item> {
    if self.i as usize >= self.grid.data.len() {
      return None;
    }
    let p = Point::new(self.i % self.grid.width, self.i / self.grid.width);
    self.i += 1;
    Some((p, self.grid[p]))
  }
}

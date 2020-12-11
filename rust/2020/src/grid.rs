use std::fmt;
use std::fmt::Write;
use std::iter::{Cloned, Iterator};
use std::ops::{Add, Index, IndexMut, Sub};

use crate::fixtures::Fixture;

#[derive(Copy, Clone)]
pub struct Point {
  pub x: i32,
  pub y: i32,
}

static ALL_DIRECTIONS: [Point; 8] = [
  Point { x: -1, y: -1 },
  Point { x: -1, y: 0 },
  Point { x: -1, y: 1 },
  Point { x: 0, y: -1 },
  Point { x: 0, y: 1 },
  Point { x: 1, y: -1 },
  Point { x: 1, y: 0 },
  Point { x: 1, y: 1 },
];

impl Point {
  pub fn new(x: i32, y: i32) -> Point {
    Point { x: x, y: y }
  }

  pub fn all_directions() -> Cloned<std::slice::Iter<'static, Point>> {
    ALL_DIRECTIONS.iter().cloned()
  }

  pub fn zero() -> Point {
    Self::new(0, 0)
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

impl Sub for Point {
  type Output = Self;
  fn sub(self, other: Self) -> Self::Output {
    Self {
      x: self.x - other.x,
      y: self.y - other.y,
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
  pub fn new(width: i32, height: i32, seed: u8) -> Self {
    Grid {
      width,
      height,
      data: vec![seed; (width * height) as usize],
    }
  }

  pub fn from_fixture(name: &str) -> Self {
    Grid::from(Fixture::open(name))
  }

  pub fn iter(&self) -> Iter {
    Iter { grid: self, i: 0 }
  }

  pub fn contains(&self, p: Point) -> bool {
    self.index_of(p).is_some()
  }

  pub fn get(&self, p: Point) -> Option<u8> {
    match self.index_of(p) {
      Some(i) => Some(self.data[i]),
      None => None,
    }
  }

  pub fn count(&self, v: u8) -> i32 {
    let mut n: i32 = 0;
    for a in self.data.iter() {
      if *a == v {
        n += 1;
      }
    }
    n
  }

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

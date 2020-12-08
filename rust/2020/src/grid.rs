use std::fmt;
use std::fmt::Write;
use std::ops::{Add, Index, IndexMut, Sub};

use crate::fixtures::Fixture;


#[derive(Copy, Clone)]
pub struct Point {
  pub x: i32,
  pub y: i32,
}

impl Point {
  pub fn new(x: i32, y: i32) -> Point {
    Point { x: x, y: y }
  }

  pub fn zero() -> Point {
    Self::new(0, 0)
  }

  pub fn is_zero(&self) -> bool {
    self.x == 0 && self.y == 0
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

pub struct Grid {
  pub width: i32,
  pub height: i32,
  data: Vec<u8>,
}

impl Grid {
  pub fn from_fixture(name: &str) -> Self {
    Grid::from(Fixture::open(name))
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

  fn index_of(&self, p: Point) -> Option<usize> {
    if p.x < 0 || p.x >= self.width || p.y < 0 || p.y >= self.height {
      return None
    }
    Some(((p.y * self.width) + p.x) as usize)
  }
}

impl From<Fixture> for Grid {
  fn from(fixture: Fixture) -> Self {
    let mut grid = Self {
      data: Vec::new(),
      width: 0,
      height: 0,
    };
    let mut i = 0;
    for b in fixture.data.iter() {
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

impl Index<Point> for Grid {
  type Output = u8;

  fn index(&self, p: Point) -> &Self::Output {
    match self.index_of(p) {
      Some(i) => &self.data[i],
      None => {
          panic!(format!("position out of bounds: {}", p));
      },
    }
  }
}

impl IndexMut<Point> for Grid {
  fn index_mut(&mut self, p: Point) -> &mut Self::Output {
    match self.index_of(p) {
      Some(i) => &mut self.data[i],
      None => {
          panic!(format!("position out of bounds: {}", p));
      },
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
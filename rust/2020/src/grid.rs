use std::fmt;
use std::fmt::Write;
use std::io;
use std::ops::{Add, Index, IndexMut, Sub};

use crate::fixtures::Fixture;
use crate::ToDoError;

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

  pub fn is_zero(self) -> bool {
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
  pub fn from_vec(buf: Vec<u8>) -> Self {
    let mut grid = Self {
      data: Vec::new(),
      width: 0,
      height: 0,
    };
    let mut i = 0;
    for b in buf.iter() {
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

  pub fn from_fixture(name: &str) -> Self {
    Self::from_vec(Fixture::open(name).data.clone())
  }

  fn index_of(&self, p: Point) -> Result<usize, ToDoError> {
    if !self.contains(p) {
      return Err(ToDoError);
    }
    Ok(((p.y * self.width) + p.x) as usize)
  }

  pub fn contains(&self, p: Point) -> bool {
    p.x >= 0 && p.x < self.width && p.y >= 0 && p.y < self.height
  }
}

impl Index<Point> for Grid {
  type Output = u8;

  fn index(&self, p: Point) -> &Self::Output {
    if !self.contains(p) {
      panic!(format!("position out of bounds: {}", p));
    }
    let i = self.index_of(p).unwrap();
    &self.data[i]
  }
}

impl IndexMut<Point> for Grid {
  fn index_mut(&mut self, p: Point) -> &mut Self::Output {
    if !self.contains(p) {
      panic!(format!("position out of bounds: {}", p));
    }
    let i = self.index_of(p).unwrap();
    &mut self.data[i]
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

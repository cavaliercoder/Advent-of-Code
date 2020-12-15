use std::error::Error;
use std::fmt;
use std::io;

#[derive(Debug)]
pub struct ToDoError;

impl Error for ToDoError {
  fn source(&self) -> Option<&(dyn Error + 'static)> {
    None
  }
}

impl fmt::Display for ToDoError {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    "You have doomed Christmas".fmt(f)
  }
}

impl From<io::Error> for ToDoError {
  fn from(_: io::Error) -> Self {
    ToDoError {}
  }
}

mod fixtures;
mod grid;

mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod day06;
mod day07;
mod day08;
mod day09;
mod day10;
mod day11;
mod day12;
mod day13;
mod day14;
mod day15;

use std::fmt;

#[derive(Debug)]
pub struct ToDoError;

impl fmt::Display for ToDoError {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    "You have doomed Christmas".fmt(f)
  }
}

mod fixtures;
mod grid;

mod day01;
mod day02;
mod day03;
mod day04;

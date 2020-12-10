use std::fmt;
use std::io;

#[derive(Debug)]
pub struct ToDoError;

impl fmt::Display for ToDoError {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    "You have doomed Christmas".fmt(f)
  }
}

impl From<io::Error> for ToDoError {
    fn from(_: io::Error) -> Self {
        ToDoError{}
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

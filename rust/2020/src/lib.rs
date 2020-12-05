use std::fs::{File, read};
use std::fmt;
use std::io::{BufRead, BufReader};
use std::str::FromStr;


#[derive(Debug)]
pub struct ToDoError;

impl fmt::Display for ToDoError {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    "You have doomed Christmas".fmt(f)
  }
}

pub fn fixture(name: &str) -> String {
  format!("../../inputs/2020/{}.dat", name)
}

pub fn read_fixture(name: &str) -> Result<Vec<u8>, ToDoError> {
  read(fixture(name)).map_err(|_| ToDoError{})
}

pub fn parse_fixture<T>(name: &str) -> Result<Vec<T>, ToDoError>
where
  T: FromStr,
{
  let filename = fixture(name);
  let reader = BufReader::new(File::open(filename).expect("Failed to open fixture"));
  let mut values = Vec::new();
  for line in reader.lines() {
    values.push(
      line
        .expect("Failed to read line")
        .parse()
        .ok()
        .expect("Failed to parse line"),
    );
  }
  Ok(values)
}

mod grid;

mod day01;
mod day02;
mod day03;
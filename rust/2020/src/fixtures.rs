use std::fs::{self, File};
use std::io;
use std::io::{BufRead, BufReader, Lines};
use std::str::FromStr;

use crate::ToDoError;

pub fn fixture(name: &str) -> String {
  format!("../../inputs/2020/{}.dat", name)
}

pub fn open(name: &str) -> io::Result<File> {
  File::open(fixture(name))
}

pub fn read(name: &str) -> io::Result<Vec<u8>> {
  fs::read(fixture(name))
}

pub fn read_lines(name: &str) -> io::Result<Lines<BufReader<File>>>{
  let reader = BufReader::new(open(name)?);
  Ok(reader.lines())
}

pub fn parse<T>(name: &str) -> Result<Vec<T>, ToDoError>
where
  T: FromStr,
{
  // let reader = io::BufReader::new(open_fixture(name).expect("Failed to open fixture"));
  let mut values = Vec::new();
  for line in read_lines(name).map_err(|_| ToDoError)? {
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

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

pub fn read_lines(name: &str) -> io::Result<Lines<BufReader<File>>> {
  let reader = BufReader::new(open(name)?);
  Ok(reader.lines())
}

pub fn parse<T>(name: &str) -> Result<Vec<T>, ToDoError>
where
  T: FromStr,
{
  let mut values = Vec::new();
  for line in read_lines(name)? {
    values.push(line?.parse::<T>().map_err(|_| ToDoError)?);
  }
  Ok(values)
}

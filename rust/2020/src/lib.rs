use std::fs::File;
use std::fmt;
use std::io::{BufRead, BufReader};
use std::str::FromStr;


#[derive(Debug)]
pub struct AOCError;

impl fmt::Display for AOCError {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    "provided string was not a valid Password".fmt(f)
  }
}


pub fn fixture<T>(name: &str) -> Result<Vec<T>, AOCError>
where
  T: FromStr,
{
  let filename = format!("../../inputs/2020/{}.dat", name);
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

mod day01;
mod day02;

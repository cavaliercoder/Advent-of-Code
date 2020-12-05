#[cfg(test)]
mod tests {
  use std::collections::HashMap;
  use std::fs::File;
  use std::io::{self, BufReader, Lines};
  use lazy_static::lazy_static;
  use regex::Regex;

  use crate::{fixtures, ToDoError};

  type Passport = HashMap<String, String>;

  lazy_static! {
    static ref RE_XYR: Regex = Regex::new(r"^[12][90][0-9][0-9]$").unwrap();
    static ref RE_HGT: Regex = Regex::new(r"^(?P<value>\d+)(?P<units>cm|in)$").unwrap();
    static ref RE_HCL: Regex = Regex::new(r"^#[0-9a-f]{6}$").unwrap();
    static ref RE_ECL: Regex = Regex::new(r"^(amb|blu|brn|gry|grn|hzl|oth)$").unwrap();
    static ref RE_PID: Regex = Regex::new(r"^[0-9]{9}$").unwrap();
  }

  fn read_passport(lines: &mut Lines<BufReader<File>>) -> io::Result<Option<Passport>> {
    let mut passport = Passport::new();
    let mut did_read = false;
    for line in lines {
      did_read = true;
      let line = line?;
      if line.is_empty() {
        break;
      }
      let pairs = line.split(' ');
      for pair in pairs {
        let parts: Vec<&str> = pair.split(':').collect();
        assert_eq!(parts.len(), 2, "pair: \"{}\"", pair);
        passport.insert(parts[0].to_string(), parts[1].to_string());
      }
    }
    if did_read {
      return Ok(Some(passport));
    }
    Ok(None)
  }

  fn read_batch(name: &str) -> Result<Vec<Passport>, ToDoError> {
    let mut passports: Vec<Passport> = Vec::new();
    let mut lines = fixtures::read_lines(name)?;
    loop {
      let passport = read_passport(&mut lines)?;
      match passport {
        Some(passport) => {
          passports.push(passport);
        }
        None => {
          return Ok(passports);
        }
      }
    }
  }

  fn check_passport(passport: &Passport) -> bool {
    let fields: [&str; 7] = [
      "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
    ];
    for &field in fields.iter() {
      if !passport.contains_key(field) {
        return false;
      }
    }
    true
  }

  fn check_passport_strict(passport: &Passport) -> bool {
    if !check_passport(passport) {
      return false;
    }
    // check byr (birth year) in [1920, 2002]:
    let byr = &passport["byr"];
    if !RE_XYR.is_match(byr) {
      return false;
    }
    let byr: usize = byr.parse().unwrap_or(0);
    if byr < 1920 || byr > 2002 {
      return false;
    }

    // check iyr (Issue Year) in [2010, 2020]
    let iyr = &passport["iyr"];
    if !RE_XYR.is_match(iyr) {
      return false;
    }
    let iyr: usize = iyr.parse().unwrap_or(0);
    if iyr < 2010 || iyr > 2020 {
      return false;
    }

    // check eyr (Expiration Year) in [2020, 2030]
    let eyr = &passport["eyr"];
    if !RE_XYR.is_match(eyr) {
      return false;
    }
    let eyr: usize = eyr.parse().unwrap_or(0);
    if eyr < 2020 || eyr > 2030 {
      return false;
    }

    // check hgt (Height) in [150cm, 193cm] or [59in, 76in]
    let hgt = &passport["hgt"];
    match RE_HGT.captures(hgt) {
      Some(caps) => {
        let hgt: usize = caps["value"].parse().unwrap_or(0);
        let units = &caps["units"];
        if units == "cm" {
          if hgt < 150 || hgt > 193 {
            return false;
          }
        }
        if units == "in" {
          if hgt < 59 || hgt > 76 {
            return false;
          }
        }
      }
      None => {
        return false;
      }
    }

    // check hcl (Hair Color) is hex color
    let hcl = &passport["hcl"];
    if !RE_HCL.is_match(hcl) {
      return false;
    }

    // check ecl (Eye Color) in [amb, blu, brn, gry, grn, hzl, oth]
    let ecl = &passport["ecl"];
    if !RE_ECL.is_match(ecl) {
      return false;
    }

    // check pid (Passport ID) is a nine-digit number
    let pid = &passport["pid"];
    if !RE_PID.is_match(pid) {
      return false
    }

    true
  }

  #[test]
  fn test_part1() {
    let passports = read_batch("day04").unwrap();
    let mut n = 0;
    for passport in passports.iter() {
      if check_passport(&passport) {
        n += 1;
      }
    }
    assert_eq!(n, 260);
  }

  #[test]
  fn test_part2() {
    let passports = read_batch("day04").unwrap();
    let mut n = 0;
    for passport in passports.iter() {
      if check_passport_strict(&passport) {
        n += 1;
      }
    }
    assert_eq!(n, 153);
  }
}

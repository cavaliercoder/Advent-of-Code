#[cfg(test)]
mod tests {
  use std::error::Error;
  use std::fmt;
  use std::fs::File;
  use std::io::{BufRead, BufReader};
  use std::str::FromStr;

  /// CharPolicy describes the minimum and maximum number of times a character
  /// must appear in a password.
  struct CharPolicy {
    c: char,
    min: i32,
    max: i32,
  }

  /// PasswordEntry describes a password in the database and the policy that
  /// was enforced at the time it was created.
  struct PasswordEntry {
    password: String,
    policy: CharPolicy,
  }

  impl PasswordEntry {
    /// Check the password against its policy using the rules from the old sled
    /// rental place down the street.
    fn is_valid(&self) -> bool {
      let mut n = 0;
      for c in self.password.chars() {
        if c == self.policy.c {
          n += 1;
        }
      }
      n >= self.policy.min && n <= self.policy.max
    }

    /// Check the password against its policy using the correct rules from the
    /// Official Toboggan Corporate Authentication System.
    fn is_valid2(&self) -> bool {
      let mut n = 0;
      let c = self.password.chars().nth(self.policy.min as usize).unwrap();
      if c == self.policy.c {
        n += 1;
      }
      let c = self.password.chars().nth(self.policy.max as usize).unwrap();
      if c == self.policy.c {
        n += 1;
      }
      n == 1
    }
  }

  impl fmt::Display for PasswordEntry {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(
        f,
        "{}-{} {}: {}",
        self.policy.min, self.policy.max, self.policy.c, self.password
      )
    }
  }

  impl FromStr for PasswordEntry {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
      // format: `<min>-<max> <char>: <password>`
      let parser_error = "parser error";

      // read policy min
      let i = s.find('-').ok_or(parser_error)?;
      let min: i32 = s[0..i].parse()?;
      let s = &s[i + 1..];

      // read policy max
      let i = s.find(' ').ok_or(parser_error)?;
      let max: i32 = s[..i].parse()?;
      let s = &s[i + 1..];

      // read policy char
      let c = s.chars().next().ok_or(parser_error)?;

      // read password
      let password = &s[2..];

      Ok(PasswordEntry {
        password: String::from(password),
        policy: CharPolicy { c, min, max },
      })
    }
  }

  fn get_password_db() -> Vec<PasswordEntry> {
    let fixture = "../../inputs/2020/day02.dat";
    let reader = BufReader::new(File::open(fixture).unwrap());
    let mut report: Vec<PasswordEntry> = Vec::new();
    for line in reader.lines() {
      report.push(line.unwrap().parse().unwrap());
    }
    report
  }

  #[test]
  fn test_part1() {
    let report = get_password_db();
    let mut valid = 0;
    for item in report.iter() {
      if item.is_valid() {
        valid += 1;
      }
    }
    assert_eq!(645, valid);
  }

  #[test]
  fn test_part2() {
    let report = get_password_db();
    let mut valid = 0;
    for item in report.iter() {
      if item.is_valid2() {
        valid += 1;
      }
    }
    assert_eq!(737, valid);
  }
}

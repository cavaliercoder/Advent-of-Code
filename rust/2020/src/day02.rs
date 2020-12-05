#[cfg(test)]
mod tests {
  use std::fmt;
  use std::str;

  use crate::{fixture, AOCError};

  /// CharPolicy describes the minimum and maximum number of times a character
  /// must appear in a password.
  struct CharPolicy {
    c: u8,
    min: usize,
    max: usize,
  }

  /// Password describes a password in the database and the policy that
  /// was enforced at the time it was created.
  struct Password {
    password: Vec<u8>,
    policy: CharPolicy,
  }

  impl Password {
    /// Check the password against its policy using the rules from the old sled
    /// rental place down the street.
    fn is_valid(&self) -> bool {
      let mut n = 0;
      for c in self.password.iter() {
        if *c == self.policy.c {
          n += 1;
        }
      }
      n >= self.policy.min && n <= self.policy.max
    }

    /// Check the password against its policy using the correct rules from the
    /// Official Toboggan Corporate Authentication System.
    fn is_valid2(&self) -> bool {
      (self.password[self.policy.min] == self.policy.c)
        ^ (self.password[self.policy.max] == self.policy.c)
    }
  }

  impl fmt::Display for Password {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
      write!(
        f,
        "{}-{} {}: {}",
        self.policy.min,
        self.policy.max,
        self.policy.c,
        str::from_utf8(&self.password).unwrap(),
      )
    }
  }


  impl str::FromStr for Password {
    type Err = AOCError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
      // format: `<min>-<max> <char>: <password>`

      // read policy min
      let i = s.find('-').ok_or(AOCError{})?;
      let min: usize = s[0..i].parse().map_err(|_| AOCError{})?;
      let s = &s[i + 1..];

      // read policy max
      let i = s.find(' ').ok_or(AOCError{})?;
      let max: usize = s[..i].parse().map_err(|_| AOCError{})?;
      let s = &s[i + 1..];

      // read policy char
      let c = s.chars().next().ok_or(AOCError{})? as u8;

      // read password
      let password = &s[2..];

      Ok(Self {
        password: password.as_bytes().to_vec(),
        policy: CharPolicy { c, min, max },
      })
    }
  }

  #[test]
  fn test_part1() {
    let db: Vec<Password> = fixture("day02").unwrap();
    let mut valid = 0;
    for item in db.iter() {
      if item.is_valid() {
        valid += 1;
      }
    }
    assert_eq!(645, valid);
  }

  #[test]
  fn test_part2() {
    let db: Vec<Password> = fixture("day02").unwrap();
    let mut valid = 0;
    for item in db.iter() {
      if item.is_valid2() {
        valid += 1;
      }
    }
    assert_eq!(737, valid);
  }
}

#[cfg(test)]
mod tests {
    use std::collections::HashSet;

    use lazy_static::lazy_static;
    use regex::Regex;

    use crate::fixtures;

    lazy_static! {
        static ref RE_PASS: Regex = Regex::new(r"^[FB]{7}[LR]{3}$").unwrap();
    }

    fn get_seat_id(s: &str) -> i32 {
        assert!(RE_PASS.is_match(s));
        let s = s.as_bytes();
        let mut row: i32 = 0;
        let mut column: i32 = 0;
        let mut lower: i32 = 0;
        let mut upper: i32 = 127;
        for i in 0..7 {
            row = lower + (upper - lower) / 2;
            match s[i] {
                b'F' => {
                    upper = row;
                }
                b'B' => {
                    lower = row;
                }
                _ => {
                    unreachable!();
                }
            }
        }
        if s[6] == b'B' {
            row += 1;
        }
        lower = 0;
        upper = 7;
        for i in 7..10 {
            column = lower + (upper - lower) / 2;
            match s[i] {
                b'L' => {
                    upper = column;
                }
                b'R' => {
                    lower = column;
                }
                _ => {
                    unreachable!();
                }
            }
        }
        if s[9] == b'R' {
            column += 1;
        }
        (row * 8) + column
    }

    #[test]
    fn test_part1_example1() {
        assert_eq!(get_seat_id("BFFFBBFRRR"), 567);
    }
    #[test]
    fn test_part1_example2() {
        assert_eq!(get_seat_id("FFFBBBFRRR"), 119);
    }

    #[test]
    fn test_part1_example3() {
        assert_eq!(get_seat_id("BBFFBBFRLL"), 820);
    }

    #[test]
    fn test_part1() {
        let passes: Vec<String> = fixtures::parse("day05").unwrap();
        let mut max_pass = 0;
        for pass in passes.iter() {
            let pass_id = get_seat_id(pass);
            if pass_id > max_pass {
                max_pass = pass_id;
            }
        }
        assert_eq!(max_pass, 933);
    }

    #[test]
    fn test_part2() {
        let passes: Vec<String> = fixtures::parse("day05").unwrap();
        let mut seen: HashSet<i32> = HashSet::new();
        let mut min_pass: i32 = i32::MAX;
        let mut max_pass: i32 = 0;
        for pass in passes.iter() {
            let pass_id = get_seat_id(pass);
            seen.insert(pass_id);
            if pass_id < min_pass {
                min_pass = pass_id;
            }
            if pass_id > max_pass {
                max_pass = pass_id;
            }
        }
        for i in min_pass..max_pass {
            if !seen.contains(&i) {
                assert_eq!(i, 0);
                break;
            }
        }
    }
}

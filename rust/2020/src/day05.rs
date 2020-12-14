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
        let mut row_lower: i32 = 0;
        let mut row_upper: i32 = 128;
        let mut col_lower: i32 = 0;
        let mut col_upper: i32 = 8;
        for c in s.chars() {
            match c {
                'F' => {
                    row_upper = (row_lower + row_upper) / 2;
                }
                'B' => {
                    row_lower = (row_lower + row_upper) / 2;
                },
                'L' => {
                    col_upper = (col_lower + col_upper) / 2;
                }
                'R' => {
                    col_lower = (col_lower + col_upper) / 2;
                }
                _ => {
                    unreachable!();
                }
            }
        }
        (row_lower * 8) + col_lower
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
        let passes = fixtures::Fixture::open("day05");
        let mut max_pass = 0;
        for pass in passes.iter() {
            let pass_id = get_seat_id(&pass);
            if pass_id > max_pass {
                max_pass = pass_id;
            }
        }
        assert_eq!(max_pass, 933);
    }

    #[test]
    fn test_part2() {
        let passes = fixtures::Fixture::open("day05");
        let mut seen: HashSet<i32> = HashSet::new();
        let mut min_pass: i32 = i32::MAX;
        let mut max_pass: i32 = 0;
        for pass in passes.iter() {
            let pass_id = get_seat_id(&pass);
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
                assert_eq!(i, 711);
                break;
            }
        }
    }
}

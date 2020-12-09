#[cfg(test)]
mod tests {
    use std::cmp::Ordering;
    use std::collections::HashSet;
    use std::iter::{FromIterator, Iterator};
    use std::ops::Range;

    use crate::fixtures::Fixture;

    fn is_valid(n: i64, preamble: &[i64]) -> bool {
        let preamble: HashSet<i64> = HashSet::from_iter(preamble.iter().cloned());
        for m in preamble.iter() {
            if preamble.contains(&(n - *m)) {
                return true;
            }
        }
        false
    }

    fn find_first_invalid(values: &[i64], preamble_len: usize) -> Option<i64> {
        for i in preamble_len..values.len() {
            let n = &values[i];
            if !is_valid(*n, &values[i - preamble_len..i]) {
                return Some(*n);
            }
        }
        None
    }

    fn find_weakness(values: &Vec<i64>, needle: i64) -> Option<i64> {
        let mut sum = values[0] + values[1];
        let r: &mut Range<usize> = &mut Range { start: 0, end: 2 };
        for _ in 0..values.len() - 1 {
            match sum.cmp(&needle) {
                Ordering::Less => {
                    // expand end of range until >= needle
                    while sum < needle {
                        sum += values[r.end];
                        r.end += 1;
                    }
                }
                Ordering::Greater => {
                    // contract range until <= needle
                    while sum > needle {
                        sum -= values[r.end - 1];
                        r.end -= 1;
                    }
                }
                _ => (),
            }
            if sum == needle {
                // answer is sum of in and max in values[r]
                let min = values[r.clone()].iter().min().unwrap_or(&0);
                let max = values[r.clone()].iter().max().unwrap_or(&0);
                return Some(min + max);
            }
            // shift start of range forward
            sum -= values[r.start];
            r.start += 1;
        }
        None
    }

    #[test]
    fn test_part1_example1() {
        let values: Vec<i64> = vec![
            35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309,
            576,
        ];
        assert_eq!(find_first_invalid(&values, 5), Some(127));
    }

    #[test]
    fn test_part1() {
        let values: Vec<i64> = Fixture::open("day09").parse().unwrap();
        assert_eq!(find_first_invalid(&values, 25), Some(105950735));
    }

    #[test]
    fn test_part2_example1() {
        let values: Vec<i64> = vec![
            35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309,
            576,
        ];
        assert_eq!(find_weakness(&values, 127), Some(62));
    }

    #[test]
    fn test_part2() {
        let values: Vec<i64> = Fixture::open("day09").parse().unwrap();
        assert_eq!(find_weakness(&values, 105950735), Some(13826915));
    }
}

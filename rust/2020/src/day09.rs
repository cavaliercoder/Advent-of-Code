#[cfg(test)]
mod tests {
    use std::collections::HashSet;
    use std::iter::{FromIterator, Iterator};

    use crate::fixtures::Fixture;

    fn find_first_invalid(values: &[i64], preamble_len: usize) -> Option<i64> {
        let mut preamble: HashSet<i64> =
            HashSet::from_iter(values[0..preamble_len].iter().cloned());
        for i in preamble_len..values.len() {
            let n = &values[i];
            let mut valid = false;
            for m in &preamble {
                if preamble.contains(&(n - *m)) {
                    valid = true;
                    break;
                }
            }
            if !valid {
                return Some(*n);
            }
            preamble.remove(&values[i-preamble_len]);
            preamble.insert(values[i]);
        }
        None
    }

    fn find_weakness(values: &Vec<i64>, needle: i64) -> Option<i64> {
        let mut sum = values[0];
        let mut j: usize = 1;
        for i in 0..values.len() - 1 {
            while j < values.len() && sum < needle {
                sum += values[j];
                j += 1;
            }
            if sum == needle {
                let values = &values[i..j];
                return values
                    .iter()
                    .min()
                    .and_then(|min| values.iter().max().map(|max| min + max));
            }
            sum -= values[i];
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

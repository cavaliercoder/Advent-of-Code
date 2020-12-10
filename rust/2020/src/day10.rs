#[cfg(test)]
mod tests {
    use std::collections::HashSet;
    use std::iter::FromIterator;

    use crate::fixtures::Fixture;

    fn get_joltage_distribution(adapters: &Vec<i64>) -> i64 {
        let mut diffs = (0i64, 1i64); // +1 for built-in adapter at the end
        let mut adapters: HashSet<i64> = HashSet::from_iter(adapters.iter().cloned());
        adapters.insert(0); //  outlet
        for adapter in adapters.iter().cloned() {
            if adapters.contains(&(adapter + 1)) {
                diffs.0 += 1;
            } else if adapters.contains(&(adapter + 3)) {
                diffs.1 += 1;
            }
        }
        diffs.0 * diffs.1
    }

    fn count_paths(adapters: &Vec<i64>) -> i64 {
        let mut adapters = adapters.to_vec();
        adapters.push(0);
        adapters.sort();
        let mut paths = vec![0i64; adapters.len()];
        paths[0] = 1;
        for i in 0..adapters.len() {
            for j in (i + 1)..adapters.len() {
                let diff = adapters[j] - adapters[i];
                if diff > 3 {
                    break;
                }
                paths[j] += paths[i];
            }
        }
        paths[paths.len() - 1]
    }

    #[test]
    fn test_part1_example1() {
        let adapters = vec![16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4];
        assert_eq!(get_joltage_distribution(&adapters), 35);
    }

    #[test]
    fn test_part1_example2() {
        let adapters = vec![
            28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35,
            8, 17, 7, 9, 4, 2, 34, 10, 3,
        ];
        assert_eq!(get_joltage_distribution(&adapters), 220);
    }

    #[test]
    fn test_part1() {
        let adapters: Vec<i64> = Fixture::open("day10").parse().unwrap();
        assert_eq!(get_joltage_distribution(&adapters), 2263);
    }

    #[test]
    fn test_part2_example1() {
        let adapters = vec![16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4];
        assert_eq!(count_paths(&adapters), 8);
    }

    #[test]
    fn test_part2_example2() {
        let adapters = vec![
            28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35,
            8, 17, 7, 9, 4, 2, 34, 10, 3,
        ];
        assert_eq!(count_paths(&adapters), 19208);
    }

    #[test]
    fn test_part2() {
        let adapters: Vec<i64> = Fixture::open("day10").parse().unwrap();
        assert_eq!(count_paths(&adapters), 396857386627072);
    }
}

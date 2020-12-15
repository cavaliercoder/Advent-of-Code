#[cfg(test)]
mod tests {
    use std::collections::HashMap;

    fn brute_force(start: &Vec<i64>, turns: i64) -> i64 {
        let mut seen: HashMap<i64, i64> = HashMap::new();
        let mut last: i64 = 0;
        let mut next: i64 = 0;
        for i in 0i64..turns {
            if i < start.len() as i64 {
                next = start[i as usize];
            } else if let Some(&mut turn) = seen.get_mut(&next) {
                next = i - turn;
            } else {
                next = 0;
            }
            if i > 0 {
                seen.insert(last, i);
            }
            last = next;
        }
        last
    }

    #[test]
    fn test_part1_examples() {
        assert_eq!(brute_force(&vec![0, 3, 6], 2020), 436);
        assert_eq!(brute_force(&vec![1, 3, 2], 2020), 1);
        assert_eq!(brute_force(&vec![2, 1, 3], 2020), 10);
        assert_eq!(brute_force(&vec![1, 2, 3], 2020), 27);
        assert_eq!(brute_force(&vec![2, 3, 1], 2020), 78);
        assert_eq!(brute_force(&vec![3, 2, 1], 2020), 438);
        assert_eq!(brute_force(&vec![3, 1, 2], 2020), 1836);
    }

    #[test]
    fn test_part1() {
        let fixture = vec![20, 0, 1, 11, 6, 3];
        assert_eq!(brute_force(&fixture, 2020), 421);
    }

    // #[test]
    // fn test_part2_examples() {
    //     assert_eq!(brute_force(&vec![0,3,6], 30_000_000), 175594);
    //     assert_eq!(brute_force(&vec![1,3,2], 30_000_000), 2578);
    //     assert_eq!(brute_force(&vec![2,1,3], 30_000_000), 3544142);
    //     assert_eq!(brute_force(&vec![1,2,3], 30_000_000), 261214);
    //     assert_eq!(brute_force(&vec![2,3,1], 30_000_000), 6895259);
    //     assert_eq!(brute_force(&vec![3,2,1], 30_000_000), 18);
    //     assert_eq!(brute_force(&vec![3,1,2], 30_000_000), 362);
    // }

    #[test]
    fn test_part2() {
        assert_eq!(brute_force(&vec![20, 0, 1, 11, 6, 3], 30_000_000), 436);
    }
}

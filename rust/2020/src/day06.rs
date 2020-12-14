#[cfg(test)]
mod tests {
    use crate::fixtures::Fixture;

    #[test]
    fn test_part1() {
        let mut sum = 0;
        let mut group_seen: u32 = 0;
        for line in Fixture::open("day06").iter() {
            if line.is_empty() {
                sum += group_seen.count_ones();
                group_seen = 0;
                continue;
            }
            for c in line.as_bytes().iter() {
                group_seen |= 1 << (c - b'a');
            }
        }
        sum += group_seen.count_ones();
        assert_eq!(sum, 6530);
    }

    #[test]
    fn test_part2() {
        let mut sum = 0;
        let mut group_seen: u32 = 0xFFFFFFFF;
        for line in Fixture::open("day06").iter() {
            if line.is_empty() {
                sum += group_seen.count_ones();
                group_seen = 0xFFFFFFFF;
                continue;
            }
            let mut form: u32 = 0;
            for c in line.as_bytes().iter() {
                form |= 1 << (c - b'a')
            }
            group_seen &= form;
        }
        sum += group_seen.count_ones();
        assert_eq!(sum, 3323);
    }
}
#[cfg(test)]
mod tests {
    use std::collections::HashSet;
    use std::iter::FromIterator;

    use crate::fixtures::Fixture;

    const SUM: i32 = 2020;

    #[test]
    fn test_part1() {
        let expect: i32 = 299299;
        let report: Vec<i32> = Fixture::open("day01").parse().unwrap();
        let report: HashSet<&i32> = HashSet::from_iter(report.iter());
        for &n in report.iter() {
            let m = SUM - n;
            if report.contains(&m) {
                assert_eq!(expect, n * m);
                return;
            }
        }
        panic!("no result");
    }

    #[test]
    fn test_part2() {
        // TODO: Can we do this in better than O(n³)?
        let expect: i32 = 287730716;
        let report: Vec<i32> = Fixture::open("day01").parse().unwrap();
        for i in 0..report.len() - 1 {
            for j in 1..report.len() {
                let i = report[i];
                let j = report[j];
                let k = SUM - i - j;
                if report.contains(&k) {
                    assert_eq!(expect, i * j * k);
                    return;
                }
            }
        }
        panic!("no result");
    }
}

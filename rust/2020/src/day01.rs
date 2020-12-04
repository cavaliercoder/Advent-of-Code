#[cfg(test)]
mod tests {
    use std::collections::HashSet;
    use std::fs::File;
    use std::io::{BufRead, BufReader};
    use std::iter::FromIterator;

    const SUM: i32 = 2020;

    fn get_expense_report() -> Vec<i32> {
        let fixture = "../../inputs/2020/day01.dat";
        let reader = BufReader::new(File::open(fixture).unwrap());
        let mut report: Vec<i32> = Vec::new();
        for line in reader.lines() {
            report.push(line.unwrap().parse().unwrap());
        }
        report
    }

    #[test]
    fn test_part1() {
        let expect: i32 = 299299;
        let report = get_expense_report();
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
        // TODO: Can we do this in better than O(n²)?
        let expect: i32 = 287730716;
        let report = get_expense_report();
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

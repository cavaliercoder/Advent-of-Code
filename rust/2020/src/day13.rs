#[cfg(test)]
mod tests {
    use crate::fixtures::Fixture;

    fn next_bus(bus_id: i32, timestamp: i32) -> i32 {
        (timestamp / bus_id + 1) * bus_id - timestamp
    }

    #[test]
    fn test_part1() {
        let fixture = Fixture::open("day13");
        let mut iter = fixture.iter();
        let timestamp: i32 = iter.next().unwrap().parse().unwrap();
        let buses: Vec<i32> = iter
            .next()
            .unwrap()
            .split(',')
            .filter_map(|s| {
                if s == "x" {
                    // bus out of service
                    return None;
                }
                Some(s.parse::<i32>().unwrap())
            })
            .collect();
        let mut min_bus = i32::MAX;
        let mut min_wait = i32::MAX;
        for bus in buses {
            let wait = next_bus(bus, timestamp);
            if wait < min_wait {
                min_bus = bus;
                min_wait = wait;
            }
        }
        assert_eq!(min_bus * min_wait, 5946);
    }

    /// Check if n is a prime number using the 6𝑘± optimization.
    fn is_prime(n: i64) -> bool {
        if n <= 3 {
            return n > 1;
        }
        if n % 2 == 0 || n % 3 == 0 {
            return false;
        }
        let mut i: i64 = 5;
        while i.pow(2) <= n {
            if n % i == 0 || n % (i + 2) == 0 {
                return false;
            }
            i += 6;
        }
        true
    }

    /// Get buses as tuple of (timestamp offset, bus ID)
    fn parse_buses(s: &str) -> Vec<(i64, i64)> {
        let mut i: i64 = 0;
        s.split(',')
            .filter_map(|s| {
                i += 1;
                if s != "x" {
                    Some((i - 1, s.parse::<i64>().unwrap()))
                } else {
                    None // bus out of service
                }
            })
            .collect()
    }

    /// Find Tn where all buses are in sync.
    ///
    /// This solutions exploits the fact that all bus IDs are primes so we
    /// can avoid using Euclid's lemma and Extended GCD algorithm to find
    /// earlier sync events and simply find some multiple of all bus IDs that
    /// satisfies their desired offsets.
    fn find_sync(buses: &Vec<(i64, i64)>) -> i64 {
        // bus ID must be a prime number for this to work
        for bus in buses {
            assert!(is_prime(bus.1))
        }
        let mut t: i64 = buses[0].1;
        let mut step: i64 = buses[0].1;
        for bus in buses[1..].iter() {
            for c in (t..).step_by(step as usize) {
                if (c + bus.0) % bus.1 == 0 {
                    t = c;
                    step = step * bus.1;
                    break;
                }
            }
        }
        t
    }

    #[test]
    fn test_part2_examples() {
        assert_eq!(find_sync(&parse_buses("7,13,x,x,59,x,31,19")), 1_068_781);
        assert_eq!(find_sync(&parse_buses("17,x,13,19")), 3_417);
        assert_eq!(find_sync(&parse_buses("67,7,59,61")), 754_018);
        assert_eq!(find_sync(&parse_buses("67,x,7,59,61")), 779_210);
        assert_eq!(find_sync(&parse_buses("67,7,x,59,61")), 1_261_476);
        assert_eq!(find_sync(&parse_buses("1789,37,47,1889")), 1_202_161_486);
    }

    #[test]
    fn test_part2() {
        let lines: Vec<String> = Fixture::open("day13").iter().collect();
        let buses = parse_buses(&lines[1]);
        assert_eq!(find_sync(&buses), 645_338_524_823_718);
    }
}

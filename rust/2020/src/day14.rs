#[cfg(test)]
mod tests {
    use std::collections::HashMap;
    use std::str::FromStr;

    use crate::fixtures::Fixture;
    use crate::ToDoError;

    #[derive(Clone, Copy, Debug)]
    enum Instruction {
        Mask(i64, i64), // mask, value
        Mem(i64, i64),  // address, value
    }

    fn parse_mask(s: &str) -> Result<Instruction, ToDoError> {
        let s = &s[7..]; // trim "mask = "
        let mut i: usize = 36; // value width - 1
        let mut mask: i64 = 0; // all X's
        let mut value: i64 = 0; // all 1s and 0s
        for c in s.chars() {
            i -= 1;
            match c {
                '0' => {}
                '1' => {
                    value |= 1 << i;
                }
                'X' => {
                    mask |= 1 << i;
                }
                _ => {
                    return Err(ToDoError {});
                }
            }
        }
        Ok(Instruction::Mask(mask, value))
    }

    fn parse_mem(s: &str) -> Result<Instruction, ToDoError> {
        let s = &s[4..]; // trim "mem["
        let parts: Vec<&str> = s.split("] = ").collect();
        if parts.len() != 2 {
            return Err(ToDoError {});
        }
        let addr: i64 = parts[0].parse().map_err(|_| ToDoError {})?;
        let value: i64 = parts[1].parse().map_err(|_| ToDoError {})?;
        Ok(Instruction::Mem(addr, value))
    }

    impl FromStr for Instruction {
        type Err = ToDoError;

        fn from_str(s: &str) -> Result<Self, Self::Err> {
            if s.starts_with("mask = ") {
                return parse_mask(s);
            }
            if s.starts_with("mem[") {
                return parse_mem(s);
            }
            return Err(ToDoError {});
        }
    }

    #[test]
    fn test_part1() {
        let instructions: Vec<Instruction> = Fixture::open("day14").parse().unwrap();
        let mut mask_mask: i64 = 0;
        let mut mask_value: i64 = 0;
        let mut mem: HashMap<i64, i64> = HashMap::new();
        let mut sum: i64 = 0;
        for instruction in instructions {
            match instruction {
                Instruction::Mask(mask, value) => {
                    mask_mask = mask;
                    mask_value = value;
                }
                Instruction::Mem(addr, value) => {
                    mem.insert(addr, (value & mask_mask) | mask_value);
                }
            }
        }
        for v in mem.values() {
            sum += v;
        }
        assert_eq!(sum, 7_440_382_076_205);
    }

    fn bit_on(v: i64, bit: usize) -> i64 {
        v | (1 << bit)
    }

    fn bit_off(v: i64, bit: usize) -> i64 {
        v & !(1 << bit)
    }

    fn is_bit_on(v: i64, bit: usize) -> bool {
        (1i64 << bit) & v != 0
    }

    fn get_permutations(addr: i64, mask: i64) -> Vec<i64> {
        let mut addresses: Vec<i64> = Vec::new();
        let mut stack = vec![(addr, mask)];
        while !stack.is_empty() {
            let (addr, mask) = stack.pop().unwrap();
            if mask == 0 { // no more X's
                addresses.push(addr);
                continue;
            }
            for i in 0..36 {
                if is_bit_on(mask, i) {
                    stack.push((bit_on(addr, i), bit_off(mask, i)));
                    stack.push((bit_off(addr, i), bit_off(mask, i)));
                    break;
                }
            }
        }
        addresses
    }

    #[test]
    fn test_part2() {
        let instructions: Vec<Instruction> = Fixture::open("day14").parse().unwrap();
        let mut mask_mask: i64 = 0;
        let mut mask_value: i64 = 0;
        let mut mem: HashMap<i64, i64> = HashMap::new();
        for instruction in instructions {
            match instruction {
                Instruction::Mask(mask, value) => {
                    mask_mask = mask;
                    mask_value = value;
                }
                Instruction::Mem(addr, value) => {
                    let addresses = get_permutations(
                        addr | mask_value, // overwrite 1s
                        mask_mask,
                    );
                    for addr in addresses {
                        mem.insert(addr, value);
                    }
                }
            }
        }
        let mut sum: i64 = 0;
        for v in mem.values() {
            sum += v;
        }
        assert_eq!(sum, 4_200_656_704_538);
    }
}

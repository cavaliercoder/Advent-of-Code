#[cfg(test)]
mod tests {
    use std::collections::HashSet;
    use std::str::FromStr;

    use crate::fixtures::Fixture;
    use crate::ToDoError;

    #[derive(Clone, Debug)]
    enum Instruction {
        ACC(i32),
        JMP(i32),
        NOP(i32),
    }

    impl FromStr for Instruction {
        type Err = ToDoError;

        fn from_str(s: &str) -> Result<Self, Self::Err> {
            let parts: Vec<&str> = s.split_whitespace().collect();
            if parts.len() != 2 {
                return Err(ToDoError);
            }
            let n: i32 = parts[1].parse().map_err(|_| ToDoError)?;
            match parts[0] {
                "acc" => Ok(Instruction::ACC(n)),
                "jmp" => Ok(Instruction::JMP(n)),
                "nop" => Ok(Instruction::NOP(n)),
                _ => Err(ToDoError),
            }
        }
    }

    /// Run a program until it loops or program counter goes out of range.
    /// Returns the accumulator value and true if the program terminated
    /// normally.
    fn run(data: &Vec<Instruction>) -> (i32, bool) {
        let mut pc: i32 = 0;
        let mut acc: i32 = 0;
        let mut seen: HashSet<i32> = HashSet::new();
        loop {
            seen.insert(pc);
            let mut next: i32 = pc + 1;
            match &data[pc as usize] {
                Instruction::ACC(n) => {
                    acc += n;
                }
                Instruction::JMP(n) => {
                    next = pc + n;
                }
                Instruction::NOP(_) => {}
            };
            if next == data.len() as i32 {
                // terminated correctly
                return (acc, true);
            }
            if next < 0 || next > data.len() as i32 || seen.contains(&next) {
                // out of range or looped
                return (acc, false);
            }
            pc = next;
        }
    }

    #[test]
    fn test_part1() {
        let data: Vec<Instruction> = Fixture::open("day08").parse().unwrap();
        let (acc, _) = run(&data);
        assert_eq!(acc, 1749);
    }

    #[test]
    fn test_part2() {
        let data: Vec<Instruction> = Fixture::open("day08").parse().unwrap();
        for i in (0..data.len()).rev() {
            if let Instruction::ACC(_) = &data[i] {
                continue;
            }
            let mut data2 = data.clone();
            data2[i] = match &data[i] {
                Instruction::JMP(n) => Instruction::NOP(*n),
                Instruction::NOP(n) => Instruction::JMP(*n),
                _ => {
                    unreachable!();
                }
            };
            let (acc, ok) = run(&data2);
            if ok {
                assert_eq!(acc, 515);
                return;
            }
        }
        unreachable!();
    }
}

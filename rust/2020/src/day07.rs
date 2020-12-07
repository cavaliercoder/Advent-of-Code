#[cfg(test)]
mod tests {
    use std::collections::HashMap;
    use std::str::FromStr;

    use crate::fixtures::Fixture;
    use crate::ToDoError;

    #[derive(Debug)]
    struct Rule {
        bag: String,
        requirements: HashMap<String, usize>,
    }

    impl FromStr for Rule {
        type Err = ToDoError;

        fn from_str(s: &str) -> Result<Self, Self::Err> {
            let words: Vec<&str> = s.split_whitespace().collect();
            let mut rule = Rule {
                bag: words[0..2].join(" "),
                requirements: HashMap::new(),
            };
            let mut i = 4; // first child bag
            while i < words.len() {
                if words[i] == "no" {
                    break; // no other bags
                }
                let n: usize = words[i].parse().unwrap();
                let bag: String = words[i + 1..i + 3].join(" ");
                rule.requirements.insert(bag, n);
                i += 4;
            }
            Ok(rule)
        }
    }


    fn bag_contains(rules: &HashMap<&String, &Rule>, bag: &String) -> bool {
        fn bag_contains_recursive(
            rules: &HashMap<&String, &Rule>,
            seen: &mut HashMap<String, bool>,
            bag: &String,
        ) -> bool {
            if let Some(ok) = seen.get(bag) {
                return *ok;
            }
            let rule = &rules[bag];
            let mut result = false;
            if rule.requirements.contains_key("shiny gold") {
                result = true
            } else {
                for requirement in rule.requirements.keys() {
                    if bag_contains_recursive(rules, seen, requirement) {
                        result = true;
                        break;
                    }
                }
            }
            seen.insert(rule.bag.to_string(), result);
            result
        }

        let mut seen: HashMap<String, bool> = HashMap::new();
        bag_contains_recursive(rules, &mut seen, bag)
    }

    fn bag_count(rules: &HashMap<&String, &Rule>, bag: &String) -> usize {
        fn bag_count_recursive(
            rules: &HashMap<&String, &Rule>,
            seen: &mut HashMap<String, usize>,
            bag: &String,
        ) -> usize {
            if let Some(n) = seen.get(bag) {
                return *n;
            }
            let mut n: usize = 1;
            let rule = &rules[bag];
            for (child, count) in rule.requirements.iter() {
                n += count * bag_count_recursive(rules, seen, child);
            }
            seen.insert(bag.to_string(), n);
            n
        };

        let mut seen: HashMap<String, usize> = HashMap::new();
        bag_count_recursive(rules, &mut seen, bag) - 1 // -1 for top bag
    }

    #[test]
    fn test_part1() {
        let rules: Vec<Rule> = Fixture::open("day07").parse().unwrap();
        let mut rules_by_bag: HashMap<&String, &Rule> = HashMap::new();
        for rule in rules.iter() {
            rules_by_bag.insert(&rule.bag, &rule);
        }
        let mut n = 0;
        for rule in rules.iter() {
            if bag_contains(&rules_by_bag, &rule.bag) {
                n += 1;
            }
        }
        assert_eq!(n, 238);
    }

    #[test]
    fn test_part2() {
        let rules: Vec<Rule> = Fixture::open("day07").parse().unwrap();
        let mut rules_by_bag: HashMap<&String, &Rule> = HashMap::new();
        for rule in rules.iter() {
            rules_by_bag.insert(&rule.bag, &rule);
        }
        assert_eq!(
            bag_count(&rules_by_bag, &"shiny gold".to_string()),
            82930
        );
    }
}

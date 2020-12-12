#[cfg(test)]
mod tests {
    use crate::fixtures::Fixture;
    use crate::grid::{Point, NORTH, SOUTH, EAST, WEST};
    use crate::ToDoError;

    #[derive(Debug)]
    struct Ship {
        position: Point,
        waypoint: Point,
        orientation: Point,
    }

    impl Ship {
        fn new() -> Ship {
            Ship {
                position: Point::zero(),
                waypoint: Point::new(10, -1), // 10 east, 1 north
                orientation: EAST,
            }
        }

        fn dispatch(&mut self, action: &str) -> Result<(), ToDoError> {
            let verb = &action[..1];
            let param: i32 = action[1..].parse().map_err(|_| ToDoError)?;
            match verb {
                "N" => self.position += NORTH * param,
                "S" => self.position += SOUTH * param,
                "E" => self.position += EAST * param,
                "W" => self.position += WEST * param,
                "L" => {
                    self.orientation = self.orientation.rotate(-param as f64);
                }
                "R" => {
                    self.orientation = self.orientation.rotate(param as f64);
                }
                "F" => self.position += self.orientation * param,
                _ => return Err(ToDoError),
            }
            Ok(())
        }

        fn dispatch_with_waypoint(&mut self, action: &str) -> Result<(), ToDoError> {
            let verb = &action[..1];
            let param: i32 = action[1..].parse().map_err(|_| ToDoError)?;
            match verb {
                "N" => self.waypoint += NORTH * param,
                "S" => self.waypoint += SOUTH * param,
                "E" => self.waypoint += EAST * param,
                "W" => self.waypoint += WEST * param,
                "L" => {
                    self.waypoint = self.waypoint.rotate(-param as f64);
                }
                "R" => {
                    self.waypoint = self.waypoint.rotate(param as f64);
                }
                "F" => self.position += self.waypoint * param,
                _ => return Err(ToDoError),
            }
            Ok(())
        }
    }

    #[test]
    fn test_part1() {
        let actions: Vec<String> = Fixture::open("day12").parse().unwrap();
        let mut ship = Ship::new();
        for action in actions {
            ship.dispatch(&action).unwrap();
        }
        assert_eq!(ship.position.manhattan_distance(), 1645);
    }

    #[test]
    fn test_part2_example1() {
        let actions = vec!["F10", "N3", "F7", "R90", "F11"];
        let mut ship = Ship::new();
        for action in actions {
            ship.dispatch_with_waypoint(action).unwrap();
        }
        assert_eq!(ship.position.manhattan_distance(), 286);
    }

    #[test]
    fn test_part2() {
        let actions: Vec<String> = Fixture::open("day12").parse().unwrap();
        let mut ship = Ship::new();
        for action in actions {
            ship.dispatch_with_waypoint(&action).unwrap();
        }
        assert_eq!(ship.position.manhattan_distance(), 35292);
    }
}

#include "lib/aoc.h"

namespace aoc2022 {

class Day21 {
 public:
  int Part1(aoc::Input in) { return 21; }
  int Part2(aoc::Input in) { return 21; };
};

TEST(Day21, Part1) { EXPECT_EQ(Day21().Part1(aoc::Input(2022, 21)), 21); }
TEST(Day21, Part2) { EXPECT_EQ(Day21().Part2(aoc::Input(2022, 21)), 21); }

}  // namespace aoc2022

#include "lib/aoc.h"

namespace aoc2022 {

class Day25 {
 public:
  int Part1(aoc::Input in) { return 25; }
  int Part2(aoc::Input in) { return 25; };
};

TEST(Day25, Part1) { EXPECT_EQ(Day25().Part1(aoc::Input(2022, 25)), 25); }
TEST(Day25, Part2) { EXPECT_EQ(Day25().Part2(aoc::Input(2022, 25)), 25); }

}  // namespace aoc2022

#include <algorithm>

#include "lib/aoc.h"

namespace aoc2022 {

class Day06 {
  int index_marker(std::string s, int length) {
    std::unordered_set<char> A;
    for (int i = length - 1; i < s.length(); ++i) {
      A.clear();
      for (int j = 0; j < length; ++j) A.insert(s[i - j]);
      if (A.size() == length) return i + 1;
    }
    return -1;
  }

 public:
  int Part1(aoc::Input in) { return index_marker(in.get_line(), 4); }
  int Part2(aoc::Input in) { return index_marker(in.get_line(), 14); };
};

TEST(Day06, Part1) { EXPECT_EQ(Day06().Part1(aoc::Input(2022, 6)), 1287); }
TEST(Day06, Part2) { EXPECT_EQ(Day06().Part2(aoc::Input(2022, 6)), 3716); }

}  // namespace aoc2022

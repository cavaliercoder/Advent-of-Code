#include "lib/aoc.h"

namespace aoc2022 {

class Day03 {
  int priority(const char c) {
    if (c >= 'a' && c <= 'z') return 1 + static_cast<int>(c - 'a');
    if (c >= 'A' && c <= 'Z') return 27 + static_cast<int>(c - 'A');
    return -1;
  }

 public:
  int Part1(aoc::Input in) {
    int sum = 0;
    for (auto& s : in) {
      aoc::Set<char> backpack;
      for (int i = 0; i < s.size() / 2; ++i) {
        backpack += s[i];
      }
      for (int i = s.length() / 2; i < s.size(); ++i) {
        if (backpack.contains(s[i])) {
          sum += priority(s[i]);
          break;
        }
      }
    }
    return sum;
  }

  int Part2(aoc::Input in) {
    int i = 0, sum = 0;
    aoc::Set<char> group;
    for (auto& s : in) {
      aoc::Set<char> backpack;
      for (const char c : s) backpack += c;
      if (i == 0) {
        group = backpack;
      } else {
        group &= backpack;
      }
      if (i == 2) {
        assert(group.size() == 1);
        sum += priority(*group.begin());
        i = 0;
      } else {
        ++i;
      }
    }
    return sum;
  }
};

TEST(Day03, Part1) { EXPECT_EQ(Day03().Part1(aoc::Input(2022, 3)), 7793); }
TEST(Day03, Part2) { EXPECT_EQ(Day03().Part2(aoc::Input(2022, 3)), 2499); }

}  // namespace aoc2022

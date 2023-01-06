#include "lib/aoc.h"

namespace aoc2022 {

class Day04 {
  using Range = aoc::Range<int>;

  void parse(aoc::Input& in, Range& a, Range& b) {
    in.get_uint(a.start).expect('-').get_uint(a.limit).expect(',');
    in.get_uint(b.start).expect('-').get_uint(b.limit).expect('\n');
    ++a.limit;
    ++b.limit;
  }

 public:
  int Part1(aoc::Input in) {
    int n = 0;
    Range a, b;
    while (in) {
      parse(in, a, b);
      if (a.contains(b) || b.contains(a)) ++n;
    }
    return n;
  }

  int Part2(aoc::Input in) {
    int n = 0;
    Range a, b;
    while (in) {
      parse(in, a, b);
      if (a & b) ++n;
    }
    return n;
  }
};

TEST(Day04, Part1) { EXPECT_EQ(Day04().Part1(aoc::Input(2022, 4)), 569); }
TEST(Day04, Part2) { EXPECT_EQ(Day04().Part2(aoc::Input(2022, 4)), 936); }

}  // namespace aoc2022

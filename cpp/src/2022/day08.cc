#include <algorithm>
#include <cassert>
#include <numeric>
#include <vector>

#include "lib/aoc.h"

namespace aoc2022 {

using Grid = aoc::Grid<char>;
using Point = aoc::Point<2, int>;

class Day08 {
  const Point kUDLR[4] = {{0, 1}, {0, -1}, {-1, 0}, {1, 0}};

  bool is_visible(const Grid& g, const Point p) {
    for (auto d : kUDLR) {
      bool visible = true;
      for (Point q = p + d; g.contains(q); q += d) {
        if (g[q] >= g[p]) {
          visible = false;
          break;
        }
      }
      if (visible) return true;
    }
    return false;
  }

  int scenic_score(const Grid& g, const Point p) {
    int score = 1;
    for (auto d : kUDLR) {
      int dir_score = 0;
      for (Point q = p + d; g.contains(q); q += d) {
        ++dir_score;
        if (g[q] >= g[p]) break;
      }
      score *= dir_score;
    }
    return score;
  }

 public:
  int Part1(aoc::Input in) {
    auto g = in.grid();
    int count = 0;
    for (auto it = g.begin(); it < g.end(); ++it)
      if (is_visible(g, it.point())) ++count;
    return count;
  }

  int Part2(aoc::Input in) {
    auto g = in.grid();
    int best = 0;
    for (auto it = g.begin(); it < g.end(); ++it) {
      int score = scenic_score(g, it.point());
      if (score > best) best = score;
    }
    return best;
  }
};

TEST(Day08, Part1) { EXPECT_EQ(Day08().Part1(aoc::Input(2022, 8)), 1688); }
TEST(Day08, Part2) { EXPECT_EQ(Day08().Part2(aoc::Input(2022, 8)), 410400); }

}  // namespace aoc2022
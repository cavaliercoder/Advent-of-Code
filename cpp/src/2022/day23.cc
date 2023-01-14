#include <array>

#include "lib/aoc.h"

namespace aoc2022 {

using namespace aoc;

class Day23 {
  using Point = Point<2, int>;

  // Order of moves relative to input.
  const Point moves[4] = {
      /* North */ {0, -1},
      /* South */ {0, 1},
      /* West */ {-1, 0},
      /* East */ {1, 0},
  };

  // Buffer space needed for each move.
  const Point buffer[12] = {
      /* North */ {-1, -1}, {0, -1}, {1, -1},
      /* South */ {-1, 1},  {0, 1},  {1, 1},
      /* West */ {-1, -1},  {-1, 0}, {-1, 1},
      /* East */ {1, -1},   {1, 0},  {1, 1},
  };

  auto parse(aoc::Input& in) {
    auto g = in.grid();
    std::vector<Point> a;
    for (auto it = g.begin(); it < g.end(); ++it)
      if (*it == '#') a.push_back(it.point());
    return a;
  }

  int score(const std::vector<Point>& elves) {
    Point minp = elves[0], maxp = elves[0];
    for (auto p : elves) {
      minp = minp.min(p);
      maxp = maxp.max(p);
    }
    auto size = (maxp - minp) + 1;
    return size.x() * size.y() - elves.size();
  }

  // Run a single round and return the resulting position of each elf.
  std::vector<Point> step(const std::vector<Point>& src, const int first_move) {
    auto dst = src;
    auto proposers = std::unordered_map<Point, int>();
    const auto occupied = Set<Point>(src.begin(), src.end());
    for (int elf = 0; elf < src.size(); ++elf) {
      // Check for adjacent neighbors
      bool ok = false;
      const Point p = src[elf];
      for (int y = p.y() - 1; y <= p.y() + 1; ++y) {
        for (int x = p.x() - 1; x <= p.x() + 1; ++x) {
          Point q = Point(x, y);
          if (p == q) continue;
          if (occupied.count(q)) {
            ok = true;
          }
        }
      }
      if (!ok) continue;

      // Try the four moves.
      for (int i = 0; i < 4; ++i) {
        auto move = (first_move + i) % 4;
        auto isclear = true;
        for (int j = 0; j < 3; ++j) {
          if (occupied.count(p + buffer[(move * 3) + j])) {
            isclear = false;
            break;
          }
        }
        if (!isclear) continue;
        dst[elf] = src[elf] + moves[move];
        ++proposers[dst[elf]];
        break;
      }
    }

    // Check the proposals.
    for (int elf = 0; elf < src.size(); ++elf) {
      if (src[elf] == dst[elf]) continue;
      if (proposers.at(dst[elf]) > 1) dst[elf] = src[elf];
    }
    return dst;
  }

 public:
  int Part1(aoc::Input in, const int rounds = 10) {
    auto elves = parse(in);
    for (int r = 0; r < 10; ++r) {
      elves = step(elves, r % 4);
    }
    return score(elves);
  }

  int Part2(aoc::Input in) {
    auto elves = parse(in);
    int r = 0;
    while (true) {
      auto next = step(elves, r++ % 4);
      if (next == elves) return r;
      elves = next;
    }
  };
};

TEST(Day23, Part1) { EXPECT_EQ(Day23().Part1(aoc::Input(2022, 23)), 3862); }
TEST(Day23, Part2) { EXPECT_EQ(Day23().Part2(aoc::Input(2022, 23)), 913); }

}  // namespace aoc2022

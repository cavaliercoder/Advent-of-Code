#include <algorithm>
#include <array>

#include "lib/aoc.h"

namespace aoc2022 {

using namespace aoc;

class Day23 {
  using Point = Point<2, int>;

  // Flags represent which neighbors an elf has.
  enum Neighbors : char {
    None = 0x00,
    N = 0x01,
    S = 0x02,
    E = 0x04,
    W = 0x08,
  };

  // Order of moves to take (y-axis inverted).
  const Point moves[4] = {
      /* North */ {0, -1},
      /* South */ {0, 1},
      /* West */ {-1, 0},
      /* East */ {1, 0},
  };

  // Clearance required for each move.
  const Neighbors buffer[4] = {N, S, W, E};

  // Position of every elf in the input.
  std::vector<Point> elves_;

  int score(const std::vector<Point>& elves) const {
    Point minp = elves[0], maxp = elves[0];
    for (auto p : elves) {
      minp = minp.min(p);
      maxp = maxp.max(p);
    }
    auto size = (maxp - minp) + 1;
    return size.x() * size.y() - elves.size();
  }

  // Run a single round and return the resulting position of each elf.
  std::vector<Point> step(const std::vector<Point>& elves,
                          const int round) const {
    // Determine which elves are allowed to move and stored the location of
    // their neighbors as flags.
    //
    // I originally used a hash table of all the positions which made the code
    // very simple to read. Unfortunately it requires ~50M hash computations
    // to check for neighbors of each elf, which results in a runtime > 1s for
    // part 2 - even after switching up hash implementations.
    //
    // Doing a O(n2) comparison of each elf might seem slower on paper, but it
    // is shadowed by the constant-time cost of the hash table lookups.
    //
    // We save time below by sorting the elves by their x-axis and only
    // searching the subset of near neighbors. If the arrangement is reasonably
    // square, actual runtime is more like O(n * sqrt(n)).
    //
    // Runtime dropped to <300ms for part 2.
    auto src = elves;
    auto flags = std::vector<char>(src.size(), 0);
    std::sort(src.begin(), src.end());
    for (int i = 0; i < src.size() - 1; ++i) {
      for (int j = i + 1; j < src.size(); ++j) {
        auto delta = src[j] - src[i];
        if (delta.x() > 1) break;
        if (delta.abs().y() > 1) continue;
        if (delta.x() == 1) {
          flags[i] |= E;
          flags[j] |= W;
        }
        if (delta.y() == 1) {
          flags[i] |= S;
          flags[j] |= N;
        } else if (delta.y() == -1) {
          flags[i] |= N;
          flags[j] |= S;
        }
      }
    }

    // Propose a destination for each elf.
    auto dst = src;
    auto proposers = std::unordered_map<Point, int>();
    for (int elf = 0; elf < src.size(); ++elf) {
      if (!flags[elf]) continue;
      for (int i = 0; i < 4; ++i) {
        auto move = (round + i) % 4;
        if (flags[elf] & buffer[move]) continue;
        dst[elf] = src[elf] + moves[move];
        ++proposers[dst[elf]];
        break;
      }
    }

    // Reset elves with colliding proposals.
    for (int elf = 0; elf < src.size(); ++elf) {
      if (src[elf] == dst[elf]) continue;
      if (proposers.at(dst[elf]) > 1) dst[elf] = src[elf];
    }
    return dst;
  }

 public:
  Day23(aoc::Input in) {
    auto g = in.grid();
    for (auto it = g.begin(); it < g.end(); ++it)
      if (*it == '#') elves_.push_back(it.point());
  }

  int Part1(const int rounds = 10) const {
    auto elves = elves_;
    for (int r = 0; r < 10; ++r) {
      elves = step(elves, r);
    }
    return score(elves);
  }

  int Part2() const {
    auto elves = elves_;
    int r = 0;
    while (true) {
      auto next = step(elves, r++);
      if (next == elves) return r - 1;
      elves = next;
    }
  };
};

TEST(Day23, Part1) { EXPECT_EQ(Day23(aoc::Input(2022, 23)).Part1(), 3862); }

TEST(Day23, Part2) { EXPECT_EQ(Day23(aoc::Input(2022, 23)).Part2(), 913); }

}  // namespace aoc2022

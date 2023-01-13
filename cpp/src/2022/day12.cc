#include <limits.h>

#include "lib/aoc.h"

namespace aoc2022 {

class Day12 {
  using Point = aoc::Point<2, int>;

 public:
  int Part1(aoc::Input in) {
    auto g = in.grid();
    auto udlr = Point().orth();

    Point S, E;
    for (auto it = g.begin(); it < g.end(); ++it) {
      if (*it == 'S') {
        S = it.point();
        g[S] = 'a';
      }
      if (*it == 'E') {
        E = it.point();
        g[E] = 'z';
      }
    }

    // Distance of each node from S.
    auto distance = aoc::Grid<int>(g.width(), g.height(), INT_MAX);
    distance[S] = 0;

    // Min-heap of nodes to visit.
    auto Q = aoc::Heap(S, [&distance](const Point a, const Point b) -> bool {
      return distance[a] > distance[b];
    });
    while (Q) {
      auto p = Q.pop();
      assert(distance[p] != INT_MAX);
      if (p == E) break;
      for (auto dir : udlr) {
        auto q = p + dir;
        if (!g.contains(q)) continue;   // Out of bounds.
        if (g[q] > g[p] + 1) continue;  // Too high.
        if (distance[q] > distance[p] + 1) {
          distance[q] = distance[p] + 1;
          Q.push(q);
        }
      }
    }
    return distance[E];
  }

  int Part2(aoc::Input in) {
    auto g = in.grid();
    auto udlr = Point().orth();

    Point E;
    for (auto it = g.begin(); it < g.end(); ++it) {
      if (*it == 'S') {
        g[it] = 'a';
      }
      if (*it == 'E') {
        E = it.point();
        g[E] = 'z';
      }
    }

    // Distance of each node from S.
    auto distance = aoc::Grid<int>(g.width(), g.height(), INT_MAX);
    distance[E] = 0;
    int best = INT_MAX;

    // Min-heap of nodes to visit.
    auto Q = aoc::Heap(E, [&distance](const Point& a, const Point& b) -> bool {
      return distance[a] > distance[b];
    });
    while (Q) {
      auto p = Q.pop();
      if (distance[p] > best) continue;
      assert(distance[p] != INT_MAX);
      for (auto dir : udlr) {
        auto q = p + dir;
        if (!g.contains(q)) continue;   // Out of bounds.
        if (g[q] < g[p] - 1) continue;  // Too high.
        if (distance[q] > distance[p] + 1) {
          distance[q] = distance[p] + 1;
          if (g[q] == 'a') {
            best = std::min(best, distance[q]);
          } else {
            Q.push(q);
          }
        }
      }
    }
    return best;
  }
};

TEST(Day12, Part1) { EXPECT_EQ(Day12().Part1(aoc::Input(2022, 12)), 456); }
TEST(Day12, Part2) { EXPECT_EQ(Day12().Part2(aoc::Input(2022, 12)), 454); }

}  // namespace aoc2022

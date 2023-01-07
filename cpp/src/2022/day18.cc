#include <queue>

#include "lib/aoc.h"

namespace aoc2022 {

class Day18 {
  using Point = aoc::Point<3, int>;
  using Space = aoc::Space<3, int, char>;

  enum : char {
    Air = '\0',
    Lava = '#',
    Water = '~',
  };

  Space parse(aoc::Input& in) {
    int x, y, z;
    auto m = Space();
    while (in) {
      in.get_uint(x).expect(',').get_uint(y).expect(',').get_uint(z).expect(
          '\n');
      auto p = Point(x, y, z);
      m[p] = Lava;
    }
    return m;
  }

  // Fill all empty positions in the space with water using a DFS from the
  // source.
  void flood(Space& m, Point source) {
    assert(!m.contains(source));
    auto min = source - 1, max = m.max() + 1;
    std::queue<Point> Q;
    Q.push(source);
    while (!Q.empty()) {
      auto p = Q.front();
      Q.pop();
      if (p.x() < min.x() || p.y() < min.y() || p.z() < min.z()) continue;
      if (p.x() > max.x() || p.y() > max.y() || p.z() > max.z()) continue;
      if (m[p] != Air) continue;
      m[p] = Water;
      for (auto next : p.udlrfb()) Q.push(next);
    }
  }

 public:
  // Solves part 1 by checking the neighboring positions of each droplet.
  int Part1(aoc::Input in) {
    auto m = parse(in);
    int count = 0;
    for (auto [p, _] : m)
      for (auto q : p.udlrfb())
        if (!m.contains(q)) ++count;
    return count;
  }

  // Solves part 2 by flooding the space with water using a DFS from outside,
  // then checking each droplet for contact with the water.
  int Part2(aoc::Input in) {
    auto m = parse(in);
    int count = 0;
    flood(m, Point());
    for (auto [p, v] : m) {
      if (v != Lava) continue;
      for (auto q : p.udlrfb())
        if (m.at(q) == Water) ++count;
    }
    return count;
  };
};

TEST(Day18, Part1) { EXPECT_EQ(Day18().Part1(aoc::Input(2022, 18)), 4636); }
TEST(Day18, Part2) { EXPECT_EQ(Day18().Part2(aoc::Input(2022, 18)), 2572); }

}  // namespace aoc2022

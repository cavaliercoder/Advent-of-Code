#include <algorithm>

#include "lib/aoc.h"

namespace aoc2022 {

class Day14 {
  using Point = aoc::Point<2, int>;
  using Path = std::vector<Point>;
  using Space = aoc::Space<2, int, char>;

  enum Cell : unsigned char {
    Air = ' ',
    Rock = '#',
    Sand = 'o',
  };

  Space parse(aoc::Input& in) {
    std::vector<Path> paths;
    int x, y;
    auto path = Path();
    while (in) {
      in.get_uint(x).expect(',').get_uint(y);
      path.push_back(Point(x, y));
      if (in.branch('\n')) {
        paths.push_back(path);
        path = Path();
        continue;
      }
      in.expect(" -> ");
    }
    paths.push_back(path);
    auto m = Space(Air);
    for (const auto& path : paths) {
      for (int i = 1; i < path.size(); ++i) {
        for (Point p = path[i - 1], end = path[i];; p = p.nudge(end)) {
          m.insert(p, Rock);
          if (p == end) break;
        }
      }
    }
    return m;
  }

  Point drop(const Space& m, const Point src, const int yfloor) {
    Point p = src;
    while (true) {
      if (p.y() == yfloor - 1) return p;
      bool moved = false;
      // NB: inverted y-axis
      Point moves[] = {p.up(), p.up().left(), p.up().right()};
      for (auto next : moves) {
        if (m[next] == Air) {
          p = next;
          moved = true;
          break;
        }
      }
      if (!moved) break;
    }
    return p;
  }

 public:
  int Part1(aoc::Input in) {
    auto m = parse(in);
    int count = 0;
    Point src = {500, 0};
    int yfloor = m.max().y() + 2;
    while (true) {
      auto p = drop(m, src, yfloor);
      m[p] = Sand;
      if (p.y() == yfloor - 1) break;
      ++count;
    }
    return count;
  }

  int Part2(aoc::Input in) {
    auto m = parse(in);
    int count = 0;
    Point src = {500, 0};
    int yfloor = m.max().y() + 2;
    while (true) {
      auto p = drop(m, src, yfloor);
      m[p] = Sand;
      ++count;
      if (p == src) break;
    }
    return count;
  };
};

TEST(Day14, Part1) { EXPECT_EQ(Day14().Part1(aoc::Input(2022, 14)), 674); }
TEST(Day14, Part2) { EXPECT_EQ(Day14().Part2(aoc::Input(2022, 14)), 24958); }

}  // namespace aoc2022

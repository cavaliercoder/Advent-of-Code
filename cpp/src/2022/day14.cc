#include <algorithm>

#include "lib/aoc.h"

namespace aoc2022 {

class Day14 {
  using Point = aoc::Point<2, int>;
  using Path = std::vector<Point>;
  using Grid = aoc::Grid<char>;

  Grid grid_;
  int floor_ = 0;

  enum Cell : unsigned char {
    Air = ' ',
    Rock = '#',
    Sand = 'o',
  };

  Point drop(const Grid& m, const Point src = {500, 0}) const {
    Point p = src;
    while (true) {
      if (p.y() == floor_ - 1) return p;
      bool moved = false;
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
  Day14(aoc::Input in) {
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

    grid_ = Grid(1024, 1024, Air);
    for (const auto& path : paths) {
      for (int i = 1; i < path.size(); ++i) {
        for (Point p = path[i - 1], end = path[i];; p = p.nudge(end)) {
          if (!grid_.contains(p)) throw "input too large";
          grid_[p] = Rock;
          floor_ = std::max(floor_, p.y() + 2);
          if (p == end) break;
        }
      }
    }
  }

  int Part1() const {
    auto m = grid_;
    int count = 0;
    while (true) {
      auto p = drop(m);
      m[p] = Sand;
      if (p.y() == floor_ - 1) break;
      ++count;
    }
    return count;
  }

  int Part2() const {
    auto m = grid_;
    int count = 0;
    Point src = {500, 0};
    while (true) {
      auto p = drop(m, src);
      m[p] = Sand;
      ++count;
      if (p == src) break;
    }
    return count;
  };
};

TEST(Day14, Part1) { EXPECT_EQ(Day14(aoc::Input(2022, 14)).Part1(), 674); }
TEST(Day14, Part2) { EXPECT_EQ(Day14(aoc::Input(2022, 14)).Part2(), 24958); }

}  // namespace aoc2022

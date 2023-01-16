#include "lib/aoc.h"

namespace aoc2022 {

class Day09 {
  using Point = aoc::Point<2, int>;

  struct Vector {
    Point dir;
    int mag;

    static Vector parse_one(aoc::Input& in) {
      char dir;
      int mag;
      in.get(dir).expect(' ').get_uint(mag).expect('\n');
      switch (dir) {
        case 'U':
          return Vector{Point(0, 1), mag};
        case 'D':
          return Vector{Point(0, -1), mag};
        case 'L':
          return Vector{Point(-1, 0), mag};
        case 'R':
          return Vector{Point(1, 0), mag};
      };
      assert(0);
    }

    static std::vector<Vector> parse(aoc::Input& in) {
      std::vector<Vector> vectors;
      while (in) vectors.push_back(parse_one(in));
      return vectors;
    }
  };

  Point move_tail(const Point head, const Point tail) {
    Point delta = (tail - head).abs();
    if (delta.x() <= 1 && delta.y() <= 1) return Point();
    assert(delta <= Point(2, 2));

    // Simple horizontal/vertical move
    if (delta.x() == 0 || delta.y() == 0) {
      if (tail.x() != head.x()) return Point(tail.x() < head.x() ? 1 : -1, 0);
      return Point(0, tail.y() < head.y() ? 1 : -1);
    }

    // Diagonal move
    return Point(tail.x() < head.x() ? 1 : -1, tail.y() < head.y() ? 1 : -1);
  }

 public:
  int Part1(aoc::Input in) {
    auto vectors = Vector::parse(in);
    Point head, tail;
    aoc::Set<Point> seen;
    seen += tail;
    for (auto v : vectors) {
      for (int i = 0; i < v.mag; ++i) {
        head += v.dir;
        tail += move_tail(head, tail);
        seen += tail;
      }
    }
    return seen.size();
  }

  int Part2(aoc::Input in) {
    auto vectors = Vector::parse(in);
    std::vector<Point> rope = {10, Point()};
    aoc::Set<Point> seen;
    seen += Point();
    for (auto v : vectors) {
      for (int i = 0; i < v.mag; ++i) {
        rope[0] += v.dir;
        for (int i = 1; i < rope.size(); ++i)
          rope[i] += move_tail(rope[i - 1], rope[i]);
        seen += rope.back();
      }
    }
    return seen.size();
  }
};

TEST(Day09, Part1) { EXPECT_EQ(Day09().Part1(aoc::Input(2022, 9)), 6044); }
TEST(Day09, Part2) { EXPECT_EQ(Day09().Part2(aoc::Input(2022, 9)), 2384); }

}  // namespace aoc2022
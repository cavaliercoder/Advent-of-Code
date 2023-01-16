#include "point.h"

#include <algorithm>

#include "testing.h"

namespace aoc {

TEST(AoC, Point) {
  using Point = aoc::Point<2, int>;

  // Default constructor.
  Point p;
  EXPECT_EQ(p, Point());
  EXPECT_EQ(p.size(), 2);
  EXPECT_EQ(p.x(), 0);
  EXPECT_EQ(p.y(), 0);

  // Assignement
  p = Point(3, 4);
  EXPECT_EQ(p, Point(3, 4));

  // Components
  EXPECT_EQ(Point(5, 6).x(), 5);
  EXPECT_EQ(Point(5, 6).y(), 6);

  // Accessors.
  p = Point(5, 6);
  EXPECT_EQ(p[0], 5);
  EXPECT_EQ(p[1], 6);
  p[0] = -1;
  p[1] = -2;
  EXPECT_EQ(p, Point(-1, -2));

  // Comparison.
  EXPECT_EQ(Point(1, 2), Point(1, 2));
  EXPECT_EQ(Point(68, 70), Point(68, 70));
  EXPECT_NE(Point(1, 2), Point(3, 4));
  EXPECT_NE(Point(50, 0), Point(50, -22));

  // Sorting
  std::vector<Point> a = {{3, 4}, {-10, 500}, {1, -5}, {1, -6}};
  std::vector<Point> b = {{-10, 500}, {1, -6}, {1, -5}, {3, 4}};
  std::sort(a.begin(), a.end());
  for (int i = 0; i < a.size(); ++i) {
    EXPECT_EQ(a[i], b[i]);
    EXPECT_LE(a[i], b[i]);
    EXPECT_GE(a[i], b[i]);
  }
  EXPECT_LT(Point(3, 4), Point(3, 5));
  EXPECT_LT(Point(2, 4), Point(3, 5));
  EXPECT_LE(Point(3, 4), Point(3, 5));
  EXPECT_LE(Point(2, 4), Point(3, 5));

  // Addition.
  p = Point(1, 1);
  EXPECT_EQ(p + Point(2, 2), Point(3, 3));
  EXPECT_EQ(p, Point(1, 1));
  EXPECT_EQ(p += Point(3, 4), Point(4, 5));
  EXPECT_EQ(p, Point(4, 5));
  EXPECT_EQ(p++, Point(4, 5));
  EXPECT_EQ(++p, Point(6, 7));

  // Rotation
  p = {3, 6};
  EXPECT_EQ(p = p.cw(), Point(6, -3));
  EXPECT_EQ(p = p.cw(), Point(-3, -6));
  EXPECT_EQ(p = p.cw(), Point(-6, 3));
  EXPECT_EQ(p = p.cw(), Point(3, 6));

  p = {2, 4};
  EXPECT_EQ(p = p.ccw(), Point(-4, 2));
  EXPECT_EQ(p = p.ccw(), Point(-2, -4));
  EXPECT_EQ(p = p.ccw(), Point(4, -2));
  EXPECT_EQ(p = p.ccw(), Point(2, 4));

  // Orthogonal
  p = {1, 1};
  auto [l, r, d, u] = p.orth();
  EXPECT_EQ(u, Point(1, 2));
  EXPECT_EQ(d, Point(1, 0));
  EXPECT_EQ(l, Point(0, 1));
  EXPECT_EQ(r, Point(2, 1));
  EXPECT_EQ(p, p.up().right().down().left());

  // Negation.
  EXPECT_EQ(-Point(-1, 2), Point(1, -2));

  // Absolute value.
  EXPECT_EQ(Point(0, 0).abs(), Point());
  EXPECT_EQ(Point(1, 2).abs(), Point(1, 2));
  EXPECT_EQ(Point(-1, 2).abs(), Point(1, 2));
  EXPECT_EQ(Point(1, -2).abs(), Point(1, 2));
  EXPECT_EQ(Point(-1, -2).abs(), Point(1, 2));

  // Min/max.
  EXPECT_EQ(Point(1, 2).min(Point(-1, 3)), Point(-1, 2));
  EXPECT_EQ(Point(1, 2).max(Point(-1, 3)), Point(1, 3));

  // Orientation.
  EXPECT_EQ(Point().orientation(), Point());
  EXPECT_EQ(Point(10, 10).orientation(), Point(1, 1));
  EXPECT_EQ(Point(123, -456).orientation(), Point(1, -1));
  EXPECT_EQ(Point(-123, 456).orientation(), Point(-1, 1));
  EXPECT_EQ(Point(-123, 456).orientation(),
            Point(-123, 456).orientation().orientation());

  // Nudging.
  // Dimensions: {inc, inc*2, dec, dec*2, static}.
  using Point5D = aoc::Point<5, int>;
  Point5D p_start = {1, 2, 3, 4, 5};
  Point5D p_next = {2, 3, 2, 3, 5};
  Point5D p_end = {2, 4, 2, 2, 5};
  EXPECT_EQ(p_start = p_start.nudge(p_start), p_start);
  EXPECT_EQ(p_start = p_start.nudge(p_end), p_next);
  EXPECT_EQ(p_start = p_start.nudge(p_end), p_end);
  EXPECT_EQ(p_start = p_start.nudge(p_end), p_end);

  // constexpr
  static constexpr auto c =
      (Point(-3, 40).orientation() * 4).abs().max({0, 6}).nudge({}).down().y();
  EXPECT_EQ(c, 4);
}

}  // namespace aoc
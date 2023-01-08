#include "point.h"

#include "testing.h"

namespace aoc {

TEST(AoC, Point) {
  using Point = aoc::Point<2, int>;

  auto p = Point(1, 1);
  EXPECT_EQ(p + Point(2, 2), Point(3, 3));
  EXPECT_EQ(p, Point(1, 1));
  EXPECT_EQ(p += Point(3, 4), Point(4, 5));
  EXPECT_EQ(p, Point(4, 5));

  auto [u, d, l, r] = Point(1, 1).udlr();
  EXPECT_EQ(u, Point(1, 2));
  EXPECT_EQ(d, Point(1, 0));
  EXPECT_EQ(l, Point(0, 1));
  EXPECT_EQ(r, Point(2, 1));

  EXPECT_EQ(-Point(-1, 2), Point(1, -2));

  EXPECT_EQ(Point(0, 0).abs(), Point());
  EXPECT_EQ(Point(1, 2).abs(), Point(1, 2));
  EXPECT_EQ(Point(-1, 2).abs(), Point(1, 2));
  EXPECT_EQ(Point(1, -2).abs(), Point(1, 2));
  EXPECT_EQ(Point(-1, -2).abs(), Point(1, 2));

  EXPECT_EQ(Point(1, 2).min(Point(-1, 3)), Point(-1, 2));
  EXPECT_EQ(Point(1, 2).max(Point(-1, 3)), Point(1, 3));

  // Dimensions: {inc, inc*2, dec, dec*2, static}.
  using Point5D = aoc::Point<5, int>;
  Point5D p_start = {1, 2, 3, 4, 5};
  Point5D p_next = {2, 3, 2, 3, 5};
  Point5D p_end = {2, 4, 2, 2, 5};
  EXPECT_EQ(p_start = p_start.nudge(p_start), p_start);
  EXPECT_EQ(p_start = p_start.nudge(p_end), p_next);
  EXPECT_EQ(p_start = p_start.nudge(p_end), p_end);
  EXPECT_EQ(p_start = p_start.nudge(p_end), p_end);
}

}  // namespace aoc
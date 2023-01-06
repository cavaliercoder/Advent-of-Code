
#include "aoc.h"

#include <algorithm>
#include <iostream>
#include <vector>

namespace aoc {

std::vector<const BaseTestFactory*> test_factories;

int register_test(const BaseTestFactory* factory) {
  test_factories.push_back(factory);
  return test_factories.size();
}

std::vector<BaseTest*> get_tests() {
  std::vector<BaseTest*> tests;
  for (auto& factory : test_factories) {
    tests.push_back(factory->make());
  }
  std::sort(
      tests.begin(), tests.end(),
      [](const BaseTest* a, const BaseTest* b) -> bool { return *a < *b; });
  return tests;
}

int run_tests() {
  int fail_count = 0;
  auto tests = get_tests();
  StopWatch sw, sw_all;
  sw_all.start();
  for (int i = 0; i < tests.size(); ++i) {
    auto test = tests[i];
    std::cout << "[" << (i + 1) << "/" << tests.size() << "]";
    std::cout << " " << test->suite() << "::" << test->name();
    sw.start();
    test->run();
    sw.stop();
    std::cout << " " << sw;
    if (*test) {
      std::cout << " [\u001b[32mPASS\u001b[0m]\n";
    } else {
      ++fail_count;
      std::cout << " [\u001b[31mFAIL\u001b[0m]\n";
      for (auto& err : test->errv()) {
        std::cout << "  " << err.file() << ":" << err.line()
                  << " Expression evaluates to false: " << err.expr() << " "
                  << err.op() << " " << err.expect() << "\n"
                  << "    " << err.expr() << " → " << err.actual() << "\n";
      }
      std::cout << "\n";
    }
  }
  sw_all.stop();

  std::cout << "\nPassed " << (tests.size() - fail_count) << "/" << tests.size()
            << " tests in " << sw_all << ".\n";
  for (auto& test : tests) delete test;
  return fail_count;
}

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
}

TEST(AoC, Grid) {
  using Grid = aoc::Grid<char>;
  using Point = aoc::Point<2, int>;

  auto g = Grid(4, 3, ' ');
  EXPECT_EQ(g.width(), 4);
  EXPECT_EQ(g.height(), 3);
  EXPECT_EQ(g.size(), 12);
  EXPECT_EQ(g.count(' '), 12);

  // Test boundaries
  EXPECT_FALSE(g.contains(-1));
  EXPECT_FALSE(g.contains(12));
  EXPECT_FALSE(g.contains(Point(-1, -1)));
  EXPECT_FALSE(g.contains(Point(4, 3)));
  EXPECT_FALSE(g.contains(Point(4, 1)));
  EXPECT_FALSE(g.contains(Point(1, 3)));

  // Rewrite the whole grid with unique values
  int i = 0;
  for (auto it = g.begin(); it != g.end(); ++it) {
    EXPECT_EQ(*it, ' ');
    char c = 'A' + i++;
    g[it] = c;
    EXPECT_EQ(*it, c);
  }

  // Test formatting
  EXPECT_EQ(g.str(), "ABCD\nEFGH\nIJKL\n");

  // Test direct access
  std::array<Point, 12> points = {
      Point(0, 0), Point(1, 0), Point(2, 0), Point(3, 0),
      Point(0, 1), Point(1, 1), Point(2, 1), Point(3, 1),
      Point(0, 2), Point(1, 2), Point(2, 2), Point(3, 2),
  };
  for (int i = 0; i < points.size(); ++i) {
    Point p = points[i];
    char c = 'A' + i;
    EXPECT_TRUE(g.contains(p));
    EXPECT_EQ(g[p], g[i]);
    EXPECT_EQ(g[p], c);
    EXPECT_EQ(g.count(c), 1);
  }

  // Test for-each loop
  i = 0;
  for (auto c : g) {
    EXPECT_EQ(c, 'A' + i++);
  }
  EXPECT_EQ(i, 12);

  // Test forward iteration
  i = 0;
  auto it = g.begin();
  for (int y = 0; y < g.height(); ++y) {
    for (int x = 0; x < g.width(); ++x) {
      EXPECT_NE(it, g.end());
      EXPECT_LT(it, g.end());
      EXPECT_LE(it, g.end());
      EXPECT_GT(g.end(), it);
      EXPECT_GE(g.end(), it);
      EXPECT_EQ(*it, 'A' + i);
      EXPECT_EQ(it.index(), i);
      EXPECT_EQ(it.point(), Point(x, y));
      ++it;
      ++i;
    }
  }
  EXPECT_EQ(it, g.end());
  EXPECT_LE(it, g.end());
  EXPECT_EQ(g.end(), it);

  // Test seek iteration
  it = g.begin();
  EXPECT_EQ(it.point(), Point());
  EXPECT_EQ(it.index(), 0);
  EXPECT_EQ(*it, 'A');
  EXPECT_EQ(g[it], 'A');
  EXPECT_EQ(g[it.index()], 'A');
  EXPECT_EQ(g[it.point()], 'A');

  // Post-fix increment
  EXPECT_EQ(*(it++), 'A');
  EXPECT_EQ(it.point(), Point(1, 0));
  EXPECT_EQ(it.index(), 1);
  EXPECT_EQ(*it, 'B');
  EXPECT_EQ(g[it], 'B');
  EXPECT_EQ(g[it.index()], 'B');
  EXPECT_EQ(g[it.point()], 'B');

  // Prefix increment
  EXPECT_EQ(*(++it), 'C');
  EXPECT_EQ(it.point(), Point(2, 0));
  EXPECT_EQ(it.index(), 2);
  EXPECT_EQ(*it, 'C');
  EXPECT_EQ(g[it], 'C');
  EXPECT_EQ(g[it.index()], 'C');
  EXPECT_EQ(g[it.point()], 'C');

  // Increment N
  EXPECT_EQ(*(it += 5), 'H');
  EXPECT_EQ(it.point(), Point(3, 1));
  EXPECT_EQ(it.index(), 7);
  EXPECT_EQ(*it, 'H');
  EXPECT_EQ(g[it], 'H');
  EXPECT_EQ(g[it.index()], 'H');
  EXPECT_EQ(g[it.point()], 'H');
}

}  // namespace aoc
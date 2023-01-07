#include "aoc.h"

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

TEST(AoC, Heap) {
  auto data = {1, 3, 0, 4, 2};

  auto maxheap = Heap(data);
  for (int i = 4; i >= 0; --i) {
    EXPECT_TRUE(maxheap);
    EXPECT_EQ(maxheap.pop(), i);
  }
  EXPECT_FALSE(maxheap);
  maxheap.push(data.begin(), data.end());
  EXPECT_EQ(maxheap.pop(), 4);

  auto minheap = Heap(data, std::greater());
  for (int i = 0; i < 5; ++i) {
    EXPECT_TRUE(minheap);
    EXPECT_EQ(minheap.pop(), i);
  }
  EXPECT_FALSE(minheap);
  minheap.push(data.begin(), data.end());
  EXPECT_EQ(minheap.pop(), 0);
}

TEST(AoC, Hex) {
  EXPECT_EQ(hex(uint8_t(0x12)), "12");
  EXPECT_EQ(hex(uint16_t(0x12)), "0012");
  EXPECT_EQ(hex(uint16_t(0x1200)), "1200");
  EXPECT_EQ(hex(uint32_t(0x12)), "00000012");
  EXPECT_EQ(hex(uint16_t(0x3456)), "3456");
  EXPECT_EQ(hex(uint32_t(0x789ABCDE)), "789abcde");
  EXPECT_EQ(hex(uint64_t(0x0123456789ABCDEF)), "0123456789abcdef");
}

TEST(AoC, CRC32) {
  using CRC32 = CRC32<>;

  // Cases: https://rosettacode.org/wiki/CRC-32#C++
  char buf[32];
  for (int i = 0; i < 32; ++i) buf[i] = 0;
  EXPECT_EQ(CRC32(&buf[0], 32), 0x190A55AD);
  for (int i = 0; i < 32; ++i) buf[i] = 0xFF;
  EXPECT_EQ(CRC32(&buf[0], 32), 0xFF6CAB0B);
  for (int i = 0; i < 32; ++i) buf[i] = i;
  EXPECT_EQ(CRC32(&buf[0], 32), 0x91267E8A);

  auto quote = "The quick brown fox jumps over the lazy dog";
  EXPECT_EQ(CRC32(quote, 43).str(), "414fa339");

  CRC32 crc32;
  for (auto c : std::string(quote)) {
    crc32.update(&c, sizeof(c));
  }
  EXPECT_EQ(*crc32, 0x414FA339);
}

TEST(AoC, Digest) {
  const char* data = "The quick brown fox jumps over the lazy dog";
  const size_t len = 43;

  auto sha1 = SHA1(data, len);
  EXPECT_EQ(sha1().str(), "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12");

  auto sha256 = SHA256();
  EXPECT_EQ(sha256.str(),
            "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855");
  sha256.update(data, len);
  EXPECT_EQ(sha256.str(),
            "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592");
  sha256.reset();
  EXPECT_EQ(sha256.str(),
            "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855");
}

}  // namespace aoc

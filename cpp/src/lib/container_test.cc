#include "container.h"

#include "testing.h"

namespace aoc {

TEST(Aoc, Set) {
  using Set = aoc::Set<int>;

  // Comparison
  EXPECT_EQ(Set(), Set());
  EXPECT_EQ(Set({1, 2, 3}), Set({1, 2, 3}));
  EXPECT_NE(Set(), Set({1, 2, 3}));

  EXPECT_LT(Set({1}), Set({1, 2, 3}));
  EXPECT_FALSE(Set({1, 2, 3}) < Set({1, 2, 3}));
  EXPECT_FALSE(Set({1, 2, 3}) < Set({1}));

  EXPECT_LE(Set(), Set());
  EXPECT_LE(Set({1, 2, 3}), Set({1, 2, 3}));
  EXPECT_LE(Set({1, 2}), Set({1, 2, 3}));
  EXPECT_FALSE(Set({1, 2, 3, 4}) <= Set({1, 2, 3}));

  EXPECT_GT(Set({1, 2, 3}), Set({1, 2}));
  EXPECT_FALSE(Set({1, 2, 3}) > Set({1, 2, 3}));
  EXPECT_FALSE(Set({1}) > Set({1, 2}));

  EXPECT_GE(Set(), Set());
  EXPECT_GE(Set({1, 2, 3}), Set({1, 2, 3}));
  EXPECT_GE(Set({1, 2, 3}), Set({1, 2}));
  EXPECT_FALSE(Set({1, 2, 3}) >= Set({1, 2, 3, 4}));

  // Membership
  EXPECT_TRUE(Set({1, 2, 3}).contains(2));
  EXPECT_FALSE(Set({1, 2, 3}).contains(4));
  EXPECT_TRUE(Set({1, 2, 3}).contains(Set({1, 2})));
  EXPECT_FALSE(Set({1, 2, 3}).contains(Set({1, 2, 4})));

  // Addition
  Set s = {1, 2, 3};
  EXPECT_EQ(s += 4, Set({1, 2, 3, 4}));
  EXPECT_EQ(s += Set({5, 6}), Set({1, 2, 3, 4, 5, 6}));
  EXPECT_EQ(Set({1, 2, 3}) + 4, Set({1, 2, 3, 4}));
  EXPECT_EQ(Set({1, 2, 3}) + Set({4, 5, 6}), Set({1, 2, 3, 4, 5, 6}));
  EXPECT_EQ(Set() + Set(), Set());

  // Subtraction
  s = {1, 2, 3, 4, 5, 6};
  EXPECT_EQ(s -= 4, Set({1, 2, 3, 5, 6}));
  EXPECT_EQ(s -= Set({5, 6}), Set({1, 2, 3}));
  EXPECT_EQ(Set({1, 2, 3, 4}) - 4, Set({1, 2, 3}));
  EXPECT_EQ(Set({1, 2, 3, 4, 5, 6}) - Set({4, 5, 6}), Set({1, 2, 3}));
  EXPECT_EQ(Set() - Set(), Set());

  // Union
  s = {1, 2, 3};
  EXPECT_EQ(s |= Set({3, 4, 5, 6}), Set({1, 2, 3, 4, 5, 6}));
  EXPECT_EQ(Set({1, 2, 3}) | Set({1, 4, 5, 6}), Set({1, 2, 3, 4, 5, 6}));
  EXPECT_EQ(Set() | Set(), Set());

  // Intersect
  s = {1, 2, 3};
  EXPECT_EQ(s &= Set({2, 3, 4}), Set({2, 3}));
  EXPECT_EQ(Set({1, 2, 3, 4}) & Set({3, 4, 5, 6}), Set({3, 4}));
  EXPECT_EQ(Set({1, 2}) & Set({3, 4}), Set());
  EXPECT_EQ(Set() & Set(), Set());

  // Difference
  s = {1, 2, 3};
  EXPECT_EQ(s ^= Set({2, 3, 4}), Set({1, 4}));
  EXPECT_EQ(Set({1, 2, 3, 4}) ^ Set({3, 4, 5, 6}), Set({1, 2, 5, 6}));
  EXPECT_EQ(Set({1, 2}) ^ Set({3, 4}), Set({1, 2, 3, 4}));
  EXPECT_EQ(Set() ^ Set(), Set());

  // Formatting
  std::stringstream ss;
  ss << Set({3, 4, 2, 1});
  EXPECT_EQ(ss.str(), "{1, 2, 3, 4}")
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

TEST(Aoc, UniqueStack) {
  auto stack = aoc::UniqueStack<int>();
  for (int n = 8; n > 0; --n) {
    for (int i = 0; i < 3; ++i) stack.push(n);
    for (int i = n; i <= 8; ++i) stack.push(i);
    EXPECT_EQ(stack.peek(), n);
    EXPECT_EQ(stack.size(), 8 - n + 1);
  }
  EXPECT_EQ(stack.max_size(), 8);
  EXPECT_EQ(stack.stacked(), 8);
  for (int n = 1; n <= 8; ++n) {
    EXPECT_EQ(stack.pop(), n);
  }
  EXPECT_FALSE(stack);
  stack.clear();
  stack.push(1);
  EXPECT_TRUE(stack);
  EXPECT_EQ(stack.pop(), 1);
}

}  // namespace aoc

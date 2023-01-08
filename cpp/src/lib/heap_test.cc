#include "heap.h"

#include "testing.h"

namespace aoc {

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

}  // namespace aoc

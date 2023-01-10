#include "range.h"

#include "testing.h"

namespace aoc {

TEST(AoC, Range) {
  auto r = Range(10, 20);
  int i = 10;
  for (auto it = r.begin(); it != r.end(); ++it) {
    EXPECT_TRUE(i < 20);
    EXPECT_EQ(*it, i++);
  }
}

}  // namespace aoc

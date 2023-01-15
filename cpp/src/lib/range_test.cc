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

  // Regression case for flattening.
  std::vector<Range<int64_t>> a = {
      {-1092964, 1125469}, {-71069, 2729596},  {121154, 902799},
      {1332603, 2601208},  {2729595, 3034484}, {2729595, 3293868},
      {3091625, 3097306},  {3097305, 4163648}, {3809745, 3996538},
      {86103, 771242},
  };
  Range<int64_t>::flatten(a);
  EXPECT_EQ(a.size(), 1);
  EXPECT_EQ(a[0], Range(-1092964, 4163648));
}

}  // namespace aoc

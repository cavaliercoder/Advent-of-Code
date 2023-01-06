#include <algorithm>
#include <cassert>
#include <numeric>
#include <vector>

#include "lib/aoc.h"

namespace aoc2022 {

class Day01 {
  int sumTopNCalorieElves(aoc::Input& in, const int n) {
    std::vector<int> a;

    auto append_elf_cals = [&a, n](int cals) {
      a.push_back(cals);
      sort(a.begin(), a.end(), std::greater<int>());
      if (a.size() > n) {
        a.resize(n);
      }
    };

    int cals = 0, elf_cals = 0;
    for (const auto& s : in) {
      if (s.empty()) {
        append_elf_cals(elf_cals);
        elf_cals = 0;
        continue;
      }
      cals = stoi(s);
      elf_cals += cals;
    }
    append_elf_cals(elf_cals);
    assert(a.size() == n);
    return std::accumulate(a.begin(), a.end(), 0);
  }

 public:
  int Part1(aoc::Input in) { return sumTopNCalorieElves(in, 1); }
  int Part2(aoc::Input in) { return sumTopNCalorieElves(in, 3); }
};

TEST(Day01, Part1) { EXPECT_EQ(Day01().Part1(aoc::Input(2022, 1)), 69528); }
TEST(Day01, Part2) { EXPECT_EQ(Day01().Part2(aoc::Input(2022, 1)), 206152); }

}  // namespace aoc2022
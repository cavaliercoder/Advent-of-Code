#include <iostream>
#include <vector>

#include "lib/aoc.h"

namespace aoc2022 {

class Day20 {
  std::vector<int64_t> parse(aoc::Input& in) {
    std::vector<int64_t> a;
    int64_t n;
    while (in) {
      in.get_int(n).expect('\n');
      a.push_back(n);
    }
    return a;
  }

  int64_t mix(std::vector<int64_t>& a, const int rounds = 1,
              const int64_t key = 1) {
    int64_t len = a.size() - 1;
    std::unordered_map<int64_t, int64_t> m;  // old -> new
    for (int64_t i = 0; i < a.size(); ++i) {
      a[i] *= key;
      m[i] = i;
    }
    for (int r = 0; r < rounds; ++r) {
      for (int64_t id = 0; id < a.size(); ++id) {
        int64_t delta = a[id];
        int64_t src = m[id];
        int64_t dst = (src + delta) % len;
        if (dst < 0) dst += len;
        assert(dst >= 0 && dst < len);
        if (src < dst) {
          for (auto& [_, i] : m)
            if (i > src && i <= dst) --i;
        } else if (dst < src) {
          for (auto& [_, i] : m)
            if (i >= dst && i < src) ++i;
        }
        m[id] = dst;
      }
    }

    // Apply new positions
    auto b = std::vector<int64_t>(a.size(), INT_MAX);
    for (auto [id, i] : m) b[i] = a[id];
    a = b;
    return score(a);
  }

  int64_t score(std::vector<int64_t>& a) {
    for (int i = 0; i < a.size(); ++i) {
      if (a[i] != 0) continue;
      int64_t sum = 0;
      sum += a[(1000 + i) % a.size()];
      sum += a[(2000 + i) % a.size()];
      sum += a[(3000 + i) % a.size()];
      return sum;
    }
    return 0;
  }

 public:
  int64_t Part1(aoc::Input in) {
    auto a = parse(in);
    return mix(a);
  }

  int64_t Part2(aoc::Input in) {
    auto a = parse(in);
    return mix(a, 10, 811589153);
  };
};

TEST(Day20, Part1) { EXPECT_EQ(Day20().Part1(aoc::Input(2022, 20)), 3466); }

TEST(Day20, Part2) {
  EXPECT_EQ(Day20().Part2(aoc::Input(2022, 20)), 9995532008348);
}

}  // namespace aoc2022

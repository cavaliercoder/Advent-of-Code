#include <cstdint>

#include "lib/aoc.h"

namespace aoc2022 {

class Day25 {
 public:
  template <typename T = uint64_t>
  class Snafu {
    T value_ = 0;

   public:
    Snafu() = default;
    Snafu(const T n) : value_(n) {}

    template <class Iter>
    Snafu(Iter begin, Iter end) {
      for (auto it = begin; begin < end; ++it) {
        if (!(*it)) break;
        value_ *= 5;
        switch (*it) {
          case '0':
          case '1':
          case '2':
            value_ += *it - '0';
            break;
          case '-':
            value_ -= 1;
            break;
          case '=':
            value_ -= 2;
            break;
          default:
            throw "bad snafu";
        }
      }
    }

    std::string str() const {
      static const auto table = "012=-";
      std::vector<char> s;
      T n = value_;
      while (n) {
        T m = n % 5;
        s.push_back(table[m]);
        if (m >= 3) n += 5;
        n /= 5;
      }
      std::reverse(s.begin(), s.end());
      return std::string(s.begin(), s.end());
    }

    operator T() const { return value_; }

    Snafu& operator+=(const Snafu& s) {
      value_ += s.value_;
      return *this;
    }
  };

  std::string Part1(aoc::Input in) {
    Snafu sum;
    for (auto s : in) {
      sum += Snafu(s.begin(), s.end());
    }
    return sum.str();
  }
};

TEST(Day25, Part1) {
  EXPECT_EQ(Day25().Part1(aoc::Input(2022, 25)), "2---1010-0=1220-=010");
}

}  // namespace aoc2022

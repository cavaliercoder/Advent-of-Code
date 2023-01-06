#include "lib/aoc.h"

namespace aoc2022 {

class Day02 {
  enum Throw : int { Rock = 1, Paper = 2, Scissors = 3 };
  enum Outcome : char { Loss = 'X', Draw = 'Y', Win = 'Z' };

  Throw beat(const Throw t) {
    switch (t) {
      case Rock:
        return Paper;
      case Paper:
        return Scissors;
      case Scissors:
        return Rock;
    }
  }

  Throw lose_to(const Throw t) {
    switch (t) {
      case Rock:
        return Scissors;
      case Paper:
        return Rock;
      case Scissors:
        return Paper;
    }
  }

  inline int score(const Throw t) { return static_cast<int>(t); }

  int score(const Throw theirs, const Throw mine) {
    int n = score(mine);
    if (mine == theirs) return n + 3;        // Draw
    if (mine == beat(theirs)) return n + 6;  // Win
    return n;                                // Loss
  }

 public:
  int Part1(aoc::Input in) {
    int n = 0;
    char ca = 0, cb = 0;
    while (in) {
      in.get(ca).expect(' ').get(cb).expect('\n');
      const auto theirs = Throw(ca - 'A' + 1);
      const auto mine = Throw(cb - 'X' + 1);
      n += score(theirs, mine);
    }
    return n;
  }

  int Part2(aoc::Input in) {
    int n = 0;
    char ca = 0, cb = 0;
    while (in) {
      in.get(ca).expect(' ').get(cb).expect('\n');
      const auto theirs = Throw(ca - 'A' + 1);
      const auto outcome = Outcome(cb);
      Throw mine;
      switch (outcome) {
        case Loss:
          mine = lose_to(theirs);
          break;
        case Draw:
          mine = theirs;
          break;
        case Win:
          mine = beat(theirs);
          break;
      }
      n += score(theirs, mine);
    }
    return n;
  }
};

TEST(Day02, Part1) { EXPECT_EQ(Day02().Part1(aoc::Input(2022, 2)), 15337); }
TEST(Day02, Part2) { EXPECT_EQ(Day02().Part2(aoc::Input(2022, 2)), 11696); }

}  // namespace aoc2022
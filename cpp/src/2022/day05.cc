#include <algorithm>

#include "lib/aoc.h"

namespace aoc2022 {

class Day05 {
  struct Move {
    int count;
    int src;
    int dst;
  };

  void parse(aoc::Input& in, std::vector<std::vector<char>>& stacks,
             std::vector<Move>& moves) {
    // Parse stacks
    while (in) {
      if (in.branch('\n')) break;  // Empty line delimiter
      int stack = 0;
      while (!in.branch('\n')) {
        if (stack >= stacks.size()) stacks.push_back({});
        if (stack > 0) in.expect(' ');
        if (in.branch('[')) {
          stacks[stack].push_back(in.get());
          in.expect(']');
        } else {
          in.ignore(3);  // Empty slot
        }
        ++stack;
      }
    }
    for (auto& stack : stacks) {
      std::reverse(stack.begin(), stack.end());
    }

    // Parse moves
    while (in) {
      if (!in.branch('m')) break;
      Move m;
      in.ignore(4).get_uint(m.count);
      in.ignore(6).get_uint(m.src);
      in.ignore(4).get_uint(m.dst).expect('\n');
      --m.src;
      --m.dst;
      moves.push_back(m);
    }
  }

  std::string answer(std::vector<std::vector<char>> stacks) {
    auto s = std::string();
    for (auto& stack : stacks) s += stack.back();
    return s;
  }

 public:
  std::string Part1(aoc::Input in) {
    std::vector<std::vector<char>> stacks;
    std::vector<Move> moves;
    parse(in, stacks, moves);
    for (auto& m : moves) {
      for (int i = 0; i < m.count; ++i) {
        stacks[m.dst].push_back(stacks[m.src].back());
        stacks[m.src].pop_back();
      }
    }
    return answer(stacks);
  }

  std::string Part2(aoc::Input in) {
    std::vector<std::vector<char>> stacks;
    std::vector<Move> moves;
    parse(in, stacks, moves);
    for (auto& m : moves) {
      auto& src = stacks[m.src];
      auto& dst = stacks[m.dst];
      for (auto it = std::prev(src.end(), m.count); it < src.end(); ++it)
        dst.push_back(*it);
      for (int i = 0; i < m.count; ++i) src.pop_back();
    }
    return answer(stacks);
  }
};

TEST(Day05, Part1) {
  EXPECT_EQ(Day05().Part1(aoc::Input(2022, 5)), "QNNTGTPFN");
}

TEST(Day05, Part2) {
  EXPECT_EQ(Day05().Part2(aoc::Input(2022, 5)), "GGNPJBTTR");
}

}  // namespace aoc2022

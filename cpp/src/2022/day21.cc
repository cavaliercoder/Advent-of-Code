#include "lib/aoc.h"

namespace aoc2022 {

class Day21 {
  struct Monkey {
    std::string id;
    std::string lhs;
    std::string rhs;
    char op = 0;
    int64_t value = 0;

    int64_t eval(std::unordered_map<std::string, Monkey>& all) {
      if (!op) return value;
      auto a = all[lhs], b = all[rhs];
      switch (op) {
        case '+':
          return a.eval(all) + b.eval(all);
        case '-':
          return a.eval(all) - b.eval(all);
        case '*':
          return a.eval(all) * b.eval(all);
        case '/':
          return a.eval(all) / b.eval(all);
      }
      throw std::runtime_error("bad operator");
    }

    friend std::ostream& operator<<(std::ostream& os, Monkey& m) {
      if (!m.op) {
        return os << m.id << ": " << m.value;
      }
      return os << m.id << ": " << m.lhs << " " << m.op << " " << m.rhs;
    }
  };

  auto parse(aoc::Input& in) {
    std::unordered_map<std::string, Monkey> all;
    while (in) {
      Monkey m;
      in.get(m.id, 4).expect(": ");
      if (in.isdigit()) {
        in.get_uint(m.value).expect('\n');
      } else {
        in.get(m.lhs, 4).expect(' ').get(m.op).expect(' ');
        in.get_token(m.rhs).expect('\n');
      }
      all[m.id] = m;
    }
    return all;
  }

 public:
  int64_t Part1(aoc::Input in) {
    auto all = parse(in);
    return all["root"].eval(all);
  }

  int64_t Part2(aoc::Input in) { return 21; };
};

TEST(Day21, Part1) {
  EXPECT_EQ(Day21().Part1(aoc::Input(2022, 21)), 84244467642604);
}

TEST(Day21, Part2) { EXPECT_EQ(Day21().Part2(aoc::Input(2022, 21)), 21); }

}  // namespace aoc2022

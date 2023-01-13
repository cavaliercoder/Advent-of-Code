#include "lib/aoc.h"

namespace aoc2022 {

class Day21 {
  struct Monkey {
    std::string id;
    std::string lhs;
    std::string rhs;
    char op = 0;
    int64_t value = 0;

    int64_t eval(std::unordered_map<std::string, Monkey>& monkeys) {
      if (!op) return value;
      auto &a = monkeys[lhs], &b = monkeys[rhs];
      switch (op) {
        case '+':
          value = a.eval(monkeys) + b.eval(monkeys);
          break;
        case '-':
          value = a.eval(monkeys) - b.eval(monkeys);
          break;
        case '*':
          value = a.eval(monkeys) * b.eval(monkeys);
          break;
        case '/':
          value = a.eval(monkeys) / b.eval(monkeys);
          break;
      }
      return value;
    }

    // Return the left-hand side value required to eval to x.
    int64_t solve_left(const int64_t solution,
                       std::unordered_map<std::string, Monkey>& monkeys) const {
      auto b = monkeys[rhs].value;
      switch (op) {
        case '+':
          return solution - b;
        case '-':
          return solution + b;
        case '*':
          return solution / b;
        case '/':
          return solution * b;
      }
      return solution;
    }

    // Return the right-hand side value required to eval to x.
    int64_t solve_right(
        const int64_t solution,
        std::unordered_map<std::string, Monkey>& monkeys) const {
      auto a = monkeys[lhs].value;
      switch (op) {
        case '+':
          return solution - a;
        case '-':
          return a - solution;
        case '*':
          return solution / a;
        case '/':
          return a / solution;
      }
      return solution;
    }

    friend std::ostream& operator<<(std::ostream& os, const Monkey& m) {
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
    auto monkeys = parse(in);
    return monkeys["root"].eval(monkeys);
  }

  int64_t Part2(aoc::Input in) {
    int64_t answer;
    auto monkeys = parse(in);
    monkeys["root"].eval(monkeys);

    // DFS for the target expression, tracking the required solution to each
    // expression in the stack.  When we get to the target, we know exactly what
    // value it needs to represent.
    std::function<void(const std::string&, const std ::string&, const int64_t)>
        solve = [&](const std::string& start, const std::string& target,
                    const int64_t solution) {
          if (start == target) answer = solution;
          if (answer) return;
          auto& m = monkeys[start];
          if (!m.op) return;
          solve(m.lhs, target, m.solve_left(solution, monkeys));
          solve(m.rhs, target, m.solve_right(solution, monkeys));
        };

    // Instead of implementing equality, lets express it in terms of
    // subtraction: a - b == 0.
    monkeys["root"].op = '-';
    solve("root", "humn", 0);
    return answer;
  }
};

TEST(Day21, Part1) {
  EXPECT_EQ(Day21().Part1(aoc::Input(2022, 21)), 84244467642604);
}

TEST(Day21, Part2) {
  EXPECT_EQ(Day21().Part2(aoc::Input(2022, 21)), 3759569926192);
}

}  // namespace aoc2022

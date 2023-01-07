#include <algorithm>
#include <cassert>
#include <numeric>
#include <queue>

#include "lib/aoc.h"

namespace aoc2022 {

class Day11 {
  struct Monkey {
    std::queue<uint64_t> items;
    uint64_t score = 0;
    char op = 0;
    uint64_t arg = 0;
    uint64_t test = 0;
    int true_dst = 0;
    int false_dst = 0;

    std::pair<int, uint64_t> inspect(const uint64_t lcm = 0) {
      if (items.empty()) return {-1, -1};
      ++score;
      auto worry = items.front();
      items.pop();
      switch (op) {
        case '*':
          worry *= arg;
          break;
        case '+':
          worry += arg;
          break;
        case '^':
          worry = aoc::pow(worry, arg);
          break;
        default:
          assert(false);
          break;
      }
      if (lcm == 0)
        worry /= 3;
      else
        worry %= lcm;
      return {worry % test == 0 ? true_dst : false_dst, worry};
    }

    void push(const uint64_t item) { items.push(item); }

    uint64_t pop() {
      uint64_t result = items.front();
      items.pop();
      return result;
    }

    friend std::ostream& operator<<(std::ostream& os, Monkey* m) {
      os << "Monkey:\n";
      os << "  Starting items: ";
      for (int j = 0; j < m->items.size(); ++j) {
        if (j) os << ", ";
        os << m->items.front();
        m->items.push(m->items.front());
        m->items.pop();
      }
      os << "\n";
      os << "  Operation: new = old " << m->op << " " << m->arg << "\n";
      os << "  Test: divisible by " << m->test << "\n";
      os << "    If true: throw to monkey " << m->true_dst << "\n";
      os << "    If false: throw to monkey " << m->false_dst << "\n";
      os << "\n";
      return os;
    }

    static std::vector<Monkey*> parse(aoc::Input& in) {
      std::vector<Monkey*> result;
      Monkey* m;
      while (in) {
        in.expect("Monkey ").ignore().expect(":\n");
        m = new Monkey();
        result.push_back(m);
        in.expect("  Starting items: ");
        while (in) {
          m->push(in.get_uint<uint64_t>());
          if (in.branch('\n')) break;
          in.expect(", ");
        }
        in.expect("  Operation: new = old ").get(m->op).expect(' ');
        if (in.isdigit()) {
          in.get_uint(m->arg).expect('\n');
        } else {
          assert(m->op == '*');
          in.expect("old\n");
          m->op = '^';
          m->arg = 2;
        }
        in.expect("  Test: divisible by ").get_uint(m->test).expect('\n');
        in.expect("    If true: throw to monkey ")
            .get_uint(m->true_dst)
            .expect('\n');
        in.expect("    If false: throw to monkey ")
            .get_uint(m->false_dst)
            .expect('\n');
        in.discard('\n');
      }
      return result;
    }
  };

  uint64_t run(std::vector<Monkey*> monkeys, const int n,
               const uint64_t lcm = 0) {
    for (int i = 0; i < n; ++i) {
      for (auto m : monkeys) {
        while (true) {
          auto [dst, worry] = m->inspect(lcm);
          if (dst < 0) break;
          monkeys[dst]->push(worry);
        }
      }
    }
    std::vector<uint64_t> scores;
    for (auto m : monkeys) {
      scores.push_back(m->score);
      delete m;
    }
    std::sort(scores.begin(), scores.end(), std::greater<uint64_t>());
    return scores[0] * scores[1];
  }

 public:
  uint64_t Part1(aoc::Input in, const int n = 20) {
    return run(Monkey::parse(in), n);
  }

  uint64_t Part2(aoc::Input in, const int n = 10000) {
    auto monkeys = Monkey::parse(in);
    uint64_t lcm = 1;
    for (auto m : monkeys) lcm *= m->test;
    return run(monkeys, n, lcm);
  }
};

TEST(Day11, Part1) { EXPECT_EQ(Day11().Part1(aoc::Input(2022, 11)), 56120); }

TEST(Day11, Part2) {
  EXPECT_EQ(Day11().Part2(aoc::Input(2022, 11)), 24389045529);
}

}  // namespace aoc2022
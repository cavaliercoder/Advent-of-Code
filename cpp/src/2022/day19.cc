#include <iostream>
#include <stack>

#include "lib/aoc.h"

namespace aoc2022 {

class Day19 {
 public:
  enum Type : int { Ore = 0, Clay = 1, Obsidian = 2, Geodes = 3 };

  struct Bill {
    int resources[4];

    int operator[](const int i) const { return resources[i]; }
    int& operator[](const int i) { return resources[i]; }

#define BILL_MATH(op)                                         \
  friend Bill operator op(const Bill& lhs, const Bill& rhs) { \
    Bill b;                                                   \
    for (int i = 0; i < 4; ++i)                               \
      b.resources[i] = lhs.resources[i] op rhs.resources[i];  \
    return b;                                                 \
  }
    BILL_MATH(+)
    BILL_MATH(-)

#define BILL_CMP(op)                                             \
  friend bool operator op(const Bill& lhs, const Bill& rhs) {    \
    for (int i = 0; i < 4; ++i)                                  \
      if (!(lhs.resources[i] op rhs.resources[i])) return false; \
    return true;                                                 \
  }
    BILL_CMP(==)
    BILL_CMP(>=)
  };

  struct Blueprint {
    int id;
    Bill costs[4];
    Bill limits;

    Bill& operator[](const int i) { return costs[i]; }
  };

  struct State {
    Bill balance;
    Bill yield;
    int ttl;

    friend bool operator==(const State& lhs, const State& rhs) {
      return lhs.balance == rhs.balance && lhs.yield == rhs.yield &&
             lhs.ttl == rhs.ttl;
    }
  };

  std::vector<Blueprint> parse(aoc::Input& in) {
    std::vector<Blueprint> blueprints;
    while (in) {
      Blueprint b = {};
      in.expect("Blueprint ").get_uint(b.id);
      in.expect(": Each ore robot costs ").get_uint(b[Ore][Ore]);
      in.expect(" ore. Each clay robot costs ").get_uint(b[Clay][Ore]);
      in.expect(" ore. Each obsidian robot costs ")
          .get_uint(b[Obsidian][Ore])
          .expect(" ore and ")
          .get_uint(b[Obsidian][Clay]);
      in.expect(" clay. Each geode robot costs ")
          .get_uint(b[Geodes][Ore])
          .expect(" ore and ")
          .get_uint(b[Geodes][Obsidian]);
      in.expect(" obsidian.\n");

      // Compute maximum desired robot counts.
      for (int i = 0; i < 4; ++i) {
        for (int j = 0; j < 4; ++j) {
          b.limits[i] = std::max(b.limits[i], b.costs[j][i]);
        }
      }
      b.limits[Geodes] = INT_MAX;
      blueprints.push_back(b);
    }
    return blueprints;
  }

  int crack_geodes(Blueprint& bp, const int ttl) {
    int best = 0;
    auto Q = aoc::Stack<State>({{}, {1}, ttl});
    while (Q) {
      auto state = Q.pop();
      best = std::max(best, state.balance[Geodes]);
      if (!state.ttl) continue;

      // Prune branches that could never improve on the best score.
      int maximal = state.balance[Geodes] + (state.ttl * state.yield[Geodes]) +
                    ((state.ttl * state.ttl) / 2);
      if (maximal < best) continue;

      // Build each robot when next available.
      for (int i = 0; i < 4; ++i) {
        if (state.yield[i] >= bp.limits[i]) continue;
        Bill cost = bp.costs[i];
        State next = state;
        while (next.ttl) {
          if (next.balance >= cost) {
            --next.ttl;
            next.balance = next.balance + next.yield - cost;
            ++next.yield[i];
            Q.push(next);
            break;
          }
          --next.ttl;
          next.balance = next.balance + next.yield;
        }
      }
    }
    return best;
  }

 public:
  int Part1(aoc::Input in) {
    auto blueprints = parse(in);
    int sum = 0;
    for (auto bp : blueprints) {
      int best = crack_geodes(bp, 24);
      sum += (best * bp.id);
    }
    return sum;
  }

  int Part2(aoc::Input in) {
    auto blueprints = parse(in);
    int sum = 1;
    for (int i = 0; i < 3; ++i) {
      auto bp = blueprints[i];
      sum *= crack_geodes(bp, 32);
    }
    return sum;
  };
};

TEST(Day19, Part1) { EXPECT_EQ(Day19().Part1(aoc::Input(2022, 19)), 1681); }
TEST(Day19, Part2) { EXPECT_EQ(Day19().Part2(aoc::Input(2022, 19)), 5394); }

}  // namespace aoc2022

#include <unordered_map>

#include "lib/aoc.h"

namespace aoc2022 {

class Day16 {
  // Represents a valve as a unique bit, or multiple valves as a bitfield.
  // Must be wider that the total number of valves.
  using Mask = uint64_t;

  static constexpr Mask AA = 1;

  std::unordered_map<std::string, Mask> masks_ = {{"AA", AA}};
  Mask last_ = AA;

  // Maps a valve label to a unique bit mask.
  Mask get_mask(std::string valve) {
    if (masks_.count(valve)) return masks_[valve];
    last_ <<= 1;
    masks_.insert({valve, last_});
    return last_;
  }

  void parse(aoc::Input& in, std::unordered_map<Mask, int>& flows,
             std::unordered_map<Mask, std::unordered_map<Mask, int>>& edges) {
    std::string valve;
    Mask mask;
    int rate;
    while (in) {
      in.expect("Valve ")
          .get(valve, 2)
          .expect(" has flow rate=")
          .get_uint(rate);
      mask = get_mask(valve);
      edges[mask] = {};
      if (rate) flows[mask] = rate;
      in.expect("; tunnel");
      if (in.peek() == ' ')
        in.expect(" leads to valve ");
      else
        in.expect("s lead to valves ");
      while (in) {
        in.get(valve, 2);
        edges[mask][get_mask(valve)] = 1;
        if (in.branch('\n')) break;
        in.expect(", ");
      }
    }
  }

  // Returns the best flow possible for each possible set valves that can be
  // visited within the TTL.
  std::unordered_map<Mask, int> get_best_flows(aoc::Input& in, const int ttl) {
    std::unordered_map<Mask, int> flows;
    std::unordered_map<Mask, std::unordered_map<Mask, int>> edges;
    parse(in, flows, edges);

    // Compute lowest path weight between all valves.
    auto weights = aoc::floyd_warshall<Mask>(edges);

    // DFS all possible states for the best flow for each.
    struct State {
      Mask current = 0;
      Mask visited = 0;
      int ttl = 0;
      int flow = 0;
    };
    int seen = 1;
    std::stack<State> stack;
    stack.push(State{AA, 0, ttl, 0});
    std::unordered_map<Mask, int> best_flows;
    while (!stack.empty()) {
      auto state = stack.top();
      stack.pop();
      best_flows[state.visited] =
          std::max(best_flows[state.visited], state.flow);
      for (auto& [next, _] : flows) {
        auto next_ttl = state.ttl - weights[state.current][next] - 1;
        if (next_ttl <= 0 || state.visited & next) continue;
        ++seen;
        stack.push(State{next, state.visited | next, next_ttl,
                         state.flow + next_ttl * flows[next]});
      }
    }
    return best_flows;
  }

 public:
  int Part1(aoc::Input in) {
    int best = 0;
    auto best_flows = get_best_flows(in, 30);
    for (auto& [_, flow] : best_flows) best = std::max(best, flow);
    return best;
  }

  int Part2(aoc::Input in) {
    int best = 0;
    auto best_flows = get_best_flows(in, 26);
    for (auto& [visited1, flow1] : best_flows) {    // Me
      for (auto& [visited2, flow2] : best_flows) {  // Elephant
        if (visited1 & visited2) continue;          // Collision
        best = std::max(best, flow1 + flow2);
      }
    }
    return best;
  }
};

TEST(Day16, Part1) { EXPECT_EQ(Day16().Part1(aoc::Input(2022, 16)), 1659); }
TEST(Day16, Part2) { EXPECT_EQ(Day16().Part2(aoc::Input(2022, 16)), 2382); }

}  // namespace aoc2022

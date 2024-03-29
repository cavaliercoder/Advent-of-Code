#include <algorithm>
#include <climits>

#include "lib/aoc.h"

namespace aoc2022 {

class Day24 {
  using Grid = aoc::Grid<char>;
  using Point = aoc::Point<2, int>;

  struct Blizzard {
    Point init;
    Point vec;

    Blizzard(const Point init, const char id) : init(init) {
      switch (id) {
        case '^':
          vec = {0, -1};
          break;
        case 'v':
          vec = {0, 1};
          break;
        case '<':
          vec = {-1, 0};
          break;
        case '>':
          vec = {1, 0};
          break;
      }
    }
  };

  struct State {
    Point p;
    int t;
    int priority;

    bool operator<(const State& v) const { return priority < v.priority; }
    bool operator>(const State& v) const { return priority > v.priority; }
    bool operator==(const State& v) const { return p == v.p && t == v.t; }

    struct Hash {
      size_t operator()(const State& v) const {
        std::size_t h = std::hash<Point>()(v.p);
        h ^= (std::hash<int>()(v.t) << 1);
        return h;
      }
    };
  };

  Point S;
  Point E;
  Point bounds;
  aoc::Grid<char> grid;
  std::vector<aoc::Grid<char>> grid_cache = {};
  std::vector<Blizzard> blizzards;

  const Point move[5] = {{0, 0}, {-1, 0}, {0, -1}, {0, 1}, {1, 0}};

  // Fast-forward the position of a blizzard to time t.
  Point ff(const Blizzard& b, const int t) const {
    auto init = b.init - 1;   // Less top-left walls
    auto field = bounds - 2;  // Inner field size
    auto delta = Point(t % field.x(), t % field.y()) * b.vec;
    return ((init + delta + field) % field) + 1;
  }

  // Returns true if point p collides with a blizzard at time t.
  bool collides(const Point p, const int t) {
    // Caching grid states trades memory for a ~50% speed increase.
    //
    // There should only be W*H possible states but we find our destination
    // within W+H so there seems to be no point applying t%(W*H).
    if (t >= grid_cache.size()) {
      grid_cache.resize(std::max<size_t>(t * 2, 64), Grid());
    }
    auto& g = grid_cache[t - 1];
    if (g.empty()) {
      g = Grid(grid);
      for (auto& b : blizzards) g[ff(b, t)] = 'B';
    }
    return g[p] != '.';
  }

 public:
  Day24(aoc::Input in) {
    grid = in.grid();
    bounds = Point(grid.width(), grid.height());
    S = {1, 0};
    E = {bounds.x() - 2, bounds.y() - 1};
    for (auto it = grid.begin(); it < grid.end(); ++it) {
      if (*it == '#' || *it == '.') continue;
      auto b = Blizzard(it.point(), *it);
      blizzards.push_back(b);
      grid[it] = '.';
    }
  }

  int search(const Point src, const Point dst, const int T = 0) {
    auto dist = aoc::Grid<int>(bounds.x(), bounds.y(), INT_MAX);
    dist[src] = T;

    auto best = dist[dst];
    auto init = State{src, dist[src]};
    auto seen = aoc::Set<State, State::Hash>(init);
    auto heap = aoc::Heap<State, std::greater<State>>(init);
    while (heap) {
      const auto state = heap.pop();
      if (state.t > best) continue;
      for (int i = 0; i < 5; ++i) {
        int t = state.t + 1;
        Point p = state.p + move[i];
        if (!grid.contains(p) || grid[p] == '#') continue;
        if (collides(p, t)) continue;
        if (dist[p] > t) {
          dist[p] = t;
          if (p == dst) best = t;
        }
        if (p == dst) continue;
        State next = {p, t, (t * grid.size()) + p.manhattan(dst)};
        if (seen.contains(next)) continue;
        seen.insert(next);
        heap.push(next);
      }
    }
    return dist[dst];
  }

  int Part1() { return search(S, E); }

  int Part2() {
    auto t = search(S, E);
    t = search(E, S, t);
    t = search(S, E, t);
    return t;
  };
};

TEST(Day24, Part1) { EXPECT_EQ(Day24(aoc::Input(2022, 24)).Part1(), 343); }
TEST(Day24, Part2) { EXPECT_EQ(Day24(aoc::Input(2022, 24)).Part2(), 960); }

}  // namespace aoc2022

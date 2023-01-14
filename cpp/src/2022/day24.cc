#include <algorithm>

#include "lib/aoc.h"

namespace aoc2022 {

class Day24 {
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
  std::vector<std::vector<Blizzard>> blizzards_at_x;
  std::vector<std::vector<Blizzard>> blizzards_at_y;

  const Point move[5] = {{0, 0}, {-1, 0}, {0, -1}, {0, 1}, {1, 0}};

  // Fast-forward the position of a blizzard to time t.
  Point ff(const Blizzard& b, const int t) const {
    auto init = b.init - 1;   // Less top-left walls
    auto field = bounds - 2;  // Inner field size
    auto delta = Point(t % field.x(), t % field.y()) * b.vec;
    return ((init + delta + field) % field) + 1;
  }

  // Returns true if point p collides with a blizzard at time t.
  bool collides(const Point p, const int t) const {
    for (auto& b : blizzards_at_x[p.x()])
      if (p == ff(b, t)) return true;
    for (auto& b : blizzards_at_y[p.y()])
      if (p == ff(b, t)) return true;
    return false;
  }

 public:
  Day24(aoc::Input in) {
    grid = in.grid();
    bounds = Point(grid.width(), grid.height());
    S = {1, 0};
    E = {bounds.x() - 2, bounds.y() - 1};
    blizzards_at_x.resize(bounds.x());
    blizzards_at_y.resize(bounds.y());
    for (auto it = grid.begin(); it < grid.end(); ++it) {
      if (*it == '#' || *it == '.') continue;
      auto b = Blizzard(it.point(), *it);
      if (!b.vec.x()) blizzards_at_x[b.init.x()].push_back(b);
      if (!b.vec.y()) blizzards_at_y[b.init.y()].push_back(b);
    }
  }

  int search(const Point src, const Point dst, const int T = 0) const {
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
        State next = {p, t, t + p.manhattan(dst)};
        if (seen.contains(next)) continue;
        seen.insert(next);
        heap.push(next);
      }
    }
    return dist[dst];
  }

  int Part1() const { return search(S, E); }

  int Part2() const {
    auto t = search(S, E);
    t = search(E, S, t);
    t = search(S, E, t);
    return t;
  };
};

TEST(Day24, Part1) { EXPECT_EQ(Day24(aoc::Input(2022, 24)).Part1(), 343); }
TEST(Day24, Part2) { EXPECT_EQ(Day24(aoc::Input(2022, 24)).Part2(), 960); }

}  // namespace aoc2022

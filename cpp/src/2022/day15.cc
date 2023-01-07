#include <stack>

#include "lib/aoc.h"

namespace aoc2022 {

class Day15 {
  using Point = aoc::Point<2, int64_t>;
  using Rect = aoc::Rect<int64_t>;
  using Range = aoc::Range<int64_t>;

  struct Object {
    enum Type : char {
      Empty = ' ',
      Sensor = 'S',
      Beacon = 'B',
    };

    Type key;
    Point pos;
    Point beacon;
    int radius;

    Object(const Type key, const Point pos, const Point beacon = Point())
        : key(key), pos(pos), beacon(beacon), radius(pos.manhattan(beacon)) {}

    Object(const Point sensor, const Point beacon)
        : Object(Sensor, sensor, beacon) {}

    explicit Object(const Point beacon) : Object(Beacon, beacon) {}

    inline bool is_sensor() const { return key == Sensor; }
    inline bool is_beacon() const { return key == Beacon; }

    inline bool in_range(const Point p) const {
      if (is_beacon()) return pos == p;
      return pos.manhattan(p) <= radius;
    }

    bool in_range(const Rect& r) const {
      return in_range(r.tl()) && in_range(r.tr() + Point(-1, 0)) &&
             in_range(r.bl() + Point(0, -1)) &&
             in_range(r.br() + Point(-1, -1));
    }

    Range range_y(const int64_t y) const {
      int64_t delta_y = std::abs(pos.y() - y);
      if (delta_y > radius) return {};
      int64_t delta_x = radius - delta_y;
      return {pos.x() - delta_x, pos.x() + delta_x + 1};
    }
  };

  using Space = aoc::Space<2, int64_t, Object*>;

  Space parse(aoc::Input& in) {
    int x, y;
    auto m = Space(nullptr);
    while (in) {
      in.expect("Sensor at x=").get_int(x).expect(", y=").get_int(y);
      auto sensor = Point(x, y);
      in.expect(": closest beacon is at x=")
          .get_int(x)
          .expect(", y=")
          .get_int(y)
          .expect('\n');
      auto beacon = Point(x, y);
      if (!m.contains(beacon)) m.insert(beacon, new Object(beacon));
      assert(!m.contains(sensor));
      m.insert(sensor, new Object(sensor, beacon));
    }
    return m;
  }

  void delete_space(Space& m) {
    for (auto [_, obj] : m) {
      delete obj;
    }
  }

 public:
  uint64_t Part1(aoc::Input in, const int64_t target_y = 2000000) {
    auto m = parse(in);
    std::vector<Range> ranges;
    for (auto [p, obj] : m) {
      if (obj->is_beacon()) continue;
      ranges.push_back(obj->range_y(target_y));
    }
    Range::flatten(ranges);
    int count = 0;
    for (auto& r : ranges) {
      count += r.size();
      for (auto [p, _] : m) {
        if (p.y() != target_y) continue;
        if (r.contains(p.x())) --count;
      }
    }
    delete_space(m);
    return count;
  }

  uint64_t Part2(aoc::Input in, const int64_t max_v = 4000000) {
    auto m = parse(in);
    int64_t score = -1;
    auto stack = std::stack<Rect>();
    stack.push(Rect(Point(), Point(max_v + 1, max_v + 1)));
    while (!stack.empty()) {
      auto r = stack.top();
      stack.pop();
      bool skip = false;
      for (auto [_, obj] : m) {
        if (obj->in_range(r)) {
          skip = true;
          break;
        }
      }
      if (skip) continue;
      if (r.size() == 1) {
        auto p = r.tl();
        score = (p.x() * max_v) + p.y();
        break;
      }
      for (auto subr : r.split()) stack.push(subr);
    }
    delete_space(m);
    return score;
  }
};

TEST(Day15, Part1) { EXPECT_EQ(Day15().Part1(aoc::Input(2022, 15)), 5256611); }

TEST(Day15, Part2) {
  EXPECT_EQ(Day15().Part2(aoc::Input(2022, 15)), 13337919186981);
}

}  // namespace aoc2022

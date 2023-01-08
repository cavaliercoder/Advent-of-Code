#ifndef AOC_SPACE_H
#define AOC_SPACE_H

#include <unordered_map>

#include "point.h"

namespace aoc {

// Space is a container of elements arranged in an infinite N-dimensional
// space.
template <int N, typename K, typename V>
class Space {
  using Point = Point<N, K>;
  using Map = std::unordered_map<Point, V>;

  V zero_;
  Map map_;
  Point min_;
  Point max_;

 public:
  explicit Space(const V zero = V()) : zero_(zero) {}
  inline int size() const { return map_.size(); }
  inline bool empty() const { return map_.empty(); }
  inline Point min() const { return min_; }
  inline Point max() const { return max_; }
  inline K width() const { return max_.x() - min_.x() + 1; }
  inline K height() const { return max_.y() - min_.y() + 1; }
  inline K depth() const { return max_.z() - min_.z() + 1; }
  inline bool contains(const Point p) const { return map_.count(p); }
  inline auto begin() const { return map_.begin(); }
  inline auto end() const { return map_.end(); }

  inline int count(const Point p) const { return map_.count(p); }

  int count(const V value) const {
    int count = 0;
    for (auto& [_, v] : map_) {
      if (v == value) ++count;
    }
    return count;
  }

  void clear() {
    map_.clear();
    min_ = {}, max_ = {};
  }

  V at(const Point p) const {
    if (!contains(p)) return zero_;
    return map_.at(p);
  }

  void insert(const std::pair<const Point, V> value) {
    map_.insert(value);
    min_ = min_.min(value.first);
    max_ = max_.max(value.first);
  }

  inline void insert(const Point p, V value) { return insert({p, value}); }

  inline V operator[](const Point p) const { return at(p); }

  inline V& operator[](const Point p) {
    if (!contains(p)) insert(p, zero_);
    return map_.at(p);
  }
};

}  // namespace aoc

#endif  // AOC_SPACE_H

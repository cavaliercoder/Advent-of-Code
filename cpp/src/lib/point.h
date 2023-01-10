#ifndef AOC_POINT_H
#define AOC_POINT_H

#include <array>
#include <iostream>

namespace aoc {

// Point represents an N-dimensional coordinate with components of type T.
template <int N = 2, typename T = int>
struct Point {
  T data[N] = {};

  // Returns a point at {0, ...}.
  Point() = default;

  // Returns a point with the given components.
  Point(std::initializer_list<T> init) {
    assert(init.size() == N);
    std::copy(init.begin(), init.end(), data);
  }

  // Returns a 2-dimensional point at {x, y}.
  Point<2, T>(const int x, const int y) : Point{x, y} {};

  // Returns a 3-dimensional point at {x, y, z}.
  Point<3, T>(const T x, const T y, const T z) : Point{x, y, z} {}

  // Returns the x-component (width) of a point with at least one dimension.
  inline T x() const { return data[0]; }

  // Returns the y-component (height) of a point with at least two dimensions.
  inline T y() const { return data[1]; }

  // Returns the z-component (depth) of a point with at least three dimensions.
  inline T z() const { return data[2]; }

  // Returns true if all components are zero.
  bool empty() const {
    for (int i = 0; i < N; ++i)
      if (data[i]) return false;
    return true;
  }

  // Returns a new point where each component is the result of f(n).
  Point map(const std::function<T(T)>& f) const {
    Point p = {};
    for (int i = 0; i < N; ++i) p.data[i] = f(data[i]);
    return p;
  }

  // Returns a new point where each component is the result of f(n, m).
  Point map(const Point p, const std::function<T(T, T)>& f) const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = f(data[i], p.data[i]);
    return q;
  }

  // Returns a new point where each component is the absolute value of its
  // old value.
  inline Point abs() const {
    return map([](const T n) -> T { return std::abs(n); });
  }

  // Returns a new point where each component is the lower of the same component
  // in this point and p.
  inline Point min(const Point p) const {
    return map(p, [](const T a, const T b) -> T { return std::min(a, b); });
  }

  // Returns a new point where each component is the higher of the same
  // component in this point and p.
  inline Point max(const Point p) const {
    return map(p, [](const T a, const T b) -> T { return std::max(a, b); });
  }

  // Returns the Manhattan distance from this point to p.
  inline T manhattan(const Point p) const {
    T n = 0;
    for (int i = 0; i < N; ++i) n += std::abs(data[i] - p.data[i]);
    return n;
  }

  // Increment each component toward its complement in p.
  //
  // Once each component reaches its complement in p, it is no longer
  // incremented.
  //
  // Calling p.nudge(p) always returns p.
  Point nudge(const Point p) const {
    return map(p, [](const T a, const T b) -> T {
      if (a < b) return std::min(a + 1, b);
      if (a > b) return std::max(a - 1, b);
      return a;
    });
  }

#define POINT_MOVE(name, index, op)    \
  inline Point name(const T n) const { \
    Point p = *this;                   \
    p.data[index] op## = n;            \
    return p;                          \
  }                                    \
                                       \
  inline Point name() const {          \
    Point p = *this;                   \
    op##op p.data[index];              \
    return p;                          \
  }

  POINT_MOVE(left, 0, -)
  POINT_MOVE(right, 0, +)
  POINT_MOVE(up, 1, +)
  POINT_MOVE(down, 1, -)
  POINT_MOVE(forward, 2, +)
  POINT_MOVE(backward, 2, -)

  // Returns the adjacent points immediately up, down, left and right.
  std::array<Point, 4> udlr() const { return {up(), down(), left(), right()}; }

  // Returns the adjacent points immediately up, down, left, right, forward and
  // back.
  std::array<Point, 6> udlrfb() const {
    return {up(), down(), left(), right(), forward(), backward()};
  }

  // Negates each component.
  inline Point operator-() const {
    return map([](const T n) -> T { return -n; });
  }

  // Returns true if any component is non-zero.
  inline operator bool() { return !empty(); }

#define POINT_ARITHMATIC(op)                                               \
  inline friend Point operator op(const Point lhs, const Point rhs) {      \
    return lhs.map(rhs, [](const T a, const T b) -> T { return a op b; }); \
  }                                                                        \
                                                                           \
  inline friend Point operator op(const Point lhs, const T rhs) {          \
    return lhs.map([rhs](const T n) -> T { return n op rhs; });            \
  }                                                                        \
                                                                           \
  inline friend Point operator op##=(Point& lhs, const Point rhs) {        \
    return lhs = lhs op rhs;                                               \
  }                                                                        \
                                                                           \
  inline friend Point operator op##=(Point& lhs, const T rhs) {            \
    return lhs = lhs op rhs;                                               \
  }

  POINT_ARITHMATIC(+)
  POINT_ARITHMATIC(-)
  POINT_ARITHMATIC(*)
  POINT_ARITHMATIC(/)
  POINT_ARITHMATIC(%)

#define POINT_INCREMENT(op)                          \
  inline friend Point operator op(Point& p) {        \
    return p = p.map([](T n) -> T { return op n; }); \
  }                                                  \
                                                     \
  inline friend Point operator op(Point& p, int) {   \
    Point old = p;                                   \
    p = p.map([](T n) -> T { return op n; });        \
    return old;                                      \
  }

  POINT_INCREMENT(--)
  POINT_INCREMENT(++)

#define POINT_CMP(op)                                                \
  inline friend bool operator op(const Point lhs, const Point rhs) { \
    for (int i = 0; i < N; ++i) {                                    \
      if (lhs.data[i] op rhs.data[i]) continue;                      \
      return false;                                                  \
    }                                                                \
    return true;                                                     \
  }

  POINT_CMP(==)
  POINT_CMP(!=)
  POINT_CMP(<)
  POINT_CMP(>)
  POINT_CMP(<=)
  POINT_CMP(>=)

  friend std::ostream& operator<<(std::ostream& os, const Point p) {
    os << "{";
    for (int i = 0; i < N; ++i) {
      if (i) os << ", ";
      os << p.data[i];
    }
    return os << "}";
  }
};

}  // namespace aoc

namespace std {

// Hasher for aoc::Point so it can be used in std::unordered_set.
template <int N, typename T>
struct std::hash<aoc::Point<N, T>> {
  std::size_t operator()(const aoc::Point<N, T>& p) const noexcept {
    std::size_t h = 0;
    for (int i = 0; i < N; ++i) h ^= (std::hash<int>()(p.data[i]) << i);
    return h;
  }
};

}  // namespace std

#endif  // AOC_POINT_H

#ifndef AOC_POINT_H
#define AOC_POINT_H

#include <array>
#include <iostream>

namespace aoc {

// Point represents an N-dimensional coordinate with components of type T.
template <size_t N = 2, typename T = int>
struct Point {
  T data[N] = {};

  // Returns a point at {0, ...}.
  constexpr Point() = default;

  // Returns a point with the given components.
  constexpr Point(std::initializer_list<T> init) {
    assert(init.size() == N);
    int i = 0;
    for (auto n : init) data[i++] = n;
  }

  // Returns a 2-dimensional point at {x, y}.
  constexpr Point<2, T>(const int x, const int y) : Point{x, y} {};

  // Returns a 3-dimensional point at {x, y, z}.
  constexpr Point<3, T>(const T x, const T y, const T z) : Point{x, y, z} {}

  // Returns the x-component (width) of a point with at least one dimension.
  constexpr T x() const { return data[0]; }

  // Returns the y-component (height) of a point with at least two dimensions.
  constexpr T y() const { return data[1]; }

  // Returns the z-component (depth) of a point with at least three dimensions.
  constexpr T z() const { return data[2]; }

  // Returns the number of dimensions in the point.
  constexpr size_t size() const { return N; }

  // Returns true if all components are zero.
  bool constexpr empty() const {
    for (int i = 0; i < N; ++i)
      if (data[i]) return false;
    return true;
  }

  // Returns a new point where each component is the result of f(n).
  Point map(const std::function<T(T)>& f) const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = f(data[i]);
    return q;
  }

  // Returns a new point where each component is the result of f(n, m).
  Point map(const Point p, const std::function<T(T, T)>& f) const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = f(data[i], p.data[i]);
    return q;
  }

  // Returns a new point where each component is the absolute value of its
  // old value.
  constexpr Point abs() const {
    Point q = {};
    for (int i = 0; i < N; ++i) {
      // std::abs is not a constexpr until C++23
      q.data[i] = data[i] < 0 ? -data[i] : data[i];
    }
    return q;
  }

  // Returns a new point where each component is either -1, 0 or 1 - whichever
  // is closest. For example, the orientation of {-10, 0 53} is {-1, 0, 1}.
  constexpr Point orientation() const {
    Point q = {};
    for (int i = 0; i < N; ++i)
      q.data[i] = (data[i] == 0) ? 0 : (data[i] > 0) ? 1 : -1;
    return q;
  }

  constexpr Point orientation(const Point toward) const {
    return (toward - *this).orientation();
  }

  // Returns a new point where each component is the lower of the same component
  // in this point and p.
  constexpr Point min(const Point p) const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = std::min(data[i], p.data[i]);
    return q;
  }

  // Returns a new point where each component is the higher of the same
  // component in this point and p.
  constexpr Point max(const Point p) const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = std::max(data[i], p.data[i]);
    return q;
  }

  // Returns the Manhattan distance from this point to p.
  constexpr T manhattan(const Point p = {}) const {
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
  constexpr Point nudge(const Point p) const {
    Point q = {};
    for (int i = 0; i < N; ++i) {
      q.data[i] = (data[i] == p[i])  ? data[i]
                  : (data[i] < p[i]) ? std::min(data[i] + 1, p[i])
                                     : std::max(data[i] - 1, p[i]);
    }
    return q;
  }

  // Returns the point rotated 90° clockwise around {0, 0}.
  constexpr Point<2, T> cw() const { return Point(y(), -x()); }

  // Returns the point rotated 90° counterclockwise around {0, 0.
  constexpr Point<2, T> ccw() const { return Point(-y(), x()); }

  // Returns the point on the edge of a square of width w, rotated -90° in place
  // to the next clockwise edge.
  constexpr Point<2, T> icw(const T w) const {
    assert(w >= 0 && x() >= 0 && x() <= w && y() >= 0 && y() <= w);
    return Point(y(), w - x() - 1);
  }

  // Returns the point on the edge of a square of width w, rotated 90° in place
  // to the next counterclockwise edge.
  constexpr Point<2, T> iccw(const T w) const {
    assert(w >= 0 && x() >= 0 && x() <= w && y() >= 0 && y() <= w);
    return Point(w - y() - 1, x());
  }

#define POINT_MOVE(name, index, op)       \
  constexpr Point name(const T n) const { \
    Point p = *this;                      \
    p.data[index] op## = n;               \
    return p;                             \
  }                                       \
                                          \
  constexpr Point name() const {          \
    Point p = *this;                      \
    op##op p.data[index];                 \
    return p;                             \
  }

  POINT_MOVE(left, 0, -)
  POINT_MOVE(right, 0, +)
  POINT_MOVE(up, 1, +)
  POINT_MOVE(down, 1, -)
  POINT_MOVE(forward, 2, +)
  POINT_MOVE(backward, 2, -)

  // Returns all immediate orthogonal neighboring points.
  //
  // The returned order is deterministic and is always -1 then +1 for each
  // subsequent dimension.
  constexpr std::array<Point, N * 2> orth() const {
    std::array<Point, N * 2> a;
    for (int i = 0; i < N; ++i) {
      auto p = Point(*this);
      p.data[i] = data[i] - 1;
      a[i * 2] = p;
      p.data[i] = data[i] + 1;
      a[(i * 2) + 1] = p;
    }
    return a;
  }

  // Negates each component.
  constexpr Point operator-() const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = -data[i];
    return q;
  }

  // Returns true if any component is non-zero.
  constexpr operator bool() { return !empty(); }

  constexpr T& operator[](const int i) { return data[i]; }
  constexpr T operator[](const int i) const { return data[i]; }

#define POINT_ARITHMATIC(op)                                             \
  constexpr friend Point operator op(const Point lhs, const Point rhs) { \
    Point q = {};                                                        \
    for (int i = 0; i < N; ++i) q.data[i] = lhs.data[i] op rhs.data[i];  \
    return q;                                                            \
  }                                                                      \
                                                                         \
  constexpr friend Point operator op(const Point lhs, const T rhs) {     \
    Point q = {};                                                        \
    for (int i = 0; i < N; ++i) q.data[i] = lhs.data[i] op rhs;          \
    return q;                                                            \
  }                                                                      \
                                                                         \
  constexpr friend Point operator op##=(Point& lhs, const Point rhs) {   \
    return lhs = lhs op rhs;                                             \
  }                                                                      \
                                                                         \
  constexpr friend Point operator op##=(Point& lhs, const T rhs) {       \
    return lhs = lhs op rhs;                                             \
  }

  POINT_ARITHMATIC(+)
  POINT_ARITHMATIC(-)
  POINT_ARITHMATIC(*)
  POINT_ARITHMATIC(/)
  POINT_ARITHMATIC(%)

#define POINT_INCREMENT(op)                           \
  constexpr friend Point operator op(Point& p) {      \
    for (int i = 0; i < N; ++i) op p.data[i];         \
    return p;                                         \
  }                                                   \
                                                      \
  constexpr friend Point operator op(Point& p, int) { \
    Point old = p;                                    \
    for (int i = 0; i < N; ++i) op p.data[i];         \
    return old;                                       \
  }

  POINT_INCREMENT(--)
  POINT_INCREMENT(++)

#define POINT_CMP(op)                                                   \
  constexpr friend bool operator op(const Point lhs, const Point rhs) { \
    for (int i = 0; i < N; ++i) {                                       \
      if (lhs.data[i] op rhs.data[i]) continue;                         \
      return false;                                                     \
    }                                                                   \
    return true;                                                        \
  }

  POINT_CMP(==)
  POINT_CMP(<)
  POINT_CMP(>)
  POINT_CMP(<=)
  POINT_CMP(>=)

  constexpr friend bool operator!=(const Point lhs, const Point rhs) {
    for (int i = 0; i < N; ++i) {
      if (lhs.data[i] != rhs.data[i]) return true;
    }
    return false;
  }

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

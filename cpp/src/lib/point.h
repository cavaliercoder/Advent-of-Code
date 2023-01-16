#ifndef AOC_POINT_H
#define AOC_POINT_H

#include <array>
#include <cassert>
#include <functional>
#include <iostream>

#include "checksum.h"

namespace aoc {

// Point represents an N-dimensional coordinate with components of type T.
template <size_t Size = 2, typename T = int>
struct Point {
  T data[Size] = {};

  // FNV Hash runs faster in most cases and with a better collision rate than
  // std::hash.
  struct FNV32 {
    std::size_t operator()(const aoc::Point<Size, T>& p) const noexcept {
      return aoc::FNV32(&p.data[0], sizeof(T) * Size);
    }
  };

  // Returns a point at {0, ...}.
  constexpr Point() = default;

  // Returns a point with the given components.
  constexpr Point(std::initializer_list<T> init) {
    assert(init.size() == Size);
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
  constexpr size_t size() const { return Size; }

  // Returns true if all components are zero.
  bool constexpr empty() const {
    for (int i = 0; i < Size; ++i)
      if (data[i]) return false;
    return true;
  }

  // Returns a new point where each component is the result of f(n).
  Point map(const std::function<T(T)>& f) const {
    Point q = {};
    for (int i = 0; i < Size; ++i) q.data[i] = f(data[i]);
    return q;
  }

  // Returns a new point where each component is the result of f(n, m).
  Point map(const Point p, const std::function<T(T, T)>& f) const {
    Point q = {};
    for (int i = 0; i < Size; ++i) q.data[i] = f(data[i], p.data[i]);
    return q;
  }

  // Returns a new point where each component is the absolute value of its
  // old value.
  constexpr Point abs() const {
    Point q = {};
    for (int i = 0; i < Size; ++i) {
      // std::abs is not a constexpr until C++23
      q.data[i] = data[i] < 0 ? -data[i] : data[i];
    }
    return q;
  }

  // Returns a new point where each component is either -1, 0 or 1 - whichever
  // is closest. For example, the orientation of {-10, 0 53} is {-1, 0, 1}.
  constexpr Point orientation() const {
    Point q = {};
    for (int i = 0; i < Size; ++i)
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
    for (int i = 0; i < Size; ++i) q.data[i] = std::min(data[i], p.data[i]);
    return q;
  }

  // Returns a new point where each component is the higher of the same
  // component in this point and p.
  constexpr Point max(const Point p) const {
    Point q = {};
    for (int i = 0; i < Size; ++i) q.data[i] = std::max(data[i], p.data[i]);
    return q;
  }

  // Returns the Manhattan distance from this point to p.
  constexpr T manhattan(const Point p = {}) const {
    T n = 0;
    for (int i = 0; i < Size; ++i) n += std::abs(data[i] - p.data[i]);
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
    for (int i = 0; i < Size; ++i) {
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

  // Returns all immediate orthogonal neighboring points at distance n.
  //
  // The returned order is deterministic and is always -1 then +1 for each
  // subsequent dimension.
  constexpr std::array<Point, Size * 2> orth(const T n = 1) const {
    std::array<Point, Size * 2> a;
    for (int i = 0; i < Size; ++i) {
      auto p = Point(*this);
      p.data[i] = data[i] - n;
      a[i * 2] = Point(p);
      p.data[i] = data[i] + n;
      a[(i * 2) + 1] = Point(p);
    }
    return a;
  }

  // Returns all immediate adjacent neighboring points at distance n.
  // Ordering is determinate but an implementation detail and should not be
  // relied upon.
  constexpr std::vector<Point> adj(const T n = 1) const {
    Point p;
    std::vector<Point> a;
    for (int i = 0; i < Size; ++i) {
      auto qn = a.size();
      for (int q = 0; q < qn; ++q) {
        p = Point(a[q]);
        p.data[i] = data[i] - n;
        a.push_back(p);
        p.data[i] = data[i] + n;
        a.push_back(p);
      }
      p = Point(*this);
      p.data[i] = data[i] - n;
      a.push_back(p);
      p.data[i] = data[i] + n;
      a.push_back(p);
    }
    return a;
  }

#define POINT_ORTH(name, index, offset)       \
  constexpr Point name(const T n = 1) const { \
    Point p = *this;                          \
    p.data[index] += n * offset;              \
    return p;                                 \
  }

  POINT_ORTH(left, 0, -1)
  POINT_ORTH(right, 0, +1)
  POINT_ORTH(up, 1, +1)
  POINT_ORTH(down, 1, -1)
  POINT_ORTH(forward, 2, +1)
  POINT_ORTH(backward, 2, -1)

  POINT_ORTH(N, 1, +1)
  POINT_ORTH(S, 1, -1)
  POINT_ORTH(E, 0, +1)
  POINT_ORTH(W, 0, -1)

#define POINT_DIAG(name, X, Y)                \
  constexpr Point name(const T n = 1) const { \
    Point p = *this;                          \
    p.data[0] += X * n;                       \
    p.data[1] += Y * n;                       \
    return p;                                 \
  }

  POINT_DIAG(NW, -1, 1)
  POINT_DIAG(NE, 1, 1)
  POINT_DIAG(SW, -1, -1)
  POINT_DIAG(SE, 1, -1)

  // Negates each component.
  constexpr Point operator-() const {
    Point q = {};
    for (int i = 0; i < Size; ++i) q.data[i] = -data[i];
    return q;
  }

  // Returns true if any component is non-zero.
  constexpr operator bool() { return !empty(); }

  constexpr T& operator[](const int i) { return data[i]; }
  constexpr T operator[](const int i) const { return data[i]; }

  constexpr friend bool operator==(const Point lhs, const Point rhs) {
    for (int i = 0; i < Size; ++i) {
      if (lhs.data[i] != rhs.data[i]) return false;
    }
    return true;
  }

  constexpr friend bool operator!=(const Point lhs, const Point rhs) {
    for (int i = 0; i < Size; ++i) {
      if (lhs.data[i] != rhs.data[i]) return true;
    }
    return false;
  }

#define POINT_ARITHMATIC(op)                                               \
  constexpr friend Point operator op(const Point lhs, const Point rhs) {   \
    Point q = {};                                                          \
    for (int i = 0; i < Size; ++i) q.data[i] = lhs.data[i] op rhs.data[i]; \
    return q;                                                              \
  }                                                                        \
                                                                           \
  constexpr friend Point operator op(const Point lhs, const T rhs) {       \
    Point q = {};                                                          \
    for (int i = 0; i < Size; ++i) q.data[i] = lhs.data[i] op rhs;         \
    return q;                                                              \
  }                                                                        \
                                                                           \
  constexpr friend Point operator op##=(Point& lhs, const Point rhs) {     \
    return lhs = lhs op rhs;                                               \
  }                                                                        \
                                                                           \
  constexpr friend Point operator op##=(Point& lhs, const T rhs) {         \
    return lhs = lhs op rhs;                                               \
  }

  POINT_ARITHMATIC(+)
  POINT_ARITHMATIC(-)
  POINT_ARITHMATIC(*)
  POINT_ARITHMATIC(/)
  POINT_ARITHMATIC(%)

#define POINT_INCREMENT(op)                           \
  constexpr friend Point operator op(Point& p) {      \
    for (int i = 0; i < Size; ++i) op p.data[i];      \
    return p;                                         \
  }                                                   \
                                                      \
  constexpr friend Point operator op(Point& p, int) { \
    Point old = p;                                    \
    for (int i = 0; i < Size; ++i) op p.data[i];      \
    return old;                                       \
  }

  POINT_INCREMENT(--)
  POINT_INCREMENT(++)

#define POINT_CMP(op)                                                   \
  constexpr friend bool operator op(const Point lhs, const Point rhs) { \
    for (int i = 0; i < Size; ++i) {                                    \
      if (lhs.data[i] op rhs.data[i]) return true;                      \
      if (rhs.data[i] op lhs.data[i]) return false;                     \
    }                                                                   \
    return false;                                                       \
  }

  POINT_CMP(<)
  POINT_CMP(>)
  POINT_CMP(<=)
  POINT_CMP(>=)

  friend std::ostream& operator<<(std::ostream& os, const Point p) {
    os << "{";
    for (int i = 0; i < Size; ++i) {
      if (i) os << ", ";
      os << p.data[i];
    }
    return os << "}";
  }

  // Implements getter for structured binding (auto [x, y] = p).
  template <std::size_t I>
  std::tuple_element_t<I, Point> get() {
    return data[I];
  }
};

}  // namespace aoc

template <size_t Size, typename T>
struct std::hash<aoc::Point<Size, T>> : aoc::Point<Size, T>::FNV32 {};

// Implement structured binding to each component of a point:

template <size_t Size, typename T>
struct std::tuple_size<aoc::Point<Size, T>> {
  static constexpr size_t value = Size;
};

template <std::size_t I, size_t Size, typename T>
struct std::tuple_element<I, aoc::Point<Size, T>> {
  using type = T;
};

#endif  // AOC_POINT_H

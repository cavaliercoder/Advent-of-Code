#ifndef AOC_RECT_H
#define AOC_RECT_H

#include "point.h"

namespace aoc {

template <typename T>
struct Rect {
  using Point = Point<2, T>;

  Point a;
  Point b;

  Rect(const Point a, const Point b) : a(a), b(b) {}

  inline Point tl() const { return a; }
  inline Point tr() const { return Point(b.x(), a.y()); }
  inline Point bl() const { return Point(a.x(), b.y()); }
  inline Point br() const { return b; }
  inline int width() const { return std::abs(b.x() - a.x()); }
  inline int height() const { return std::abs(b.y() - a.y()); }
  inline int size() const { return width() * height(); }

  // Split into four quarters.
  std::array<Rect, 4> split() const {
    auto w = width() / 2;
    auto h = height() / 2;
    return {
        Rect(a, {a.x() + w, a.y() + h}),
        Rect({a.x() + w, a.y()}, {b.x(), a.y() + h}),
        Rect({a.x(), a.y() + h}, {a.x() + w, b.y()}),
        Rect({a.x() + w, a.y() + h}, b),
    };
  }
};

}  // namespace aoc

#endif  //  AOC_RECT_H

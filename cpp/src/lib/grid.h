#ifndef AOC_GRID_H
#define AOC_GRID_H

#include <sstream>
#include <vector>

#include "point.h"

namespace aoc {

// Grid is a container of elements arranged in a finite 2-dimensional space.
//
// Members may be accessed by index or x/y position. The y-axis increases as the
// index increases. I.e. the y-axis increases in a downward direction.
template <typename T = char>
class Grid {
  using Point = Point<2, int>;

  std::vector<T> data_;
  int width_ = 0;
  int height_ = 0;

  // Converts a point to an index.
  inline int ptoi(const Point p) const {
    if (!contains(p)) return size();
    return p.y() * width() + p.x();
  }

  // Converts an index to a point.
  inline Point itop(const int i) const {
    if (!contains(i)) return {width(), height()};
    return Point(i % width(), i / width());
  }

 public:
  Grid(const int width, const int height, std::vector<T> data)
      : width_(width), height_(height), data_(data) {
    assert(data_.size() == width * height);
  }

  Grid(const int width, const int height, T value)
      : Grid(width, height, std::vector(width * height, value)) {}

  inline int size() const { return data_.size(); }
  inline int width() const { return width_; }
  inline int height() const { return height_; }

  inline bool contains(const int i) const { return i >= 0 && i < size(); }

  inline bool contains(const Point p) const {
    return p.x() >= 0 && p.x() < width() && p.y() >= 0 && p.y() < height();
  }

  // Returns the count of members equal to value.
  inline int count(const T value) const {
    int n = 0;
    for (int i = 0; i < size(); ++i)
      if (data_[i] == value) ++n;
    return n;
  }

  std::string str() const {
    std::stringstream ss;
    ss << *this;
    return ss.str();
  }

  inline T operator[](const int i) const { return data_[i]; }
  inline T& operator[](const int i) { return data_[i]; }
  inline T operator[](const Point p) const { return data_[ptoi(p)]; }
  inline T& operator[](const Point p) { return data_[ptoi(p)]; }

  friend std::ostream& operator<<(std::ostream& os, const Grid& g) {
    for (int i = 0; i < g.size(); ++i) {
      os << g[i];
      if (i > 0 && i % g.width() == g.width() - 1) os << "\n";
    }
    return os;
  }

  class Iterator {
    const Grid* grid_;
    int index_ = 0;
    Point point_ = Point();

    inline void oob() {
      index_ = grid_->size();
      point_ = {grid_->width(), grid_->height()};
    }

   public:
    Iterator(const Grid* grid, const int index = 0) : grid_(grid) {
      set(index);
    }

    Iterator(const Grid* grid, const Point p) : grid_(grid) { set(p); }

    inline int index() const { return index_; }
    inline Point point() const { return point_; }
    inline T value() const { return (*grid_)[index_]; }
    inline bool ok() const { return grid_->contains(index_); }

    Iterator& set(const int i) {
      if (!ok()) return *this;
      if (!grid_->contains(i)) {
        oob();
        return *this;
      }
      index_ = i;
      point_ = grid_->itop(i);
      return *this;
    }

    Iterator& set(const Point p) {
      if (!ok()) return *this;
      if (!grid_->contains(p)) {
        oob();
        return *this;
      }
      point_ = p;
      index_ = grid_->ptoi(p);
      return *this;
    }

    inline Iterator& move(const int offset = 1) {
      return set(index() + offset);
    }

    inline Iterator& move(const Point offset) { return set(point() + offset); }

#define GRID_ITER_MOVE(dir, x, y) \
  inline Iterator& dir(const int n = 1) { return move({x, y}); }

    GRID_ITER_MOVE(left, -n, 0)
    GRID_ITER_MOVE(right, n, 0)
    GRID_ITER_MOVE(up, 0, n)
    GRID_ITER_MOVE(down, 0, -n)

    friend inline T operator*(const Iterator& it) { return it.value(); }

#define GRID_ITER_CMP(op)                                                    \
  friend inline bool operator op(const Iterator& lhs, const Iterator& rhs) { \
    return lhs.index() op rhs.index();                                       \
  }

    GRID_ITER_CMP(==)
    GRID_ITER_CMP(!=)
    GRID_ITER_CMP(<)
    GRID_ITER_CMP(>)
    GRID_ITER_CMP(<=)
    GRID_ITER_CMP(>=)

    friend inline Iterator& operator+=(Iterator& it, const int n) {
      return it.move(n);
    }

    friend inline Iterator& operator+=(Iterator& it, const Point p) {
      return it.move(p);
    }

    friend inline Iterator& operator++(Iterator& it) { return it.move(1); }

    friend inline Iterator operator++(Iterator& it, int) {
      auto old = it;
      it.move(1);
      return old;
    }

    friend std::ostream& operator<<(std::ostream& os, const Iterator& it) {
      return os << it.point();
    }
  };

  // Returns a random-access iterator to the first member of the grid.
  //
  // The first member is at index 0 and point {0, 0}.
  Iterator begin() const { return Iterator(this); }

  // Returns a random-access iterator to the first position beyond the end of
  // the grid.
  //
  // The end position is at index grid.size() and point {grid.width(),
  // grid.height()}.
  Iterator end() const { return Iterator(this, size()); }

  inline T operator[](const Iterator& it) const { return data_[it.index()]; }
  inline T& operator[](const Iterator& it) { return data_[it.index()]; }
};

}  // namespace aoc

#endif  // AOC_GRID_H

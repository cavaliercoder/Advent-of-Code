#ifndef AOC_INPUT_H
#define AOC_INPUT_H

#include <array>
#include <cassert>
#include <cctype>
#include <chrono>
#include <fstream>
#include <functional>
#include <iostream>
#include <iterator>
#include <sstream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

#ifdef __APPLE__
#include <CommonCrypto/CommonDigest.h>
#endif

namespace aoc {

/*
 * Math
 */

// Computes the value of base raised to the power exp in O(log(exp)).
template <typename T>
T pow(const T base, const T exp) {
  if (exp == 0) return 1;
  if (exp == 1) return base;
  T n = pow(base, exp / 2);
  if (exp % 2 == 0)
    return n * n;
  else
    return base * n * n;
}

/*
 * Testing
 */

class Error {
  std::string expr_;
  std::string op_;
  std::string expect_;
  std::string actual_;
  std::string file_;
  int line_;

 public:
  template <typename T>
  Error(const std::string expr, const std::string op, const std::string expect,
        const T& actual, const std::string file, const int line)
      : expr_(expr), op_(op), expect_(expect), file_(file), line_(line) {
    std::stringstream ss;
    ss << actual;
    actual_ = ss.str();
  }

  const std::string expr() const { return expr_; }
  const std::string op() const { return op_; }
  const std::string expect() const { return expect_; }
  const std::string actual() const { return actual_; }
  const std::string file() const { return file_; }
  const int line() const { return line_; }
};

class BaseTest {
 protected:
  std::vector<Error> errv_;

 public:
  virtual ~BaseTest() = default;
  virtual std::string name() const = 0;
  virtual std::string suite() const = 0;
  virtual std::string file() const = 0;
  virtual int line() const = 0;
  virtual void run() = 0;

  const std::vector<Error> errv() const { return errv_; }
  operator bool() const { return errv_.empty(); }

  bool operator<(const BaseTest& rhs) const {
    if (suite() < rhs.suite()) return true;
    if (suite() == rhs.suite()) return name() < rhs.name();
    return false;
  }
};

struct BaseTestFactory {
  virtual BaseTest* make() const = 0;
};

template <class T>
struct TestFactory : public BaseTestFactory {
  BaseTest* make() const override { return new T(); }
};

int register_test(const BaseTestFactory* factory);

std::vector<BaseTest*> get_tests();

int run_tests();

#define TEST_CLASS_NAME(suite_name, test_name) suite_name##test_name##Test

#define TEST(suite_name, test_name)                                           \
  class TEST_CLASS_NAME(suite_name, test_name) : public aoc::BaseTest {       \
   private:                                                                   \
    static const int id_;                                                     \
    std::string name() const override { return #test_name; }                  \
    std::string suite() const override { return #suite_name; }                \
    std::string file() const override { return __FILE__; }                    \
    int line() const override { return __LINE__; }                            \
    void run() override;                                                      \
  };                                                                          \
                                                                              \
  const int TEST_CLASS_NAME(suite_name, test_name)::id_ = aoc::register_test( \
      new aoc::TestFactory<TEST_CLASS_NAME(suite_name, test_name)>());        \
                                                                              \
  void TEST_CLASS_NAME(suite_name, test_name)::run()

#define EXPECT(expr, op, expect)                                        \
  {                                                                     \
    auto actual = expr;                                                 \
    if (!(actual op expect))                                            \
      errv_.push_back(                                                  \
          aoc::Error(#expr, #op, #expect, actual, __FILE__, __LINE__)); \
  }

#define EXPECT_EQ(a, b) EXPECT(a, ==, b)
#define EXPECT_NE(a, b) EXPECT(a, !=, b)
#define EXPECT_LT(a, b) EXPECT(a, <, b)
#define EXPECT_GT(a, b) EXPECT(a, >, b)
#define EXPECT_LE(a, b) EXPECT(a, <=, b)
#define EXPECT_GE(a, b) EXPECT(a, >=, b)
#define EXPECT_TRUE(expr) EXPECT(expr, ==, 1)
#define EXPECT_FALSE(expr) EXPECT(expr, ==, 0)

/*
 * Timing
 */

class StopWatch {
  uint64_t start_ = 0;
  uint64_t stop_ = 0;

  inline uint64_t now() {
    return std::chrono::duration_cast<std::chrono::nanoseconds>(
               std::chrono::steady_clock::now().time_since_epoch())
        .count();
  }

 public:
  void start() {
    stop_ = 0;
    start_ = now();
  }

  uint64_t stop() {
    if (stop_ != 0) return duration();  // Already stopped
    stop_ = now();
    return duration();
  }

  uint64_t duration() {
    if (!start_) return 0;
    if (!stop_) return now() - start_;
    return stop_ - start_;
  }

  friend std::ostream& operator<<(std::ostream& os, StopWatch& sw) {
    auto d = sw.duration();
    if (d < 1000) return os << d << "ns";
    if (d < 1000000) return os << d / 1000 << "µs";
    if (d < 2000000000) return os << d / 1000000 << "ms";
    return os << d / 1000000000 << "s";
  }
};

/*
 * Set container
 */

template <typename T>
class Set : public std::unordered_set<T> {
 public:
  // Set() = default;

  bool contains(const T& elem) { return this->count(elem); }

  friend Set operator+(const Set& lhs, const Set& rhs) {
    Set u;
    u += lhs;
    u += rhs;
    return u;
  }

  friend Set& operator+=(Set& lhs, const Set& rhs) {
    lhs.insert(rhs.begin(), rhs.end());
    return lhs;
  }

  friend Set& operator+=(Set& lhs, const T& rhs) {
    lhs.insert(rhs);
    return lhs;
  }

  friend Set operator-(const Set& lhs, const Set& rhs) {
    Set u;
    for (auto& elem : lhs)
      if (!rhs.count(elem)) u.insert(elem);
    return u;
  }

  friend Set& operator-=(Set& lhs, const Set& rhs) {
    lhs.erase(rhs);
    return lhs;
  }

  friend Set& operator-=(Set& lhs, const T& rhs) {
    lhs.erase(rhs);
    return lhs;
  }

  friend Set operator|(const Set& lhs, const Set& rhs) { return lhs + rhs; }
  friend Set& operator|=(Set& lhs, const Set& rhs) { return lhs += rhs; }

  friend Set operator&(const Set& lhs, const Set& rhs) {
    Set u;
    for (auto& elem : lhs)
      if (rhs.count(elem)) u.insert(elem);
    return u;
  }

  friend Set& operator&=(Set& lhs, const Set& rhs) { return lhs = lhs & rhs; }

  friend Set operator^(const Set& lhs, const Set& rhs) {
    Set u;
    for (auto& elem : lhs)
      if (!rhs.count(elem)) u.insert(elem);
    for (auto& elem : rhs)
      if (!lhs.count(elem)) u.insert(elem);
    return u;
  }

  friend Set& operator^=(Set& lhs, const Set& rhs) { lhs = lhs ^ rhs; }
};

// Min/max heap.
//
// Wraps std::*_heap functions.
// Defaults to max-heap. Use Compare=std::greater<T> for a min-heap.
template <typename T, class Compare = std::less<T>>
class Heap {
  std::vector<T> data_;
  Compare cmp_;

 public:
  Heap(std::initializer_list<T> init, Compare cmp = Compare()) : cmp_(cmp) {
    data_ = std::vector<T>(init);
    std::make_heap(data_.begin(), data_.end(), cmp_);
  }

  Heap(T init, Compare cmp = Compare()) : Heap({init}, cmp) {}

  void push(const T x) {
    data_.push_back(x);
    std::push_heap(data_.begin(), data_.end(), cmp_);
  }

  template <class IterT>
  void push(IterT first, IterT last) {
    while (first < last) push(*(first++));
  }

  T pop() {
    std::pop_heap(data_.begin(), data_.end(), cmp_);
    auto result = data_.back();
    data_.pop_back();
    return result;
  }

  bool empty() const { return data_.empty(); }
  std::size_t size() const { return data_.size(); }

  operator bool() const { return !data_.empty(); }
};

/*
 * Value Range
 */

template <typename T>
struct Range {
  T start = 0;  // First value in the range.
  T limit = 0;  // First value after the values in the range.

 public:
  inline T size() const { return limit - start; }
  inline bool empty() const { return start == limit; }

  inline T begin() const { return start; }
  inline T end() const { return limit; }

  inline bool contains(const T n) const { return n >= start && n < limit; }

  inline bool contains(const Range& r) const {
    return contains(r.start) && r.limit <= limit;
  }

  inline bool intersects(const Range& r) const {
    return contains(r.start) || contains(r.limit - 1) || r.contains(start) ||
           r.contains(limit - 1);
  }

  inline Range extend(const T n = 1) const { return {start, limit + n}; }

  inline Range inner(const Range& r) const {
    if (!intersects(r)) return {};
    return {std::max(start, r.start), std::min(limit, r.limit)};
  }

  inline Range outer(const T n) const {
    return {std::min(start, n), std::max(limit, n + 1)};
  }

  inline Range outer(const Range& r) const {
    return {std::min(start, r.start), std::max(limit, r.limit)};
  }

  inline friend bool operator==(const Range& lhs, const Range& rhs) {
    return lhs.start == rhs.start && lhs.limit == rhs.limit;
  }

  inline friend bool operator!=(const Range& lhs, const Range& rhs) {
    return lhs.start != rhs.start || lhs.limit != rhs.limit;
  }

  inline friend bool operator<(const Range& lhs, const Range& rhs) {
    if (lhs.start == rhs.start) return lhs.limit < rhs.limit;
    return lhs.start < rhs.limit;
  }

  inline friend bool operator<=(const Range& lhs, const Range& rhs) {
    return lhs < rhs || lhs == rhs;
  }

  inline friend bool operator>(const Range& lhs, const Range& rhs) {
    if (lhs.start == rhs.start) return lhs.limit > rhs.limit;
    return lhs.start > rhs.limit;
  }

  inline friend bool operator>=(const Range& lhs, const Range& rhs) {
    return lhs > rhs || lhs == rhs;
  }

  inline friend Range operator&(const Range& lhs, const Range& rhs) {
    return lhs.inner(rhs);
  }

  inline friend Range operator|(const Range& lhs, const Range& rhs) {
    return lhs.outer(rhs);
  }

  inline operator bool() { return !empty(); }

  inline friend std::ostream& operator<<(std::ostream& os, const Range& r) {
    return os << "[" << r.start << ", " << r.limit << ")";
  }

  static void flatten(std::vector<Range>& a) {
    if (a.empty()) return;
    std::vector<Range> b;
    std::sort(a.begin(), a.end());
    Range l = a[0];
    for (int i = 1; i < a.size(); ++i) {
      Range r = a[i];
      if (l.intersects(r)) {
        l = l.outer(r);
        continue;
      }
      b.push_back(l);
      l = r;
    }
    b.push_back(l);
    a = b;
  }
};

/*
 * N-Dimensional Point
 */

template <int N = 2, typename T = int>
struct Point {
  T data[N] = {};

  Point() = default;

  Point(std::initializer_list<T> init) {
    assert(init.size() == N);
    std::copy(init.begin(), init.end(), data);
  }

  Point<2, T>(const int x, const int y) : Point{x, y} {};

  Point<3, T>(const T x, const T y, const T z) : Point{x, y, z} {}

  inline T x() const { return data[0]; }
  inline T y() const { return data[1]; }
  inline T z() const { return data[2]; }

  bool empty() const {
    for (int i = 0; i < N; ++i)
      if (data[i]) return false;
    return true;
  }

  Point map(const std::function<T(T)>& f) const {
    Point p = {};
    for (int i = 0; i < N; ++i) p.data[i] = f(data[i]);
    return p;
  }

  Point map(const Point p, const std::function<T(T, T)>& f) const {
    Point q = {};
    for (int i = 0; i < N; ++i) q.data[i] = f(data[i], p.data[i]);
    return q;
  }

  inline Point abs() const {
    return map([](const T n) -> T { return std::abs(n); });
  }

  inline Point min(const Point p) const {
    return map(p, [](const T a, const T b) -> T { return std::min(a, b); });
  }

  inline Point max(const Point p) const {
    return map(p, [](const T a, const T b) -> T { return std::max(a, b); });
  }

  inline T manhattan(const Point p) const {
    T n = 0;
    for (int i = 0; i < N; ++i) n += std::abs(data[i] - p.data[i]);
    return n;
  }

  // Increment each dimension toward p.
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

  std::array<Point, 4> udlr() const { return {up(), down(), left(), right()}; }

  std::array<Point, 6> udlrfb() const {
    return {up(), down(), left(), right(), forward(), backward()};
  }

  inline Point operator-() const {
    return map([](const T n) -> T { return -n; });
  }

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

/*
 * Data Grid
 */

// Grid is a container of elements arranged in a finite 2-dimensional space.
template <typename T = char>
class Grid {
  using Point = Point<2, int>;

  std::vector<T> data_;
  int width_ = 0;
  int height_ = 0;

  inline int ptoi(const Point p) const {
    if (!contains(p)) return size();
    return p.y() * width() + p.x();
  }

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

    inline Iterator& move(const int n = 1) { return set(index() + n); }
    inline Iterator& move(const Point p) { return set(point() + p); }

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

  Iterator begin() const { return Iterator(this); }
  Iterator end() const { return Iterator(this, size()); }

  inline T operator[](const Iterator& it) const { return data_[it.index()]; }
  inline T& operator[](const Iterator& it) { return data_[it.index()]; }
};

/*
 * Input parsing
 */

class Input {
  const int year_;
  const int day_;
  std::string path_;
  std::ifstream in_;
  bool err_;

  void set_err() { err_ = true; }

 public:
  Input() = delete;
  Input(Input&&) = default;

  Input(const int year, const int day) : year_(year), day_(day) {
    std::ostringstream ss;
    // TODO: Recurse up until ./inputs is found
    ss << "../../../inputs/" << year << "/day";
    if (day < 10) {
      ss << "0" << day;
    } else {
      ss << day;
    }
    ss << ".txt";
    path_ = ss.str();
    in_.open(path_);
    assert(!in_.fail());
  }

  ~Input() { in_.close(); }

  int year() const { return year_; }
  int day() const { return day_; }
  std::string path() { return path_; }

  class Line : public std::string {
    friend std::istream& operator>>(std::istream& is, Line& line) {
      return std::getline(is, line);
    }
  };

  // Returns an iterator that reads the input line by line.
  std::istream_iterator<Line> begin() {
    return std::istream_iterator<Line>(in_);
  }

  // Returns an iterator that represents the end-of-file marker.
  std::istream_iterator<Line> end() { return std::istream_iterator<Line>(); }

  inline char peek() { return in_.peek(); }

  inline char get() { return in_.get(); }

  Input& get(char& c) {
    c = get();
    return *this;
  }

  Input& get(std::string& s, const std::streamsize n = 1) {
    s.clear();
    for (int i = 0; i < n; ++i) s.push_back(get());
    return *this;
  }

  inline Input& get_line(std::string& s) {
    std::getline(in_, s);
    return *this;
  }

  inline std::string get_line() {
    std::string s;
    get_line(s);
    return s;
  }

  inline Input& get_token(std::string& s) {
    s.clear();
    while (isspace()) ignore();
    while (!isspace()) s.push_back(get());
    return *this;
  }

  std::string get_token() {
    std::string s;
    get_token(s);
    return s;
  }

  // Extract the next character if it equals c.
  bool branch(const char c) {
    if (!is(c)) return false;
    in_.ignore();
    return true;
  }

  Input& expect(const char c, const std::streamsize n = 1) {
    for (int i = 0; i < n; i++) {
      if (get() != c) {
        set_err();
        break;
      }
    }
    return *this;
  }

  Input& expect(const char* str) {
    for (; *str; ++str) expect(*str);
    return *this;
  }

  // Input& expect(const char* str, const size_t len) {
  //   for (int i = 0; i < len; ++i) expect(str[i]);
  // }

  // Input& expect(const std::string& str) {
  //   for (const char c : str) expect(c);
  // }

  Input& discard(const char c) {
    if (is(c)) in_.ignore();
    return *this;
  }

  Input& discard_to(const char c) {
    while (in_)
      if (get() == c) break;
    return *this;
  }

  Input& ignore(std::streamsize n = 1, int delim = EOF) {
    in_.ignore(n, delim);
    return *this;
  }

  inline bool is(const char c) { return peek() == c; }
  inline bool isdigit() { return std::isdigit(peek()); }
  inline bool isspace() { return std::isspace(peek()); }

  template <typename T>
  T get_uint() {
    char c = get();
    if (!std::isdigit(c)) {
      set_err();
      return 0;
    }
    T n = c - '0';
    while (isdigit()) {
      c = get();
      n *= 10;
      n += c - '0';
    }
    return n;
  }

  template <typename T>
  Input& get_uint(T& n) {
    n = get_uint<T>();
    return *this;
  }

  template <typename T>
  T get_int() {
    T sign = 1;
    char c = peek();
    if (c == '+') {
      ignore();
    } else if (c == '-') {
      sign = -1;
      ignore();
    }
    return sign * get_uint<T>();
  }

  template <typename T>
  Input& get_int(T& n) {
    n = get_int<T>();
    return *this;
  }

  template <typename T>
  Grid<T> grid(std::function<T(const char c)> f) {
    std::vector<T> data;
    int width = 0;
    for (auto& s : *this) {
      if (!width) width = s.size();
      assert(s.size() == width);
      for (auto c : s) data.push_back(f(c));
    }
    assert(data.size() % width == 0);
    return Grid<T>(width, data.size() / width, data);
  }

  Grid<char> grid() {
    return grid<char>([](const char c) -> char { return c; });
  }

  inline char operator*() { return peek(); }
  inline operator bool() { return !err_ && in_ && in_.peek() != EOF; }
};

/*
 * Algorithms
 */

// Computes the shortest path weight between all vertices in a graph in Θ(|V|³).
//
// Edges are provided as an AdjacencyList keyed by Key, which should map source
// vertices to their connected vertices, to the weight of each connected edge.
// For example, edges[u][v] should return the weight of the edge from u to v.
//
// The inf value should be high enough that it will exceed the maximum weight of
// any path through the graph but not so high that 2*inf creates an integer
// overflow or sign change.
//
// The returned map can be used to find the shortest path between any two
// points.
//
// To find the the shortest distance between vertices v and u:
//
//     auto weights = floyd_warshall<Key>(edges);
//     auto best = weight[v][u];
//
template <typename Key, typename Weight = int,
          typename AdjacencyList =
              std::unordered_map<Key, std::unordered_map<Key, Weight>>>
std::unordered_map<Key, std::unordered_map<Key, Weight>> floyd_warshall(
    AdjacencyList& edges, const Weight inf = 1 << 30) {
  auto distance = std::unordered_map<Key, std::unordered_map<Key, Weight>>();
  for (const auto& [v, _] : edges)
    for (const auto& [u, _] : edges) distance[v][u] = inf;
  for (const auto& [v, ve] : edges) {
    distance[v][v] = 0;
    for (const auto& [u, w] : ve) distance[v][u] = w;
  }
  for (const auto& [k, _] : edges)
    for (const auto& [i, _] : edges)
      for (const auto& [j, _] : edges)
        if (distance[i][k] < inf && distance[k][j] < inf)
          distance[i][j] =
              std::min(distance[i][j], distance[i][k] + distance[k][j]);
  return distance;
}

/*
 * Hex encoding.
 */

static constexpr char hex_lower[17] = "0123456789abcdef";

inline char* hex(char* buf, const uint8_t data) {
  *(buf++) = hex_lower[data >> 4];
  *(buf++) = hex_lower[data & 0x0F];
  return buf;
}

inline char* hex(char* buf, const void* data, size_t len) {
  const uint8_t* p = static_cast<const uint8_t*>(data);
  while (len--) buf = hex(buf, *p++);
  return buf;
}

inline char* hex(char* buf, const void* data, size_t len,
                 const bool little_endian) {
  if (!little_endian) return hex(buf, data, len);
  const uint8_t* p = static_cast<const uint8_t*>(data);
  p += len - 1;
  while (len--) buf = hex(buf, *p--);
  return buf;
}

inline std::ostream& hex(std::ostream& os, const uint8_t data) {
  os.put(hex_lower[(data >> 4)]);
  os.put(hex_lower[data & 0x0F]);
  return os;
}

inline std::ostream& hex(std::ostream& os, const void* data, size_t len) {
  const char* p = static_cast<const char*>(data);
  while (len--) hex(os, *p++);
  return os;
}

inline std::ostream& hex(std::ostream& os, const void* data, size_t len,
                         const bool little_endian) {
  if (!little_endian) return hex(os, data, len);
  const uint8_t* p = static_cast<const uint8_t*>(data);
  p += len - 1;
  while (len--) hex(os, *p--);
  return os;
}

inline std::string hex(const void* data, size_t len) {
  char buf[len * 2];
  hex(&buf[0], data, len);
  return std::string(buf, len * 2);
}

template <typename T>
inline std::string hex(T value) {
  static const size_t len = sizeof(T) * 2;
  char buf[len];
  hex(&buf[0], &value, sizeof(T), true);
  return std::string(buf, len);
}

// Computes CRC32 checksums.
template <typename T = uint_fast32_t, T Polynomial = 0xEDB88320>
class CRC32 {
  T crc_ = 0;

  static std::array<T, 256> gen_table() {
    T p = Polynomial;
    std::array<T, 256> table = {};
    for (T i = 0; i < 256; ++i) {
      T c = i;
      for (size_t j = 0; j < 8; j++) {
        if (c & 1) {
          c = p ^ (c >> 1);
        } else {
          c >>= 1;
        }
      }
      table[i] = c;
    }
    return table;
  }

 public:
  CRC32() = default;
  CRC32(const T crc) : crc_(crc) {}
  CRC32(const void* buf, size_t size) { update(buf, size); }

  T sum() const { return crc_; }
  void reset() { crc_ = 0; }
  T operator*() const { return crc_; }
  operator T() const { return crc_; }

  std::string str() const { return hex(crc_); }

  friend std::ostream& operator<<(std::ostream& os, CRC32 crc) {
    return hex(os, crc.crc_);
  }

  T update(const void* buf, size_t size) {
    static const auto table = gen_table();
    crc_ ^= 0xFFFFFFFF;
    const uint8_t* p = static_cast<const uint8_t*>(buf);
    while (size--) crc_ = table[(crc_ ^ *p++) & 0xff] ^ (crc_ >> 8);
    return crc_ ^= 0xFFFFFFFF;
  }
};

#define DEFINE_DIGEST(class_name, digest_size)                       \
  class class_name {                                                 \
    CC_##class_name##_CTX ctx_;                                      \
                                                                     \
   public:                                                           \
    class Digest {                                                   \
      unsigned char data_[32];                                       \
                                                                     \
     public:                                                         \
      int size() const;                                              \
      std::string str() const;                                       \
                                                                     \
      unsigned char operator[](const int i) const;                   \
      friend bool operator==(const Digest& lhs, const Digest& rhs);  \
      friend bool operator!=(const Digest& lhs, const Digest& rhs);  \
      friend std::ostream& operator<<(std::ostream&, const Digest&); \
                                                                     \
      friend class class_name;                                       \
    };                                                               \
                                                                     \
    class_name();                                                    \
    class_name(const void* data, size_t len);                        \
    class_name(const char* str);                                     \
    class_name(const std::string& str);                              \
                                                                     \
    void update(const void* data, size_t len);                       \
    void update(const char* str);                                    \
    void update(const std::string& str);                             \
    void reset();                                                    \
    Digest sum() const;                                              \
    std::string str() const;                                         \
                                                                     \
    Digest operator()() const;                                       \
  };

DEFINE_DIGEST(SHA1, 20);
DEFINE_DIGEST(SHA256, 32);

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

#endif  // AOC_INPUT_H

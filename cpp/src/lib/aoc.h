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
#include <unordered_set>
#include <vector>

namespace aoc {

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
    if (d < 1000000000) return os << d / 1000000 << "ms";
    return os << d / 1000000000 << "s";
  }
};

/*
 * Set and maps
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
    if (lhs.start == rhs.start) return lhs.end_ < rhs.end_;
    return lhs.start < rhs.end_;
  }

  inline friend bool operator<=(const Range& lhs, const Range& rhs) {
    return lhs < rhs || lhs == rhs;
  }

  inline friend bool operator>(const Range& lhs, const Range& rhs) {
    if (lhs.start == rhs.start) return lhs.end_ > rhs.end_;
    return lhs.start > rhs.end_;
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
  T data[N];

  Point() {
    for (int i = 0; i < N; ++i) data[i] = 0;
  }

  Point<2, T>(const int x, const int y) {
    data[0] = x;
    data[1] = y;
  };

  Point<3, T>(const T x, const T y, const T z) {
    data[0] = x;
    data[1] = y;
    data[2] = z;
  };

  inline T x() const { return data[0]; }
  inline T y() const { return data[1]; }
  inline T z() const { return data[2]; }

  bool empty() const {
    for (int i = 0; i < N; ++i)
      if (data[i]) return false;
    return true;
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

  Point map(const std::function<T(T)>& f) const {
    Point p = {};
    for (int i = 0; i < N; ++i) p.data[i] = f(data[i]);
    return p;
  }

  Point map(const std::function<T(const T, const T)>& f,
            const Point rhs) const {
    Point p = {};
    for (int i = 0; i < N; ++i) p.data[i] = f(data[i], rhs.data[i]);
    return p;
  }

  inline Point abs() const {
    return map([](T n) -> T { return std::abs(n); });
  }

  inline Point min(const Point p) const {
    return map([](const T a, const T b) -> T { return std::min(a, b); });
  }

  inline Point max(const Point p) const {
    return map([](const T a, const T b) -> T { return std::max(a, b); });
  }

#define POINT_ARITHMATIC(op)                                          \
  inline friend Point operator op(const Point lhs, const Point rhs) { \
    return lhs.map([](T a, T b) -> T { return a op b; }, rhs);        \
  }                                                                   \
                                                                      \
  inline friend Point operator op(const Point lhs, const T rhs) {     \
    return lhs.map([rhs](T n) -> T { return n op rhs; });             \
  }                                                                   \
                                                                      \
  inline friend Point operator op(const T lhs, const Point rhs) {     \
    return rhs op lhs;                                                \
  }                                                                   \
                                                                      \
  inline friend Point operator op##=(Point& lhs, const Point rhs) {   \
    return lhs = lhs op rhs;                                          \
  }                                                                   \
                                                                      \
  inline friend Point operator op##=(Point& lhs, const T rhs) {       \
    return lhs = lhs op rhs;                                          \
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

/*
 * Data Grid
 */

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

#ifndef AOC_RANGE_H
#define AOC_RANGE_H

namespace aoc {

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

}  // namespace aoc

#endif  //  AOC_RANGE_H

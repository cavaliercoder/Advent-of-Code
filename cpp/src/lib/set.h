
#ifndef AOC_SET_H
#define AOC_SET_H

#include <unordered_set>

namespace aoc {

template <typename T>
class Set : public std::unordered_set<T> {
 public:
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

}  // namespace aoc

#endif  // AOC_SET_H

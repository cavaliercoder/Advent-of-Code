#ifndef AOC_HEAP_H
#define AOC_HEAP_H

#include <algorithm>
#include <functional>
#include <iostream>
#include <unordered_set>

namespace aoc {

// Set extends std::unordered_set to support set arithmatic.
template <class T, class Compare = std::less<T>, class Hash = std::hash<T>>
class Set : public std::unordered_set<T, Hash> {
 public:
  using std::unordered_set<T, Hash>::unordered_set;

  bool contains(const T& elem) { return this->count(elem); }
  bool contains(const Set& s) { return s <= *this; }

  // Returns a vector containing all elements in an indeterminate order.
  std::vector<T> vector(Compare comp = Compare()) const {
    std::vector<T> a;
    a.reserve(this->size());
    for (auto elem : *this) {
      a.push_back(elem);
    }
    return a;
  }

  // Returns a sorted vector containing all elements.
  std::vector<T> sort(Compare comp = Compare()) const {
    auto a = vector();
    std::sort(a.begin(), a.end(), comp);
    return a;
  }

  // Returns true if lhs is a subset of rhs.
  friend bool operator<(const Set& lhs, const Set& rhs) {
    if (lhs.size() >= rhs.size()) return false;
    for (auto& elem : lhs)
      if (!rhs.count(elem)) return false;
    return true;
  }

  // Returns true if lhs is equal to or a subset of rhs.
  friend bool operator<=(const Set& lhs, const Set& rhs) {
    if (lhs.size() > rhs.size()) return false;
    for (auto& elem : lhs)
      if (!rhs.count(elem)) return false;
    return true;
  }

  // Returns true if lhs is a superset of rhs.
  friend bool operator>(const Set& lhs, const Set& rhs) {
    if (lhs.size() <= rhs.size()) return false;
    for (auto& elem : rhs)
      if (!lhs.count(elem)) return false;
    return true;
  }

  // Returns true if lhs is equal to or a superset of rhs.
  friend bool operator>=(const Set& lhs, const Set& rhs) {
    if (lhs.size() < rhs.size()) return false;
    for (auto& elem : rhs)
      if (!lhs.count(elem)) return false;
    return true;
  }

  // Returns a new set which is the union of lhs and rhs.
  friend Set operator+(const Set& lhs, const Set& rhs) {
    Set u;
    u += lhs;
    u += rhs;
    return u;
  }

  // Returns a new set which contains all elements of lhs and the rhs element.
  friend Set operator+(const Set& lhs, const T& rhs) {
    Set u;
    u += lhs;
    u += rhs;
    return u;
  }

  // Adds all elements of rhs to lhs.
  friend Set& operator+=(Set& lhs, const Set& rhs) {
    lhs.insert(rhs.begin(), rhs.end());
    return lhs;
  }

  // Adds a single element to lhs.
  friend Set& operator+=(Set& lhs, const T& rhs) {
    lhs.insert(rhs);
    return lhs;
  }

  // Removes each element in rhs from lhs.
  //
  // It is an error if rhs is not <= lhs.
  friend Set operator-(const Set& lhs, const Set& rhs) {
    Set u;
    u += lhs;
    u -= rhs;
    return u;
  }

  // Returns a new set which has all elements of lhs except rhs.
  //
  // It is an error if rhs is not in lhs.
  friend Set operator-(const Set& lhs, const T& rhs) {
    Set u;
    u += lhs;
    u -= rhs;
    return u;
  }

  // Removes all elements of rhs from lhs.
  //
  // It is an error if rhs is not in lhs.
  friend Set& operator-=(Set& lhs, const Set& rhs) {
    for (auto& elem : rhs) lhs.erase(elem);
    return lhs;
  }

  // Removes a single element from lhs.
  //
  // It is an error if rhs is not in lhs.
  friend Set& operator-=(Set& lhs, const T& rhs) {
    lhs.erase(rhs);
    return lhs;
  }

  // Returns a new set which is the union of two sets.
  friend Set operator|(const Set& lhs, const Set& rhs) { return lhs + rhs; }

  // Adds all elements of rhs to lhs.
  friend Set& operator|=(Set& lhs, const Set& rhs) { return lhs += rhs; }

  // Returns a new set which is the intersect of two sets.
  friend Set operator&(const Set& lhs, const Set& rhs) {
    Set u;
    for (auto& elem : lhs)
      if (rhs.count(elem)) u.insert(elem);
    return u;
  }

  // Removes all elements from lhs which are not in rhs.
  friend Set& operator&=(Set& lhs, const Set& rhs) { return lhs = lhs & rhs; }

  // Returns a new set which is the difference between two sets.
  friend Set operator^(const Set& lhs, const Set& rhs) {
    Set u;
    for (auto& elem : lhs)
      if (!rhs.count(elem)) u.insert(elem);
    for (auto& elem : rhs)
      if (!lhs.count(elem)) u.insert(elem);
    return u;
  }

  // Replaces the contents of lhs with the difference between lhs and rhs.
  friend Set& operator^=(Set& lhs, const Set& rhs) { return lhs = lhs ^ rhs; }

  friend std::ostream& operator<<(std::ostream& os, Set s) {
    os << "{";
    bool join = false;
    for (auto elem : s.sort()) {
      if (join) os << ", ";
      join = true;
      os << elem;
    }
    return os << "}";
  }
};

// Stack is a convenient alterative to std::stack.
template <typename T>
class Stack {
  std::size_t max_size_;
  std::size_t stacked_;

 protected:
  std::vector<T> data_;

 public:
  Stack() : Stack({}) {}
  Stack(T init) : Stack({init}) {}

  Stack(std::initializer_list<T> init) {
    data_ = std::vector<T>(init);
    data_.reserve(1024);
  }

  virtual void push(const T x) {
    data_.push_back(x);
    ++stacked_;
    max_size_ = std::max(max_size_, data_.size());
  }

  virtual void clear() noexcept { data_.clear(); }

  T peek() const { return data_.back(); }

  T pop() {
    auto result = data_.back();
    data_.pop_back();
    return result;
  }

  // Returns the largest observed length of the stack.
  std::size_t max_size() const { return max_size_; }

  // Returns the total observed number of items that were pushed to the stack.
  std::size_t stacked() const { return stacked_; }

  bool empty() const noexcept { return data_.empty(); }
  std::size_t size() const noexcept { return data_.size(); }
  operator bool() const { return !data_.empty(); }
};

// UniqueStack is a Stack that ignores items it has already seen.
//
// Useful for depth-first search, etc.
template <typename T>
class UniqueStack : public Stack<T> {
  std::unordered_set<T> seen_;

 public:
  using Stack<T>::Stack;

  void push(const T x) override {
    if (seen_.count(x)) return;
    seen_.insert(x);
    Stack<T>::push(x);
  }

  // Removes all elements from the stack, and the set of observed items.
  void clear() noexcept override {
    Stack<T>::clear();
    seen_.clear();
  }
};

// Min/max heap.
//
// Convenience wrapper for std::*_heap functions.
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

  // Implements the bool operator, so the heap can be used as follows:
  //
  //    while(heap) { auto p = heap.pop(); // do something }
  //
  operator bool() const { return !data_.empty(); }
};

}  // namespace aoc

#endif  // AOC_HEAP_H

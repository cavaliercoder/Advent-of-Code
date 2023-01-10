#ifndef AOC_HEAP_H
#define AOC_HEAP_H

#include <algorithm>

namespace aoc {

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

  // Implements the bool operator, so the heap can be used as follows:
  //
  //    while(heap) { auto p = heap.pop(); // do something }
  //
  operator bool() const { return !data_.empty(); }
};

}  // namespace aoc

#endif  // AOC_HEAP_H

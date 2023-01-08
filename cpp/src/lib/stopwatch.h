#ifndef AOC_STOPWATCH_H
#define AOC_STOPWATCH_H

#include <chrono>
#include <iostream>

namespace aoc {

class Stopwatch {
  uint64_t start_ = 0;
  uint64_t stop_ = 0;

 public:
  void start();
  uint64_t stop();
  uint64_t duration() const;
  std::string str() const;
  friend std::ostream& operator<<(std::ostream& os, Stopwatch& sw);
};

}  // namespace aoc

#endif  // AOC_STOPWATCH_H

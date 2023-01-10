#ifndef AOC_STOPWATCH_H
#define AOC_STOPWATCH_H

#include <chrono>
#include <iostream>

namespace aoc {

// Stopwatch can be used to time operations.
class Stopwatch {
  uint64_t start_ = 0;
  uint64_t stop_ = 0;

 public:
  // Start or restart the stopwatch.
  void start();

  // Stops the stopwatch and returns the time elapsed in nanoseconds.
  //
  // Returns the same result on subsequent calls until the stopwatch is
  // restarted.
  uint64_t stop();

  // Returns the time elapsed between the last start and stop.
  //
  // If the stopwatch is still running, returns the time elapsed since it
  // started.
  uint64_t duration() const;

  // Returns a string representation of the time elapsed.
  std::string str() const;

  friend std::ostream& operator<<(std::ostream& os, Stopwatch& sw);
};

}  // namespace aoc

#endif  // AOC_STOPWATCH_H

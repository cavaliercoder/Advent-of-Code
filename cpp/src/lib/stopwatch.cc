#include "stopwatch.h"

#include <iostream>
#include <sstream>

namespace aoc {

static uint64_t now() {
  return std::chrono::duration_cast<std::chrono::nanoseconds>(
             std::chrono::steady_clock::now().time_since_epoch())
      .count();
}

Stopwatch& Stopwatch::start() {
  stop_ = 0;
  start_ = now();
  return *this;
}

Stopwatch& Stopwatch::stop() {
  if (stop_ != 0) return *this;  // Already stopped
  stop_ = now();
  return *this;
}

uint64_t Stopwatch::duration() const {
  if (!start_) return 0;
  if (!stop_) return now() - start_;
  return stop_ - start_;
}

uint64_t Stopwatch::operator*() const { return duration(); }

std::string Stopwatch::str() const {
  std::stringstream ss;
  auto d = duration();
  if (d < 1000) {
    ss << d << "ns";
  } else if (d < 1000000) {
    ss << d / 1000 << "us";
  } else if (d < 2000000000) {
    ss << d / 1000000 << "ms";
  } else {
    ss << d / 1000000000 << "s";
  }
  return ss.str();
}

std::ostream& operator<<(std::ostream& os, const Stopwatch& sw) {
  return os << sw.str();
}

}  // namespace aoc

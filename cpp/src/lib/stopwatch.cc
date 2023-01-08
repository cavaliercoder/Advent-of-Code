#include "stopwatch.h"

namespace aoc {

uint64_t Stopwatch::now() {
  return std::chrono::duration_cast<std::chrono::nanoseconds>(
             std::chrono::steady_clock::now().time_since_epoch())
      .count();
}

void Stopwatch::start() {
  stop_ = 0;
  start_ = now();
}

uint64_t Stopwatch::stop() {
  if (stop_ != 0) return duration();  // Already stopped
  stop_ = now();
  return duration();
}

uint64_t Stopwatch::duration() {
  if (!start_) return 0;
  if (!stop_) return now() - start_;
  return stop_ - start_;
}

std::ostream& operator<<(std::ostream& os, Stopwatch& sw) {
  auto d = sw.duration();
  if (d < 1000) return os << d << "ns";
  if (d < 1000000) return os << d / 1000 << "µs";
  if (d < 2000000000) return os << d / 1000000 << "ms";
  return os << d / 1000000000 << "s";
}

}  // namespace aoc

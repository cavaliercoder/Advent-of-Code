#ifndef AOC_MATH_H
#define AOC_MATH_H

namespace aoc {

// Computes the value of base raised to the power exp in O(log(exp)).
template <typename T>
constexpr T pow(const T base, const T exp) {
  if (exp == 0) return 1;
  if (exp == 1) return base;
  T n = pow(base, exp / 2);
  if (exp % 2 == 0)
    return n * n;
  else
    return base * n * n;
}

}  // namespace aoc

#endif  // AOC_MATH_H

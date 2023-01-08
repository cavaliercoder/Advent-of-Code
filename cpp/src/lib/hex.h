#ifndef AOC_HEX_H
#define AOC_HEX_H

#include <iostream>

namespace aoc {

char* hex(char* buf, const uint8_t data);
char* hex(char* buf, const void* data, size_t len);
char* hex(char* buf, const void* data, size_t len, const bool little_endian);
std::ostream& hex(std::ostream& os, const uint8_t data);
std::ostream& hex(std::ostream& os, const void* data, size_t len);
std::ostream& hex(std::ostream& os, const void* data, size_t len,
                  const bool little_endian);
std::string hex(const void* data, size_t len);

template <typename T>
inline std::string hex(T value) {
  static const size_t len = sizeof(T) * 2;
  char buf[len];
  hex(&buf[0], &value, sizeof(T), true);
  return std::string(buf, len);
}

}  // namespace aoc

#endif  // AOC_HEX_H

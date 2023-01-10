#ifndef AOC_HEX_H
#define AOC_HEX_H

#include <iostream>

namespace aoc {

// Encodes a single byte to buf as a hexidecimal string.
//
// The buffer must have room for two chars.
char* hex(char* buf, const uint8_t data);

// Encodes up to len bytes of data as a hexidecimal string to buf.
//
// Uses big-endian byte ordering.
// The buffer must have room for the hexidecimal string which will be 2*len.
char* hex(char* buf, const void* data, size_t len);

// Encodes up to len bytes of data as a hexidecimal string to buf.
//
// May optionally use little-endian byte ordering.
// The buffer must have room for the hexidecimal string which will be 2*len.
char* hex(char* buf, const void* data, size_t len, const bool little_endian);

// Encodes a single byte to os as a hexidecimal string.
std::ostream& hex(std::ostream& os, const uint8_t data);

// Encodes up to len bytes of data as a hexidecimal string to os.
//
// Uses big-endian byte ordering.
std::ostream& hex(std::ostream& os, const void* data, size_t len);

// Encodes up to len bytes of data as a hexidecimal string.
//
// Uses big-endian byte ordering.
std::string hex(const void* data, size_t len);

// Encodes the given value as a hexidecimal string.
//
// Uses little-endian byte ordering.
template <typename T>
inline std::string hex(T value) {
  char buf[sizeof(T) * 2];
  hex(&buf[0], &value, sizeof(T), true);
  return std::string(buf, sizeof(T) * 2);
}

}  // namespace aoc

#endif  // AOC_HEX_H

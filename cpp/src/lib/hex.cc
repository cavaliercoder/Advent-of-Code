#include "hex.h"

namespace aoc {

static constexpr char hex_lower[17] = "0123456789abcdef";

char* hex(char* buf, const uint8_t data) {
  *(buf++) = hex_lower[data >> 4];
  *(buf++) = hex_lower[data & 0x0F];
  return buf;
}

char* hex(char* buf, const void* data, size_t len) {
  const uint8_t* p = static_cast<const uint8_t*>(data);
  while (len--) buf = hex(buf, *p++);
  return buf;
}

char* hex(char* buf, const void* data, size_t len, const bool little_endian) {
  if (!little_endian) return hex(buf, data, len);
  const uint8_t* p = static_cast<const uint8_t*>(data);
  p += len - 1;
  while (len--) buf = hex(buf, *p--);
  return buf;
}

std::ostream& hex(std::ostream& os, const uint8_t data) {
  os.put(hex_lower[(data >> 4)]);
  os.put(hex_lower[data & 0x0F]);
  return os;
}

std::ostream& hex(std::ostream& os, const void* data, size_t len) {
  const char* p = static_cast<const char*>(data);
  while (len--) hex(os, *p++);
  return os;
}

std::ostream& hex(std::ostream& os, const void* data, size_t len,
                  const bool little_endian) {
  if (!little_endian) return hex(os, data, len);
  const uint8_t* p = static_cast<const uint8_t*>(data);
  p += len - 1;
  while (len--) hex(os, *p--);
  return os;
}

std::string hex(const void* data, size_t len) {
  char buf[len * 2];
  hex(&buf[0], data, len);
  return std::string(buf, len * 2);
}

}  // namespace aoc

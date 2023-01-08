#include "hex.h"

#include "testing.h"

namespace aoc {

TEST(AoC, Hex) {
  EXPECT_EQ(hex(uint8_t(0x12)), "12");
  EXPECT_EQ(hex(uint16_t(0x12)), "0012");
  EXPECT_EQ(hex(uint16_t(0x1200)), "1200");
  EXPECT_EQ(hex(uint32_t(0x12)), "00000012");
  EXPECT_EQ(hex(uint16_t(0x3456)), "3456");
  EXPECT_EQ(hex(uint32_t(0x789ABCDE)), "789abcde");
  EXPECT_EQ(hex(uint64_t(0x0123456789ABCDEF)), "0123456789abcdef");
}

}  // namespace aoc

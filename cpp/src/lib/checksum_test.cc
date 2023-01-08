#include "checksum.h"

#include "testing.h"

namespace aoc {

TEST(AoC, Checksum) {
  const char* data = "The quick brown fox jumps over the lazy dog";
  const size_t len = 43;

  using CRC32 = CRC32<>;

  // Cases: https://rosettacode.org/wiki/CRC-32#C++
  char buf[32];
  for (int i = 0; i < 32; ++i) buf[i] = 0;
  EXPECT_EQ(CRC32(&buf[0], 32), 0x190A55AD);
  for (int i = 0; i < 32; ++i) buf[i] = 0xFF;
  EXPECT_EQ(CRC32(&buf[0], 32), 0xFF6CAB0B);
  for (int i = 0; i < 32; ++i) buf[i] = i;
  EXPECT_EQ(CRC32(&buf[0], 32), 0x91267E8A);

  EXPECT_EQ(CRC32(data, len).str(), "414fa339");

  CRC32 crc32;
  for (auto c : std::string(data)) {
    crc32.update(&c, sizeof(c));
  }
  EXPECT_EQ(*crc32, 0x414FA339);

  auto sha1 = SHA1(data, len);
  EXPECT_EQ(sha1().str(), "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12");

  auto sha256 = SHA256();
  EXPECT_EQ(sha256.str(),
            "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855");
  sha256.update(data, len);
  EXPECT_EQ(sha256.str(),
            "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592");
  sha256.reset();
  EXPECT_EQ(sha256.str(),
            "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855");
}

}  // namespace aoc

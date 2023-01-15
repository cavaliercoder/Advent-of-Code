#ifndef AOC_CHECKSUM_H
#define AOC_CHECKSUM_H

#include <array>
#include <cstring>

#ifdef __APPLE__

#include <CommonCrypto/CommonDigest.h>

#define DIGEST_CTX(digest_name) CC_##digest_name##_CTX
#define DIGEST_INIT(digest_name) CC_##digest_name##_Init
#define DIGEST_UPDATE(digest_name) CC_##digest_name##_Update
#define DIGEST_FINAL(digest_name) CC_##digest_name##_Final

#else

#include <openssl/sha.h>

typedef SHA_CTX SHA1_CTX;

#define DIGEST_CTX(digest_name) digest_name##_CTX
#define DIGEST_INIT(digest_name) digest_name##_Init
#define DIGEST_UPDATE(digest_name) digest_name##_Update
#define DIGEST_FINAL(digest_name) digest_name##_Final

#endif

#include "hex.h"

namespace aoc {

// Computes CRC32 checksums.
template <typename T = uint32_t, T Polynomial = 0xEDB88320>
class CRC32 {
  T crc_ = 0;

  static std::array<T, 256> gen_table() {
    T p = Polynomial;
    std::array<T, 256> table = {};
    for (T i = 0; i < 256; ++i) {
      T c = i;
      for (size_t j = 0; j < 8; j++) {
        if (c & 1) {
          c = p ^ (c >> 1);
        } else {
          c >>= 1;
        }
      }
      table[i] = c;
    }
    return table;
  }

 public:
  CRC32() = default;
  CRC32(const T crc) : crc_(crc) {}
  CRC32(const void* buf, size_t size) { update(buf, size); }

  T sum() const { return crc_; }
  void reset() { crc_ = 0; }
  T operator*() const { return crc_; }
  operator T() const { return crc_; }

  std::string str() const { return hex(crc_); }

  friend std::ostream& operator<<(std::ostream& os, CRC32 crc) {
    return hex(os, crc.crc_);
  }

  T update(const void* buf, size_t size) {
    static const auto table = gen_table();
    crc_ ^= 0xFFFFFFFF;
    const uint8_t* p = static_cast<const uint8_t*>(buf);
    while (size--) crc_ = table[(crc_ ^ *p++) & 0xff] ^ (crc_ >> 8);
    return crc_ ^= 0xFFFFFFFF;
  }
};

#define DIGEST_DECL(digest_name, digest_size)                        \
  class digest_name {                                                \
    DIGEST_CTX(digest_name) ctx_;                                    \
                                                                     \
   public:                                                           \
    class Digest {                                                   \
      unsigned char data_[32];                                       \
                                                                     \
     public:                                                         \
      int size() const;                                              \
      std::string str() const;                                       \
                                                                     \
      unsigned char operator[](const int i) const;                   \
      friend bool operator==(const Digest& lhs, const Digest& rhs);  \
      friend bool operator!=(const Digest& lhs, const Digest& rhs);  \
      friend std::ostream& operator<<(std::ostream&, const Digest&); \
                                                                     \
      friend class digest_name;                                      \
    };                                                               \
                                                                     \
    digest_name();                                                   \
    digest_name(const void* data, size_t len);                       \
    digest_name(const char* str);                                    \
    digest_name(const std::string& str);                             \
                                                                     \
    void update(const void* data, size_t len);                       \
    void update(const char* str);                                    \
    void update(const std::string& str);                             \
    void reset();                                                    \
    Digest sum() const;                                              \
    std::string str() const;                                         \
                                                                     \
    Digest operator()() const;                                       \
  };

DIGEST_DECL(SHA1, 20);
DIGEST_DECL(SHA256, 32);

}  // namespace aoc

#endif  // AOC_CHECKSUM_H

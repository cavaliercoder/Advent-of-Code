#include "checksum.h"

namespace aoc {

#define DIGEST_DEF(digest_name, digest_size)                                   \
  int digest_name::Digest::size() const { return digest_size; }                \
                                                                               \
  std::string digest_name::Digest::str() const {                               \
    return hex(&data_[0], digest_size);                                        \
  }                                                                            \
                                                                               \
  unsigned char digest_name::Digest::operator[](const int i) const {           \
    return data_[i];                                                           \
  }                                                                            \
                                                                               \
  bool operator==(const digest_name::Digest& lhs,                              \
                  const digest_name::Digest& rhs) {                            \
    for (int i = 0; i < digest_size; ++i)                                      \
      if (lhs[i] != rhs[i]) return false;                                      \
    return true;                                                               \
  }                                                                            \
                                                                               \
  bool operator!=(const digest_name::Digest& lhs,                              \
                  const digest_name::Digest& rhs) {                            \
    return !(lhs == rhs);                                                      \
  }                                                                            \
                                                                               \
  std::ostream& operator<<(std::ostream& os, const digest_name::Digest& sum) { \
    hex(os, &sum.data_[0], digest_size);                                       \
    return os;                                                                 \
  }                                                                            \
                                                                               \
  digest_name::digest_name() { reset(); }                                      \
                                                                               \
  digest_name::digest_name(const void* data, size_t len) : digest_name() {     \
    update(data, len);                                                         \
  }                                                                            \
                                                                               \
  digest_name::digest_name(const char* str) : digest_name() { update(str); }   \
                                                                               \
  digest_name::digest_name(const std::string& str) : digest_name() {           \
    update(str);                                                               \
  }                                                                            \
                                                                               \
  void digest_name::update(const void* data, size_t len) {                     \
    DIGEST_UPDATE(digest_name)(&ctx_, data, len);                              \
  }                                                                            \
                                                                               \
  void digest_name::update(const char* str) {                                  \
    DIGEST_UPDATE(digest_name)(&ctx_, str, std::strlen(str));                  \
  }                                                                            \
                                                                               \
  void digest_name::update(const std::string& str) {                           \
    DIGEST_UPDATE(digest_name)(&ctx_, str.c_str(), str.size());                \
  }                                                                            \
                                                                               \
  void digest_name::reset() { DIGEST_INIT(digest_name)(&ctx_); }               \
                                                                               \
  digest_name::Digest digest_name::sum() const {                               \
    DIGEST_CTX(digest_name) ctx = ctx_;                                        \
    Digest d;                                                                  \
    DIGEST_FINAL(digest_name)(&d.data_[0], &ctx);                              \
    return d;                                                                  \
  }                                                                            \
                                                                               \
  std::string digest_name::str() const { return sum().str(); }                 \
                                                                               \
  digest_name::Digest digest_name::operator()() const { return sum(); }

DIGEST_DEF(SHA1, 20);
DIGEST_DEF(SHA256, 32);

}  // namespace aoc

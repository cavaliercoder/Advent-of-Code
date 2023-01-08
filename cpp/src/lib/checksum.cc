#include "checksum.h"

namespace aoc {

#define IMPLEMENT_DIGEST(class_name, digest_size)                             \
  int class_name::Digest::size() const { return digest_size; }                \
                                                                              \
  std::string class_name::Digest::str() const {                               \
    return hex(&data_[0], digest_size);                                       \
  }                                                                           \
                                                                              \
  unsigned char class_name::Digest::operator[](const int i) const {           \
    return data_[i];                                                          \
  }                                                                           \
                                                                              \
  bool operator==(const class_name::Digest& lhs,                              \
                  const class_name::Digest& rhs) {                            \
    for (int i = 0; i < digest_size; ++i)                                     \
      if (lhs[i] != rhs[i]) return false;                                     \
    return true;                                                              \
  }                                                                           \
                                                                              \
  bool operator!=(const class_name::Digest& lhs,                              \
                  const class_name::Digest& rhs) {                            \
    return !(lhs == rhs);                                                     \
  }                                                                           \
                                                                              \
  std::ostream& operator<<(std::ostream& os, const class_name::Digest& sum) { \
    hex(os, &sum.data_[0], digest_size);                                      \
    return os;                                                                \
  }                                                                           \
                                                                              \
  class_name::class_name() { reset(); }                                       \
                                                                              \
  class_name::class_name(const void* data, size_t len) : class_name() {       \
    update(data, len);                                                        \
  }                                                                           \
                                                                              \
  class_name::class_name(const char* str) : class_name() { update(str); }     \
                                                                              \
  class_name::class_name(const std::string& str) : class_name() {             \
    update(str);                                                              \
  }                                                                           \
                                                                              \
  void class_name::update(const void* data, size_t len) {                     \
    CC_##class_name##_Update(&ctx_, data, len);                               \
  }                                                                           \
                                                                              \
  void class_name::update(const char* str) {                                  \
    CC_##class_name##_Update(&ctx_, str, std::strlen(str));                   \
  }                                                                           \
                                                                              \
  void class_name::update(const std::string& str) {                           \
    CC_##class_name##_Update(&ctx_, str.c_str(), str.size());                 \
  }                                                                           \
                                                                              \
  void class_name::reset() { CC_##class_name##_Init(&ctx_); }                 \
                                                                              \
  class_name::Digest class_name::sum() const {                                \
    CC_##class_name##_CTX ctx = ctx_;                                         \
    Digest d;                                                                 \
    CC_##class_name##_Final(&d.data_[0], &ctx);                               \
    return d;                                                                 \
  }                                                                           \
                                                                              \
  std::string class_name::str() const { return sum().str(); }                 \
                                                                              \
  class_name::Digest class_name::operator()() const { return sum(); }

IMPLEMENT_DIGEST(SHA1, 20);
IMPLEMENT_DIGEST(SHA256, 32);

}  // namespace aoc

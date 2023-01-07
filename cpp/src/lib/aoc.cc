#include "aoc.h"

#include <algorithm>
#include <iostream>
#include <vector>

namespace aoc {

std::vector<const BaseTestFactory*> test_factories;

int register_test(const BaseTestFactory* factory) {
  test_factories.push_back(factory);
  return test_factories.size();
}

std::vector<BaseTest*> get_tests() {
  std::vector<BaseTest*> tests;
  for (auto& factory : test_factories) {
    tests.push_back(factory->make());
  }
  std::sort(
      tests.begin(), tests.end(),
      [](const BaseTest* a, const BaseTest* b) -> bool { return *a < *b; });
  return tests;
}

int run_tests() {
  int errc = 0;
  auto tests = get_tests();
  StopWatch sw, sw_all;
  sw_all.start();
  for (int i = 0; i < tests.size(); ++i) {
    auto test = tests[i];
    std::cout << test->suite() << "::" << test->name();
    sw.start();
    test->run();
    sw.stop();
    std::cout << " " << sw;
    if (*test) {
      std::cout << " [\u001b[32mPASS\u001b[0m]\n";
    } else {
      ++errc;
      std::cout << " [\u001b[31mFAIL\u001b[0m]\n";
      for (auto& err : test->errv()) {
        std::cout << "  " << err.file() << ":" << err.line()
                  << " Expression evaluates to false: " << err.expr() << " "
                  << err.op() << " " << err.expect() << "\n"
                  << "    " << err.expr() << " → " << err.actual() << "\n";
      }
      std::cout << "\n";
    }
  }
  sw_all.stop();
  std::cout << "\nPassed " << (tests.size() - errc) << "/" << tests.size()
            << " tests in " << sw_all << ".\n";
  for (auto& test : tests) delete test;
  return errc;
}

/*
 * Hash digests
 */

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
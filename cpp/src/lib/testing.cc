#include "testing.h"

#include <algorithm>
#include <iomanip>
#include <iostream>

#include "stopwatch.h"

namespace aoc {

/*
 * TestError implementation.
 */

TestError::TestError(const std::string msg, const std::string file,
                     const int line)
    : msg_(msg), file_(file), line_(line) {}

const std::string TestError::msg() const { return msg_; }
const std::string TestError::file() const { return file_; }
int TestError::line() const { return line_; }
char* TestError::what() { return msg_.data(); }

/*
 * BaseTest implementation.
 */

const std::vector<TestError> BaseTest::errv() const { return errv_; }
BaseTest::operator bool() const { return errv_.empty(); }

bool BaseTest::operator<(const BaseTest& rhs) const {
  if (suite() < rhs.suite()) return true;
  if (suite() == rhs.suite()) return name() < rhs.name();
  return false;
}

std::string BaseTest::str() const {
  std::stringstream ss;
  ss << suite() << "::" << name();
  return ss.str();
}

std::ostream& operator<<(std::ostream& os, BaseTest& t) {
  return os << t.str();
}

void BaseTest::signal(void (*handler)(int)) {
  static int signals[] = {SIGABRT, SIGFPE, SIGSEGV};
  for (int i = 0; i < 3; ++i) std::signal(signals[i], handler);
}

/*
 * Test runner implementation.
 */

// Returns a global test factory list.
//
// Wrapped in a static function to avoid creating a Static Initialization Order
// Fiasco.
//
// Ref: https://en.cppreference.com/w/cpp/language/siof.
static std::vector<const BaseTestFactory*>& factories() {
  static std::vector<const BaseTestFactory*> a = {};
  return a;
}

int TestRunner::register_ctor(const BaseTestFactory* factory) {
  factories().push_back(factory);
  return factories().size();
}

std::vector<BaseTest*> TestRunner::make() {
  std::vector<BaseTest*> tests;
  for (auto& factory : factories()) tests.push_back(factory->make());
  std::sort(
      tests.begin(), tests.end(),
      [](const BaseTest* a, const BaseTest* b) -> bool { return *a < *b; });
  return tests;
}

int TestRunner::run() {
  int errc = 0;
  auto tests = make();
  int col_width = 20;
  for (auto& t : tests) col_width = std::max(col_width, int(t->str().size()));
  Stopwatch sw, sw_all;
  sw_all.start();
  for (int i = 0; i < tests.size(); ++i) {
    auto test = tests[i];
    std::cout << std::left << std::setw(col_width + 2) << test->str();
    sw.start();
    test->run();
    sw.stop();
    std::cout << std::right << std::setw(7) << sw;
    std::cout << std::right << std::setw(18);
    if (*test) {
      std::cout << "[\u001b[32mPASS\u001b[0m]\n";
    } else {
      ++errc;
      std::cout << "[\u001b[31mFAIL\u001b[0m]\n";
      for (auto& err : test->errv()) {
        std::cout << "  ";
        if (!err.file().empty()) {
          std::cout << err.file();
          if (err.line()) std::cout << ":" << err.line();
          std::cout << " ";
        }
        std::cout << err.msg() << "\n";
      }
      std::cout << std::endl;
    }
  }
  sw_all.stop();
  std::cout << "\nPassed " << (tests.size() - errc) << "/" << tests.size()
            << " tests in " << sw_all << ".\n";
  for (auto& test : tests) delete test;
  return errc;
}

}  // namespace aoc

/*
 * Optional main entry-point.
 *
 * Compile with -DTEST_MAIN to enable.
 */

#ifdef TEST_MAIN
int main(int argc, char* argv[]) { return aoc::TestRunner::run(); }
#endif

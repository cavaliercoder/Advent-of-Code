#include "testing.h"

#include <algorithm>
#include <iomanip>
#include <iostream>

#include "stopwatch.h"

namespace aoc {

/*
 * BaseTest implementation.
 */

const std::vector<Error> BaseTest::errv() const { return errv_; }
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

/*
 * Test runner.
 */

// Wraps the global test factory list in a static function to avoid the
// Static Initialization Order Fiasco.
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
  for (auto& factory : factories()) {
    tests.push_back(factory->make());
  }
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
        std::cout << "  " << err.file << ":" << err.line
                  << " Expression evaluates to false: " << err.expr << " "
                  << err.op << " " << err.expect << "\n"
                  << "    " << err.expr << " → " << err.actual << "\n";
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

#ifdef INCLUDE_TEST_RUNNER
int main(int argc, char* argv[]) { return aoc::TestRunner::run(); }
#endif

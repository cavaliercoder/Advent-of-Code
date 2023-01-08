#include "testing.h"

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
  Stopwatch sw, sw_all;
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
        std::cout << "  " << err.file << ":" << err.line
                  << " Expression evaluates to false: " << err.expr << " "
                  << err.op << " " << err.expect << "\n"
                  << "    " << err.expr << " → " << err.actual << "\n";
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

}  // namespace aoc

#ifdef INCLUDE_TEST_RUNNER
int main(int argc, char* argv[]) { return aoc::TestRunner::run(); }
#endif

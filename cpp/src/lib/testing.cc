#include "testing.h"

#include <algorithm>
#include <iomanip>
#include <iostream>
#include <regex>

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
 * TestCase implementation.
 */

const std::vector<TestError>& TestCase::errv() const { return errv_; }
const Stopwatch& TestCase::stopwatch() const { return sw_; }
const std::string TestCase::cout() const { return cout_.str(); }
TestCase::operator bool() const { return errv_.empty(); }

bool TestCase::operator<(const TestCase& rhs) const {
  if (suite() < rhs.suite()) return true;
  if (suite() == rhs.suite()) return name() < rhs.name();
  return false;
}

std::string TestCase::str() const {
  std::stringstream ss;
  ss << suite() << "::" << name();
  return ss.str();
}

std::ostream& operator<<(std::ostream& os, TestCase& t) {
  return os << t.str();
}

/*
 * Test runner implementation.
 */

std::vector<const TestFactory*>& TestRunner::factories() {
  static std::vector<const TestFactory*> a = {};
  return a;
}

bool TestRunner::filter(const TestCase* test, const std::string& filter) {
  if (filter.empty()) return false;
  std::stringstream ss;
  if (filter.front() != '^') ss << "^.*";
  ss << filter;
  if (filter.back() != '$') ss << ".*$";
  auto test_name = test->str();
  return !std::regex_match(test_name.begin(), test_name.end(),
                           std::regex(ss.str()));
}

int TestRunner::register_ctor(const TestFactory* factory) {
  factories().push_back(factory);
  return factories().size();
}

std::vector<TestCase*> TestRunner::make() const {
  std::vector<TestCase*> tests;
  for (auto& factory : factories()) tests.push_back(factory->make());
  std::sort(
      tests.begin(), tests.end(),
      [](const TestCase* a, const TestCase* b) -> bool { return *a < *b; });
  return tests;
}

static inline std::ostream& wrap_lines(std::ostream& os, const std::string& s,
                                       const std::string& prefix,
                                       const std::string& suffix,
                                       const bool skip_prefix = false) {
  if (s.empty()) return os << '\n';
  if (!skip_prefix) os << prefix;
  bool join = false;
  for (auto it = s.begin(); it < s.end(); ++it) {
    if (join) {
      os << suffix << '\n' << prefix;
      join = false;
    }
    if (*it == '\n') {
      join = true;
      continue;
    };
    os.put(*it);
  }
  os << suffix << "\n";
  return os;
}

int TestRunner::run(std::string filter) const {
  auto tests = make();
  int run_count = 0;
  int pass_count = 0;
  int fail_count = 0;
  int skip_count = 0;
  int col_width = 20;
  for (auto& t : tests) col_width = std::max(col_width, int(t->str().size()));
  auto sw = Stopwatch().start();
  for (int i = 0; i < tests.size(); ++i) {
    auto test = tests[i];
    if (this->filter(test, filter)) {
      ++skip_count;
      continue;
    }
    ++run_count;
    std::cout << std::left << std::setw(col_width + 2) << test->str();
    test->run();
    std::cout << std::right << std::setw(7) << test->stopwatch();
    std::cout << std::right << std::setw(18);
    if (*test) {
      std::cout << "[\u001b[32mPASS\u001b[0m]\n";
      ++pass_count;
      continue;
    }
    ++fail_count;
    std::cout << "[\u001b[31mFAIL\u001b[0m]\n";
    auto cout = test->cout();
    wrap_lines(std::cout, test->cout(), "  \u001b[90m", "\u001b[0m");
    for (auto& err : test->errv()) {
      std::cout << "  ";
      if (!err.file().empty()) {
        std::cout << err.file();
        if (err.line() > 0) std::cout << ":" << err.line();
        std::cout << " ";
      }
      wrap_lines(std::cout, err.msg(), "  ", "", true) << "\n";
    }
  }
  std::cout << "\nPassed " << pass_count << "/" << run_count;
  if (fail_count) std::cout << ", failed " << fail_count;
  if (skip_count) std::cout << ", skipped " << skip_count;
  std::cout << " tests in " << sw.stop() << ".\n";
  for (auto& test : tests) delete test;
  return fail_count;
}

}  // namespace aoc

/*
 * Optional main entry-point.
 *
 * Compile with -DTEST_MAIN to enable.
 */

#ifdef TEST_MAIN
void usage(std::string& cmd, std::ostream& os) {
  os << "Usage: " << cmd << " [TEST_PATTERN]\n";
}

int main(int argc, char* argv[]) {
  std::string cmd = argv[0];
  std::vector<std::string> args;
  for (int i = 1; i < argc; ++i) {
    std::string arg = argv[i];
    if (arg == "--help" || arg == "-h") {
      usage(cmd, std::cout);
      return 0;
    }
    if (arg.front() == '-') {
      usage(cmd, std::cerr);
      return 1;
    }
    args.push_back(arg);
  }
  if (args.size() > 1) {
    usage(cmd, std::cerr);
    return 1;
  }
  std::string filter;
  if (args.size() > 0) filter = args[0];
  return aoc::TestRunner().run(filter);
}
#endif

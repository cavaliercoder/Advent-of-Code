#ifndef AOC_TESTING_H
#define AOC_TESTING_H

#include <csignal>
#include <regex>
#include <sstream>
#include <string>
#include <vector>

#include "stopwatch.h"

namespace aoc {

// TestError represents a test failure.
class TestError : std::exception {
  std::string msg_;
  std::string file_;
  int line_;

 public:
  TestError(const std::string msg, const std::string file = "",
            const int line = 0);

  const std::string msg() const;
  const std::string file() const;
  int line() const;

  // Implements std::exception. Don't use this.
  char* what();
};

// Converts any thrown exception to a TestError.
// Don't use this.
#define WRAP_(expr)                                                           \
  {                                                                           \
    try {                                                                     \
      expr;                                                                   \
    } catch (aoc::TestError e) {                                              \
      errv_.push_back(e);                                                     \
    } catch (std::exception e) {                                              \
      errv_.push_back(                                                        \
          aoc::TestError("Uncaught exception: " + std::string(e.what()),      \
                         __FILE__, __LINE__));                                \
    } catch (char const* e) {                                                 \
      errv_.push_back(aoc::TestError("Uncaught exception: " + std::string(e), \
                                     __FILE__, __LINE__));                    \
    } catch (...) {                                                           \
      errv_.push_back(                                                        \
          aoc::TestError("Uncaught exception.", __FILE__, __LINE__));         \
    }                                                                         \
  }

// Base class for all user-defined test cases.
class TestCase {
 protected:
  Stopwatch sw_ = {};
  std::vector<TestError> errv_;

  std::stringstream cout_;

 public:
  virtual ~TestCase() = default;

  // Returns the name of the test case.
  virtual std::string name() const = 0;

  // Returns the name of the test suite.
  virtual std::string suite() const = 0;

  // Returns the name of the file where the test is defined.
  virtual std::string file() const = 0;

  // Returns the line in the file where the test is defined.
  virtual int line() const = 0;

  // Runs the test and converts any exceptions to TestErrors.
  virtual void run(const bool capture_stdout = true,
                   const bool capture_stderr = true) = 0;

  // // Runs the user-defined implementation of the test body.
  // virtual void run() = 0;

  // Returns a list of any errors generated by the test.
  const std::vector<TestError>& errv() const;

  // Returns a reference to the stopwatch result of the last run.
  const Stopwatch& stopwatch() const;

  const std::string cout() const;

  // Returns the tests fully qualified name.
  std::string str() const;

  // Returns true if the test passed its last run.
  operator bool() const;

  // Returns true if the test's fully qualified name is less than the other's.
  bool operator<(const TestCase& rhs) const;

  friend std::ostream& operator<<(std::ostream& os, TestCase& t);
};

// Factory base class for all user-defined test cases.
struct TestFactory {
  virtual TestCase* make() const = 0;
};

// Manages and runs test suites and cases.
class TestRunner {
 protected:
  // Returns the global test factory list.
  //
  // Wrapped in a static function to avoid a SIOF.
  static std::vector<const TestFactory*>& factories();

  static bool filter(const TestCase* test, const std::string& filter);

 public:
  // Registers the constructor for a new test case.
  //
  // Internal implementation detail.
  static int register_ctor(const TestFactory* factory);

  // Constructs and returns a new instance of each user-defined test case.
  //
  // The caller is responsible for deleting each test case.
  std::vector<TestCase*> make() const;

  // Runs all test cases and returns the number of failed cases.
  virtual int run(std::string filter = "") const;
};

// Returns a classname for the given suite and test case.
#define TEST_CLASS_NAME(suite_name, test_name) TestCase_##suite_name##test_name

// Generates a TestCase and registers it with the test runner.
#define TEST(suite_name, test_name)                                           \
  class TEST_CLASS_NAME(suite_name, test_name) : public aoc::TestCase {       \
    static const int id_;                                                     \
                                                                              \
    class Factory : public aoc::TestFactory {                                 \
      TEST_CLASS_NAME(suite_name, test_name) * make() const override {        \
        return new TEST_CLASS_NAME(suite_name, test_name)();                  \
      }                                                                       \
    };                                                                        \
                                                                              \
    std::string name() const override { return #test_name; }                  \
    std::string suite() const override { return #suite_name; }                \
    std::string file() const override { return __FILE__; }                    \
    int line() const override { return __LINE__; }                            \
                                                                              \
    static void signal_handler(const int sig) {                               \
      auto msg = std::string("Caught signal: ") + strsignal(sig);             \
      throw aoc::TestError(msg, __FILE__, __LINE__);                          \
    }                                                                         \
                                                                              \
    void test_body();                                                         \
                                                                              \
    void run(const bool capture_stdout, const bool capture_stderr) override { \
      auto cout = std::cout.rdbuf();                                          \
      auto cerr = std::cerr.rdbuf();                                          \
      if (capture_stdout) std::cout.rdbuf(cout_.rdbuf());                     \
      if (capture_stderr) std::cerr.rdbuf(cout_.rdbuf());                     \
      static int signals[] = {SIGABRT, SIGFPE, SIGSEGV};                      \
      for (int i = 0; i < 3; ++i)                                             \
        std::signal(signals[i],                                               \
                    TEST_CLASS_NAME(suite_name, test_name)::signal_handler);  \
      sw_.start();                                                            \
      WRAP_(test_body());                                                     \
      sw_.stop();                                                             \
      for (int i = 0; i < 3; ++i) std::signal(signals[i], SIG_DFL);           \
      if (capture_stdout) std::cout.rdbuf(cout);                              \
      if (capture_stderr) std::cerr.rdbuf(cerr);                              \
    }                                                                         \
  };                                                                          \
                                                                              \
  const int TEST_CLASS_NAME(suite_name, test_name)::id_ =                     \
      aoc::TestRunner::register_ctor(new Factory());                          \
                                                                              \
  void TEST_CLASS_NAME(suite_name, test_name)::test_body()

#define EXPECT_TRUE(expr)                                                     \
  {                                                                           \
    WRAP_(if (!(expr)) {                                                      \
      throw aoc::TestError("Expression evaluates to false: " #expr, __FILE__, \
                           __LINE__);                                         \
    })                                                                        \
  }

#define EXPECT_FALSE(expr)                                                   \
  {                                                                          \
    WRAP_(if (expr) {                                                        \
      throw aoc::TestError("Expression evaluates to true: " #expr, __FILE__, \
                           __LINE__);                                        \
    })                                                                       \
  }

#define EXPECT_OP(lhs, op, rhs)                                               \
  {                                                                           \
    WRAP_(auto lhsv = lhs; auto rhsv = rhs; if (!(lhsv op rhsv)) {            \
      std::stringstream ss;                                                   \
      ss << "Expression evaluates to false: " << lhsv << " " #op " " << rhsv; \
      throw aoc::TestError(ss.str(), __FILE__, __LINE__);                     \
    })                                                                        \
  }

#define EXPECT_EQ(a, b) EXPECT_OP(a, ==, b)
#define EXPECT_NE(a, b) EXPECT_OP(a, !=, b)
#define EXPECT_LT(a, b) EXPECT_OP(a, <, b)
#define EXPECT_GT(a, b) EXPECT_OP(a, >, b)
#define EXPECT_LE(a, b) EXPECT_OP(a, <=, b)
#define EXPECT_GE(a, b) EXPECT_OP(a, >=, b)

// TODO: assert

}  // namespace aoc

#endif  // AOC_TESTING_H
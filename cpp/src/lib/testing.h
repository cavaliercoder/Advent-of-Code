#ifndef AOC_TESTING_H
#define AOC_TESTING_H

#include <sstream>
#include <string>
#include <vector>

namespace aoc {

struct Error {
  std::string expr;
  std::string op;
  std::string expect;
  std::string actual;
  std::string file;
  int line;

 public:
  template <typename T>
  Error(const std::string expr, const std::string op, const std::string expect,
        const T& actual, const std::string file, const int line)
      : expr(expr), op(op), expect(expect), file(file), line(line) {
    std::stringstream ss;
    ss << actual;
    Error::actual = ss.str();
  }
};
;

class BaseTest {
 protected:
  std::vector<Error> errv_;

 public:
  virtual ~BaseTest() = default;
  virtual std::string name() const = 0;
  virtual std::string suite() const = 0;
  virtual std::string file() const = 0;
  virtual int line() const = 0;
  virtual void run() = 0;

  std::string str() const;
  const std::vector<Error> errv() const;
  operator bool() const;
  bool operator<(const BaseTest& rhs) const;
  friend std::ostream& operator<<(std::ostream& os, BaseTest& t);
};

struct BaseTestFactory {
  virtual BaseTest* make() const = 0;
};

template <class T>
struct TestFactory : public BaseTestFactory {
  BaseTest* make() const override { return new T(); }
};

struct TestRunner {
  static int register_ctor(const BaseTestFactory* factory);
  static std::vector<BaseTest*> make();
  static int run();
};

#define TEST_CLASS_NAME(suite_name, test_name) suite_name##test_name##Test

#define TEST(suite_name, test_name)                                        \
  class TEST_CLASS_NAME(suite_name, test_name) : public aoc::BaseTest {    \
   private:                                                                \
    static const int id_;                                                  \
    std::string name() const override { return #test_name; }               \
    std::string suite() const override { return #suite_name; }             \
    std::string file() const override { return __FILE__; }                 \
    int line() const override { return __LINE__; }                         \
    void run() override;                                                   \
  };                                                                       \
                                                                           \
  const int TEST_CLASS_NAME(suite_name, test_name)::id_ =                  \
      aoc::TestRunner::register_ctor(                                      \
          new aoc::TestFactory<TEST_CLASS_NAME(suite_name, test_name)>()); \
                                                                           \
  void TEST_CLASS_NAME(suite_name, test_name)::run()

#define EXPECT(expr, op, expect)                                        \
  {                                                                     \
    auto actual = expr;                                                 \
    if (!(actual op expect))                                            \
      errv_.push_back(                                                  \
          aoc::Error(#expr, #op, #expect, actual, __FILE__, __LINE__)); \
  }

#define EXPECT_TRUE(expr) EXPECT(expr, ==, 1)
#define EXPECT_FALSE(expr) EXPECT(expr, ==, 0)
#define EXPECT_EQ(a, b) EXPECT(a, ==, b)
#define EXPECT_NE(a, b) EXPECT(a, !=, b)
#define EXPECT_LT(a, b) EXPECT(a, <, b)
#define EXPECT_GT(a, b) EXPECT(a, >, b)
#define EXPECT_LE(a, b) EXPECT(a, <=, b)
#define EXPECT_GE(a, b) EXPECT(a, >=, b)

}  // namespace aoc

#endif  // AOC_TESTING_H
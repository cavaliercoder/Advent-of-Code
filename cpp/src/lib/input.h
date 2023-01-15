#ifndef AOC_INPUT_H
#define AOC_INPUT_H

#include <fstream>
#include <iterator>
#include <sstream>
#include <string>

#include "grid.h"

namespace aoc {

class Input {
  const int year_;
  const int day_;
  std::string path_;
  std::ifstream in_;
  bool err_ = 0;

  void set_err();

 public:
  Input() = delete;
  Input(Input&&) = default;
  Input(const int year, const int day, const std::string& suffix = "");
  ~Input();

  int year() const;
  int day() const;
  std::string path();

  class Line : public std::string {
    friend std::istream& operator>>(std::istream& is, Line& line);
  };

  // Returns an iterator that reads the input line by line.
  std::istream_iterator<Line> begin();

  // Returns an iterator that represents the end-of-file marker.
  std::istream_iterator<Line> end() const;

  char peek();
  char get();
  Input& get(char& c);
  Input& get(std::string& s, const std::streamsize n = 1);
  Input& get_line(std::string& s);
  std::string get_line();
  Input& get_token(std::string& s);
  std::string get_token();

  // Extract the next character and returns true if it is equal to c.
  bool branch(const char c);

  // Extracts the next n characters and sets the error flag if any of them are
  // not equal to c.
  Input& expect(const char c, const std::streamsize n = 1);

  // Extracts the next strlen(str) characters and sets the error flag if they
  // are not equal to str.
  Input& expect(const char* str);

  // Extracts the next character if it is equal to c.
  Input& discard(const char c);

  // Extracts the next characters until c is found or EOF.
  Input& discard_to(const char c);

  Input& ignore(std::streamsize n = 1, int delim = EOF);

  // Returns true if the peek() == c.
  bool is(const char c);
  bool isdigit();
  bool isspace();

  template <typename T>
  T get_uint() {
    char c = get();
    if (!std::isdigit(c)) {
      set_err();
      return 0;
    }
    T n = c - '0';
    while (isdigit()) {
      c = get();
      n *= 10;
      n += c - '0';
    }
    return n;
  }

  template <typename T>
  Input& get_uint(T& n) {
    n = get_uint<T>();
    return *this;
  }

  template <typename T>
  T get_int() {
    T sign = 1;
    char c = peek();
    if (c == '+') {
      ignore();
    } else if (c == '-') {
      sign = -1;
      ignore();
    }
    return sign * get_uint<T>();
  }

  template <typename T>
  Input& get_int(T& n) {
    n = get_int<T>();
    return *this;
  }

  // Decodes the whole input file into a 2-dimensional grid.
  template <typename T>
  Grid<T> grid(std::function<T(const char c)> f) {
    std::vector<T> data;
    int width = 0;
    for (auto& s : *this) {
      if (!width) width = s.size();
      assert(s.size() == width);
      for (auto c : s) data.push_back(f(c));
    }
    assert(data.size() % width == 0);
    return Grid<T>(width, data.size() / width, data);
  }

  // Decodes the whole input file into a 2-dimensional grid of chars.
  Grid<char> grid();

  char operator*();
  operator bool();
};

}  // namespace aoc

#endif  // AOC_INPUT_H

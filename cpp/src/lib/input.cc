#include "input.h"

namespace aoc {

void Input::set_err() { err_ = true; }

Input::Input(const int year, const int day, const std::string& suffix)
    : year_(year), day_(day) {
  std::ostringstream ss;
  // TODO: Recurse up until ./inputs is found?
  ss << "../../../inputs/" << year << "/day";
  if (day < 10) {
    ss << "0" << day;
  } else {
    ss << day;
  }
  if (!suffix.empty()) ss << "-" << suffix;
  ss << ".txt";
  path_ = ss.str();
  in_.open(path_);
  assert(in_.is_open());
  assert(!in_.fail());
}

Input::~Input() { in_.close(); }

int Input::year() const { return year_; }
int Input::day() const { return day_; }
std::string Input::path() { return path_; }

std::istream& operator>>(std::istream& is, Input::Line& line) {
  return std::getline(is, line);
}

std::istream_iterator<Input::Line> Input::begin() {
  return std::istream_iterator<Input::Line>(in_);
}

std::istream_iterator<Input::Line> Input::end() const {
  return std::istream_iterator<Input::Line>();
}

char Input::peek() { return in_.peek(); }

char Input::get() { return in_.get(); }

Input& Input::get(char& c) {
  c = get();
  return *this;
}

Input& Input::get(std::string& s, const std::streamsize n) {
  s.clear();
  for (int i = 0; i < n; ++i) s.push_back(get());
  return *this;
}

Input& Input::get_line(std::string& s) {
  std::getline(in_, s);
  return *this;
}

std::string Input::get_line() {
  std::string s;
  get_line(s);
  return s;
}

Input& Input::get_token(std::string& s) {
  s.clear();
  while (isspace()) ignore();
  while (!isspace()) s.push_back(get());
  return *this;
}

std::string Input::get_token() {
  std::string s;
  get_token(s);
  return s;
}

bool Input::branch(const char c) {
  if (!is(c)) return false;
  in_.ignore();
  return true;
}

Input& Input::expect(const char c, const std::streamsize n) {
  for (int i = 0; i < n; i++) {
    if (get() != c) {
      set_err();
      break;
    }
  }
  return *this;
}

Input& Input::expect(const char* str) {
  for (; *str; ++str) expect(*str);
  return *this;
}

Input& Input::discard(const char c) {
  if (is(c)) in_.ignore();
  return *this;
}

Input& Input::discard_to(const char c) {
  while (in_)
    if (get() == c) break;
  return *this;
}

Input& Input::ignore(std::streamsize n, int delim) {
  in_.ignore(n, delim);
  return *this;
}

bool Input::is(const char c) { return peek() == c; }
bool Input::isdigit() { return std::isdigit(peek()); }
bool Input::isspace() { return std::isspace(peek()); }

Grid<char> Input::grid() {
  return grid<char>([](const char c) -> char { return c; });
}

char Input::operator*() { return peek(); }

Input::operator bool() { return !err_ && in_ && in_.peek() != EOF; }

}  // namespace aoc

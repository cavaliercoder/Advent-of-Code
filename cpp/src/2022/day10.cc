#include "lib/aoc.h"

namespace aoc2022 {

class Day10 {
  struct VM {
    int x = 1;
    int signal = 0;
    int clock = 0;
    std::string out = "";

    VM() { out.reserve(40 * 6 + 6); }

    void tick(int cycles) {
      for (int i = 0; i < cycles; ++i) {
        ++clock;
        if ((clock - 20) % 40 == 0) signal += x * clock;

        int crt_x = (clock - 1) % 40;
        out.append(x >= crt_x - 1 && x <= crt_x + 1 ? "#" : ".");
        if (crt_x == 39) out.append("\n");
      }
    }

    void noop() { tick(1); }

    void addx(const int n) {
      tick(2);
      x += n;
    }

    VM& run(aoc::Input& in) {
      std::string sarg;
      int iarg;
      while (in) {
        in.get_token(sarg);
        if (sarg == "noop") {
          noop();
          in.expect('\n');
          continue;
        }
        if (sarg == "addx") {
          in.expect(' ').get_int(iarg).expect('\n');
          addx(iarg);
          continue;
        }
        assert(0);
      }
      return *this;
    }
  };

 public:
  int Part1(aoc::Input in) { return VM().run(in).signal; }
  std::string Part2(aoc::Input in) { return VM().run(in).out; }
};

TEST(Day10, Part1) { EXPECT_EQ(Day10().Part1(aoc::Input(2022, 10)), 13740); }

TEST(Day10, Part2) {
  const std::string expect =
      "####.#..#.###..###..####.####..##..#....\n"
      "...#.#..#.#..#.#..#.#....#....#..#.#....\n"
      "..#..#..#.#..#.#..#.###..###..#....#....\n"
      ".#...#..#.###..###..#....#....#....#....\n"
      "#....#..#.#....#.#..#....#....#..#.#....\n"
      "####..##..#....#..#.#....####..##..####.\n";

  EXPECT_EQ(Day10().Part2(aoc::Input(2022, 10)), expect);
}

}  // namespace aoc2022
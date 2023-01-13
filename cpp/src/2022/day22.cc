#include "lib/aoc.h"

namespace aoc2022 {

struct Day22 {
  using Cube = aoc::Cube;
  using Grid = aoc::Grid<char>;
  using Point = aoc::Point<2, int>;

  enum Terrain : char {
    Air = ' ',
    Path = '.',
    Wall = '#',
  };

  Grid grid;
  std::vector<int> moves;
  std::unordered_map<Point, Cube> input_orientation;
  std::unordered_map<Cube, Point> face_at;
  int size;

  /*
   * NB: Grid has an inverted y-axis relative to Point.
   */

  static constexpr Point left = Point().left();
  static constexpr Point right = Point().right();
  static constexpr Point up = Point().down();
  static constexpr Point down = Point().up();

  Day22(aoc::Input in, const int size = 50) : size(size) {
    // Parse grid
    grid = Grid(size * 4, size * 4, Air);
    Point p;
    while (in) {
      while (in && !in.is('\n')) {
        grid[p] = in.get();
        p += {1, 0};
      }
      in.expect('\n');
      p = Point(0, p.y() + 1);
      if (in.branch('\n')) break;
    }

    // Parse moves, normalized to +n for Rn and -n for Ln.
    int sign = 1;
    while (in) {
      if (in.isdigit()) moves.push_back(sign * in.get_uint<int>());
      if (in.branch('\n')) break;
      if (in.branch('L')) {
        sign = -1;
      } else {
        sign = 1;
        in.expect('R');
      }
    }

    // Determine the orientation of each cube face in the input
    auto stack =
        aoc::Stack<std::pair<Point, Cube>>({grid.find(Path).point(), Cube()});
    std::unordered_set<Point> seen;
    while (stack) {
      auto [p, c] = stack.pop();
      if (seen.count(p) || !grid.contains(p) || grid[p] == Air) continue;
      seen.insert(p);
      input_orientation[p] = c;
      for (int i = 0; i < 4; ++i) {
        // All four rotations returns the same face.
        face_at[c = c.ccw()] = p;
      }
      stack.push({{p.x() - size, p.y()}, c.right()});
      stack.push({{p.x() + size, p.y()}, c.left()});
      stack.push({{p.x(), p.y() + size}, c.up()});
    }
  }

  // Translate a move into a x/y bearing from b.
  static Point bear(const Point b, const int move) {
    auto o = b.orientation();
    if (o == up) return Point(move, 0);
    if (o == down) return Point(-move, 0);
    if (o == left) return Point(0, -move);
    if (o == right) return Point(0, move);
    throw "bad bearing";
  }

  struct Step {
    Point p;  // Current point
    Point o;  // Orientation
    int d;    // Distance
  };

  // Move along a 2D grid, wrapping empty space.
  Step move2D(const Point src, const Point bearing) const {
    auto s = Step{src, bearing.orientation(), bearing.manhattan()};
    while (s.d) {
      Step next = {s.p + s.o, s.o, s.d - 1};
      if (!grid.contains(next.p) || grid[next.p] == Air) {
        next.p = s.p;
        while (grid.contains(next.p - next.o) && grid[next.p - next.o] != Air)
          next.p -= next.o;
      }
      if (grid[next.p] == Wall) break;
      assert(grid.contains(next.p));
      assert(grid[next.p] == Path);
      s = next;
    }
    return s;
  }

  // Move around a 3D cube.
  Step move3D(const Point src, const Point bearing) const {
    auto s = Step{src, bearing.orientation(), bearing.manhattan()};
    while (s.d) {
      Step next = {s.p + s.o, s.o, s.d - 1};
      if (!grid.contains(next.p) || grid[next.p] == Air) {
        auto face = (s.p / size) * size;  // Top-left point of cube face.
        assert(input_orientation.count(face));
        auto cube = input_orientation.at(face);  // Input rotation of the face
        Point entry = ((next.p - face + size) % size);  // Relative to the face

        // Rotate the cube and find the next face
        if (s.o == up) {
          cube = cube.down();
        } else if (s.o == down) {
          cube = cube.up();
        } else if (s.o == left) {
          cube = cube.right();
        } else if (s.o == right) {
          cube = cube.left();
        } else {
          throw "bad orientation";
        }
        assert(face_at.count(cube));
        face = face_at.at(cube);

        // Rotate the face to match the input grid.
        while (cube != input_orientation.at(face)) {
          cube = cube.ccw();
          next.o = next.o.cw();
          entry = entry.icw(size);
        }
        next.p = face + entry;
      }
      if (grid[next.p] == Wall) break;
      assert(grid.contains(next.p));
      assert(grid[next.p] == Path);
      s = next;
    }
    return s;
  }

  static int score(const Point p, const Point o) {
    int n = (1000 * (p.y() + 1)) + (4 * (p.x() + 1));
    if (o == right) return n;
    if (o == down) return n + 1;
    if (o == left) return n + 2;
    if (o == up) return n + 3;
    throw "bad bearing";
  }

  int Part1() const {
    Step s = {grid.find(Path).point(), up, 0};
    for (auto m : moves) {
      assert(grid[s.p] == Path);
      s = move2D(s.p, bear(s.o, m));
    }
    return score(s.p, s.o);
  }

  int Part2() const {
    Step s = {grid.find(Path).point(), up, 0};
    for (auto m : moves) {
      assert(grid[s.p] == Path);
      s = move3D(s.p, bear(s.o, m));
    }
    return score(s.p, s.o);
  };
};

TEST(Day22, Part1) { EXPECT_EQ(Day22(aoc::Input(2022, 22)).Part1(), 1484); }
TEST(Day22, Part2) { EXPECT_EQ(Day22(aoc::Input(2022, 22)).Part2(), 142228); }

}  // namespace aoc2022

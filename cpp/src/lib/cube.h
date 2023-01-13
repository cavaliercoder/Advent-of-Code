#ifndef AOC_CUBE_H
#define AOC_CUBE_H

#include <cstdint>
#include <iostream>
#include <vector>

#include "point.h"

namespace {

/*
 * Edges are numbered from the front floor, moving clockwise around the floor,
 * then the middle vertical edges starting from the front left, then moving
 * clockwise around the top edges from the front top.
 */

static constexpr uint8_t EDGE[24] = {0, 4, 8,  7, 1, 5, 9, 4, 2, 6, 10, 5,
                                     3, 7, 11, 6, 2, 1, 0, 3, 8, 9, 10, 11};

static constexpr bool edge_has(const uint8_t e, const uint8_t a,
                               const uint8_t b) {
  return EDGE[a] == e && EDGE[b] == e;
}

static_assert(edge_has(0, 0, 18));
static_assert(edge_has(1, 4, 17));
static_assert(edge_has(2, 8, 16));
static_assert(edge_has(3, 12, 19));
static_assert(edge_has(4, 1, 7));
static_assert(edge_has(5, 5, 11));
static_assert(edge_has(6, 9, 15));
static_assert(edge_has(7, 13, 3));
static_assert(edge_has(8, 2, 20));
static_assert(edge_has(9, 6, 21));
static_assert(edge_has(10, 10, 22));
static_assert(edge_has(11, 14, 23));

/*
 * Vertices are numbered from the front left floor, moving clockwise around the
 * floor, then the same around the top from the front top left.
 */

static constexpr uint8_t VERTEX[24] = {0, 4, 7, 3, 1, 5, 4, 0, 2, 6, 5, 1,
                                       3, 7, 6, 2, 1, 0, 3, 2, 4, 5, 6, 7};

static constexpr bool vertex_has(const uint8_t v, const uint8_t a,
                                 const uint8_t b, const uint8_t c) {
  return VERTEX[a] == v && VERTEX[b] == v && VERTEX[c] == v;
}

static_assert(vertex_has(0, 0, 7, 17));
static_assert(vertex_has(1, 4, 11, 16));
static_assert(vertex_has(2, 8, 15, 19));
static_assert(vertex_has(3, 3, 12, 18));
static_assert(vertex_has(4, 1, 6, 20));
static_assert(vertex_has(5, 5, 10, 21));
static_assert(vertex_has(6, 9, 14, 22));
static_assert(vertex_has(7, 2, 13, 23));

/*
 * The following tables map an old state to a new state, if rotated 90° in the
 * associated direction.
 *
 * Each state is denoted by the ID of the front left floor corner.
 *
 * The reference cube can be reconstructed by numbering the corner of each face
 * from the bottom left in a clockwise direction. Mark each face in the
 * following order: front, left, back, right, bottom, top.
 */

static constexpr uint8_t LEFT[24] = {12, 17, 6,  23, 0,  16, 10, 20,
                                     4,  19, 14, 21, 8,  18, 2,  22,
                                     15, 11, 7,  3,  13, 1,  5,  9};

static constexpr uint8_t RIGHT[24] = {4,  21, 14, 19, 8, 22, 2,  18,
                                      12, 23, 6,  17, 0, 20, 10, 16,
                                      5,  1,  13, 9,  7, 11, 15, 3};

static constexpr uint8_t UP[24] = {16, 5,  22, 15, 19, 9, 23, 3,
                                   18, 13, 20, 7,  17, 1, 21, 11,
                                   10, 6,  2,  14, 0,  4, 8,  12};

static constexpr uint8_t DOWN[24] = {20, 13, 18, 7,  21, 1,  17, 11,
                                     22, 5,  16, 15, 23, 9,  19, 3,
                                     0,  12, 8,  4,  10, 14, 2,  6};

static constexpr uint8_t CW[24] = {3,  0,  1,  2,  7,  4,  5,  6,
                                   11, 8,  9,  10, 15, 12, 13, 14,
                                   19, 16, 17, 18, 23, 20, 21, 22};

static constexpr uint8_t CCW[24] = {1,  2,  3,  0,  5,  6,  7,  4,
                                    9,  10, 11, 8,  13, 14, 15, 12,
                                    17, 18, 19, 16, 21, 22, 23, 20};

static constexpr bool check_rotation(const uint8_t table[24]) {
  int sum = 0;
  uint8_t count[24] = {};
  for (int i = 0; i < 24; i++) {
    if (count[table[i]]) return false;
    ++count[table[i]];
    sum += i;
  }
  if (sum != 276) return false;
  return true;
}

static_assert(check_rotation(LEFT));
static_assert(check_rotation(RIGHT));
static_assert(check_rotation(UP));
static_assert(check_rotation(DOWN));
static_assert(check_rotation(CW));
static_assert(check_rotation(CCW));

/*
 * XYZ components of each possible state.
 */

static constexpr int X[24] = {0, 0, 0, 0, 0,  0,  0,  0,  0, 0, 0, 0,
                              0, 0, 0, 0, -1, -1, -1, -1, 1, 1, 1, 1};

static constexpr int Y[24] = {0,  0,  0,  0,  1, 1, 1, 1, 2, 2, 2, 2,
                              -1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0};

static constexpr int Z[24] = {0, 1, 2, -1, 0, 1, 2, -1, 0, 1, 2, -1,
                              0, 1, 2, -1, 0, 1, 2, -1, 0, 1, 2, -1};

}  // namespace

namespace aoc {

// Cube represents the rotations of a simple 3-dimensional cube that can only be
// rotated in 90° increments. It has only 24 unique states.
class Cube {
  // State is an arbitrary value that encodes the distinct state of the cube.
  // Currently, the front bottom left corner ID is used.
  uint8_t state_ = 0;

 public:
  static const int MaxStates = 24;

  struct Hash {
    std::size_t operator()(Cube const& c) const noexcept { return c.state_; }
  };

  // Face represents one of 6 uniques faces on a cube.
  //
  // Each face is named for its initial position on the cube, independent of its
  // current position.
  enum class Face : uint8_t {
    Front = 0,
    Left = 1,
    Back = 2,
    Right = 3,
    Bottom = 4,
    Top = 5
  };

  friend std::ostream& operator<<(std::ostream& os, const Face& f) {
    static char const* s[6] = {"Front", "Left",   "Back",
                               "Right", "Bottom", "Top"};
    return os << s[int(f)];
  }

  // Vertex is an arbitrary identifier representing one of 8 unique vertices on
  // a cube.
  enum class Vertex : uint8_t {
    A = 0,
    B = 1,
    C = 2,
    D = 3,
    E = 4,
    F = 5,
    G = 6,
    H = 7,
  };

  friend std::ostream& operator<<(std::ostream& os, const Vertex& v) {
    return os.put('a' + char(v));
  }

  // Edge is an arbitrary identifier representing one of 12 unique edges on a
  // cube.
  enum class Edge : uint8_t {
    A = 0,
    B = 1,
    C = 2,
    D = 3,
    E = 4,
    F = 5,
    G = 6,
    H = 7,
    I = 8,
    J = 9,
    K = 10,
    L = 11,
  };

  friend std::ostream& operator<<(std::ostream& os, Edge& e) {
    return os.put('a' + char(e));
  }

  // Corner is an arbitrary value representing one of 24 unique corners on a
  // cube.
  enum class Corner : uint8_t {
    A = 0,
    B = 1,
    C = 2,
    D = 3,
    E = 4,
    F = 5,
    G = 6,
    H = 7,
    I = 8,
    J = 9,
    K = 10,
    L = 11,
    M = 12,
    N = 13,
    O = 14,
    P = 15,
    Q = 16,
    R = 17,
    S = 18,
    T = 19,
    U = 20,
    V = 21,
    W = 22,
    X = 23,
  };

  friend std::ostream& operator<<(std::ostream& os, const Corner& c) {
    return os.put('a' + char(c));
  }

  // Returns a front-facing cube.
  constexpr Cube() = default;

  // Returns a new cube from a known state.
  //
  // Note that the state value is an implementation detail that may change over
  // time.
  constexpr Cube(const uint8_t state) : state_(state % MaxStates) {}

  // Returns a cube constructed from a point generated by the xyz() function.
  template <typename T>
  constexpr Cube(const Point<3, T> xyz) {
    auto p = ((xyz % 4) + 4) % 4;
    for (int x = 0; x < p.x(); ++x) state_ = DOWN[state_];
    for (int y = 0; y < p.y(); ++y) state_ = RIGHT[state_];
    for (int z = 0; z < p.z(); ++z) state_ = CCW[state_];
  }

  // Returns a vector containing every possible cube state.
  static std::vector<Cube> all() {
    std::vector<Cube> a;
    for (int i = 0; i < MaxStates; ++i) a.push_back(Cube(i));
    return a;
  }

  // Returns the internal state value of the cube so it can be reconstructed.
  constexpr uint8_t state() const { return state_; }

  // Returns the ID of the face which is currently facing front.
  constexpr Face face() const { return Face(state_ / 4); }

  // Returns the ID of the vertex which is currently aligned to the bottom left
  // on the face at front.
  constexpr Vertex vertex() const { return Vertex(VERTEX[state_]); }

  // Returns the ID of the edge which is currently facing the front floor.
  constexpr Edge edge() const { return Edge(EDGE[state_]); }

  // Returns the ID of the corner which is currently aligned to the bottom left
  // on the face at the front.
  constexpr Corner corner() const { return Corner(state_); }

  // Returns the 3-D rotation of the cube as the number of 90° CCW rotations of
  // each axis, in decending order of precedence.
  constexpr Point<3, int> xyz() const {
    return {X[state_], Y[state_], Z[state_]};
  }

  // Pushes the front face of the cube to the left by 90°.
  constexpr Cube left() const { return Cube(LEFT[state_]); }

  // Pushes the front face of the cube to the right by 90°.
  constexpr Cube right() const { return Cube(RIGHT[state_]); }

  // Pushes the front face of the cube up 90°.
  constexpr Cube up() const { return Cube(UP[state_]); }

  // Pushes the front face of the cube down 90°.
  constexpr Cube down() const { return Cube(DOWN[state_]); }

  // Rotates the front face 90° clockwise.
  constexpr Cube cw() const { return Cube(CW[state_]); }

  // Rotates the front face 90° counterclockwise.
  constexpr Cube ccw() const { return Cube(CCW[state_]); }

  friend constexpr bool operator==(const Cube& lhs, const Cube& rhs) {
    return lhs.state_ == rhs.state_;
  }

  friend constexpr bool operator!=(const Cube& lhs, const Cube& rhs) {
    return lhs.state_ != rhs.state_;
  }

  friend std::ostream& operator<<(std::ostream& os, const Cube& cube) {
    return os << cube.corner();
  }
};

static_assert(sizeof(Cube) == 1);

}  // namespace aoc

// Make hasher available to STL.
template <>
struct std::hash<aoc::Cube> : aoc::Cube::Hash {};

#endif  // AOC_CUBE_H

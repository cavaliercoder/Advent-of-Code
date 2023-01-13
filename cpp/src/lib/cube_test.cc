#include "cube.h"

#include "testing.h"

namespace aoc {

TEST(AoC, Cube) {
  Cube c;

  // Comparison
  EXPECT_EQ(Cube(), Cube());
  EXPECT_NE(Cube(), Cube(1));
  EXPECT_EQ(Cube().down(), Cube().up().up().up());

  // Spin

  auto expect_face = [&](Cube& c, Cube::Face face) -> void {
    EXPECT_EQ(c.face(), face);

    // Rotate 90° and back.
    EXPECT_EQ(c.up().down().face(), face);
    EXPECT_EQ(c.down().up().face(), face);
    EXPECT_EQ(c.left().right().face(), face);
    EXPECT_EQ(c.right().left().face(), face);
    EXPECT_EQ(c.cw().ccw().face(), face);
    EXPECT_EQ(c.ccw().cw().face(), face);

    // Rotate 360°.
    EXPECT_EQ(c.up().up().up().up().face(), face)
    EXPECT_EQ(c.down().down().down().down().face(), face);
    EXPECT_EQ(c.left().left().left().left().face(), face);
    EXPECT_EQ(c.right().right().right().right().face(), face);
    EXPECT_EQ(c.cw().cw().cw().cw().face(), face);
    EXPECT_EQ(c.ccw().ccw().ccw().ccw().face(), face);
  };

  // Spin up
  expect_face(c, Cube::Face::Front);
  expect_face(c = c.up(), Cube::Face::Bottom);
  expect_face(c = c.up(), Cube::Face::Back);
  expect_face(c = c.up(), Cube::Face::Top);
  expect_face(c = c.up(), Cube::Face::Front);

  // Spin down
  expect_face(c, Cube::Face::Front);
  expect_face(c = c.down(), Cube::Face::Top);
  expect_face(c = c.down(), Cube::Face::Back);
  expect_face(c = c.down(), Cube::Face::Bottom);
  expect_face(c = c.down(), Cube::Face::Front);

  // Spin left
  expect_face(c, Cube::Face::Front);
  expect_face(c = c.left(), Cube::Face::Right);
  expect_face(c = c.left(), Cube::Face::Back);
  expect_face(c = c.left(), Cube::Face::Left);
  expect_face(c = c.left(), Cube::Face::Front);

  // Spin right
  expect_face(c, Cube::Face::Front);
  expect_face(c = c.right(), Cube::Face::Left);
  expect_face(c = c.right(), Cube::Face::Back);
  expect_face(c = c.right(), Cube::Face::Right);
  expect_face(c = c.right(), Cube::Face::Front);

  // XYZ rotation.
  for (int i = 0; i < 24; ++i) {
    auto a = Cube(i);
    auto b = Cube(a.xyz());
    EXPECT_EQ(a, b);
  }

  EXPECT_EQ(Cube(Point<3, int>(0, 1, 2)), Cube(6));
}

}  // namespace aoc

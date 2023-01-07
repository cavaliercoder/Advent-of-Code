#include "lib/aoc.h"

namespace aoc2022 {

class Day17 {
  static constexpr int kSpriteHeight = 4;
  static constexpr int kStartXOffset = 3;    // Wall + padding.
  static constexpr int kStartYOffset = 3;    // Start heigh above the chamber.
  static constexpr uint16_t kWall = 0x8080;  // |.......|.......
  static constexpr char kLeft = '<';
  static constexpr char kRight = '>';
  static constexpr int kHashRows = 50;  // Collisions at < 13 on input data.

  struct Tetromino {
    // Encodes the tetromino as a mask of the enclosing chmaber, including
    // horizontal position, etc.
    uint16_t rows[kSpriteHeight];

    // Parse a tetromino from a 4x4 grid of bits in a uint16_t.
    explicit Tetromino(const uint16_t hash = 0) {
      rows[0] = (hash >> 12 & 0x0F) << (12 - kStartXOffset);  // Top
      rows[1] = (hash >> 8 & 0x0F) << (12 - kStartXOffset);
      rows[2] = (hash >> 4 & 0x0F) << (12 - kStartXOffset);
      rows[3] = (hash >> 0 & 0x0F) << (12 - kStartXOffset);  // Bottom
    }

    Tetromino hmove(const char dir) {
      auto t = Tetromino(*this);
      if (dir == kLeft) {
        for (int i = 0; i < kSpriteHeight; ++i) t.rows[i] <<= 1;
      } else {
        for (int i = 0; i < kSpriteHeight; ++i) t.rows[i] >>= 1;
      }
      return t;
    }
  };

  uint64_t get_score(const std::vector<uint16_t>& chamber) {
    uint64_t score = chamber.size() - 1;  // Less floor
    while (chamber[score] == kWall) --score;
    return score;
  }

  bool check_collision(const std::vector<uint16_t>& chamber, const Tetromino& t,
                       const int row) {
    for (int i = 0; i < kSpriteHeight; ++i)
      if ((chamber[row - i] & t.rows[i])) return true;
    return false;
  }

  std::string compute_hash(const std::vector<uint16_t>& chamber,
                           uint32_t tetramino, uint32_t move,
                           const int hash_rows = kHashRows) {
    auto h = aoc::SHA1();
    h.update((const char*)&tetramino, sizeof(uint32_t));
    h.update((const char*)&move, sizeof(uint32_t));
    for (int i = chamber.size() - 1; i > chamber.size() - hash_rows; --i) {
      h.update((const char*)&chamber[i], sizeof(uint16_t));
    }
    return h.str();
  }

  uint64_t run(aoc::Input& in, const uint64_t count) {
    static const std::array<Tetromino, 5>& Queue = {
        Tetromino(0x000F), Tetromino(0x04E4), Tetromino(0x022E),
        Tetromino(0x8888), Tetromino(0x00CC)};

    auto moves = in.get_line();
    uint64_t score_boost = 0;
    std::vector<uint16_t> chamber = {0xFFFF /* Floor */};
    for (int i = 0; i < kStartYOffset + kSpriteHeight; ++i)
      chamber.push_back(kWall);

    // Map chamber state hash to source iteration and score.
    std::unordered_map<std::string, std::pair<uint64_t, uint64_t>> hashes = {};

    uint64_t mi = 0;  //
    for (uint64_t ti = 0; ti < count; ++ti) {
      // Add headroom
      int row = chamber.size() - 1;
      for (int i = 0; i < kStartYOffset + kSpriteHeight; ++i)
        if (chamber[row - i] != kWall) chamber.push_back(kWall);

      // Send next tetromino
      Tetromino t = Queue[ti % Queue.size()];
      for (row = chamber.size() - 1; row > 0; --row) {
        Tetromino lr = t.hmove(moves[mi++ % moves.size()]);
        if (!check_collision(chamber, lr, row)) t = lr;
        if (!check_collision(chamber, t, row - 1)) continue;
        for (int i = 0; i < kSpriteHeight; ++i) chamber[row - i] |= t.rows[i];
        break;
      }

      // Hash the current state and check for loops
      if (chamber.size() < kHashRows) continue;
      auto h = compute_hash(chamber, ti % Queue.size(), mi % moves.size());
      auto it = hashes.find(h);
      if (it == hashes.end()) {
        hashes.insert({h, {ti, get_score(chamber)}});
        continue;
      }
      auto prev_ti = it->second.first;
      auto prev_score = it->second.second;

      // Fast-forward
      auto period = ti - prev_ti;
      auto loops = (count - ti) / period;
      score_boost = loops * (get_score(chamber) - prev_score);
      ti += loops * period;
      hashes.clear();  // Effect: no further loops
    }

    return get_score(chamber) + score_boost;
  }

 public:
  uint64_t Part1(aoc::Input in) { return run(in, 2022); }
  uint64_t Part2(aoc::Input in) { return run(in, 1000000000000); };
};

TEST(Day17, Part1) { EXPECT_EQ(Day17().Part1(aoc::Input(2022, 17)), 3065); }

TEST(Day17, Part2) {
  EXPECT_EQ(Day17().Part2(aoc::Input(2022, 17)), 1562536022966);
}

}  // namespace aoc2022

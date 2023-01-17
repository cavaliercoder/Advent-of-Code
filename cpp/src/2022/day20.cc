#include <iostream>
#include <vector>

#include "lib/aoc.h"

namespace aoc2022 {

class Day20 {
  struct Node {
    int64_t id = 0;
    int64_t value = 0;
    Node* prev = nullptr;
    Node* next = nullptr;

    void remove() {
      prev->next = next;
      next->prev = prev;
      next = nullptr;
      prev = nullptr;
    }

    void insert(Node* after) {
      prev = after;
      next = after->next;
      next->prev = this;
      after->next = this;
    }

    Node* nextn(int64_t n) {
      Node* node = this;
      for (uint64_t i = 0; i < std::abs(n); ++i) {
        if (n > 0) {
          node = node->next;
        } else {
          node = node->prev;
        }
      }
      return node;
    }
  };

  std::vector<Node> make_nodes(const int64_t key = 1) const {
    auto node =
        std::vector<Node>(numbers_.size(), Node());  // needs pointer stability
    for (int i = 0; i < numbers_.size(); ++i) {
      node[i].id = i;
      node[i].value = numbers_[i] * key;
      node[i].prev = &node[(i - 1 + numbers_.size()) % numbers_.size()];
      node[i].next = &node[(i + 1 + numbers_.size()) % numbers_.size()];
    }
    return node;
  }

  int64_t mix(const int rounds = 1, const int64_t key = 1) const {
    auto node = make_nodes(key);
    int64_t len = numbers_.size() - 1;
    Node* zero = nullptr;
    for (int r = 0; r < rounds; ++r) {
      for (int64_t id = 0; id < numbers_.size(); ++id) {
        auto& n = node[id];
        if (n.value == 0) zero = &n;
        auto after = n.prev;
        n.remove();
        after = after->nextn(n.value % len);
        n.insert(after);
      }
    }
    uint64_t sum = 0;
    zero = zero->nextn(1000 % numbers_.size());
    sum += zero->value;
    zero = zero->nextn(1000 % numbers_.size());
    sum += zero->value;
    zero = zero->nextn(1000 % numbers_.size());
    sum += zero->value;
    return sum;
  }

  std::vector<int64_t> numbers_;

 public:
  Day20(aoc::Input in) {
    int64_t n;
    while (in) {
      in.get_int(n).expect('\n');
      numbers_.push_back(n);
    }
  }

  int64_t Part1() const { return mix(); }
  int64_t Part2() const { return mix(10, 811589153); };
};

TEST(Day20, Part1) { EXPECT_EQ(Day20(aoc::Input(2022, 20)).Part1(), 3466); }

TEST(Day20, Part2) {
  EXPECT_EQ(Day20(aoc::Input(2022, 20)).Part2(), 9995532008348);
}

}  // namespace aoc2022

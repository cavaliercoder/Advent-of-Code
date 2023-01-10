#include <iostream>
#include <vector>

#include "lib/aoc.h"

namespace aoc2022 {

class Day20 {
  std::vector<int64_t> parse(aoc::Input& in) {
    std::vector<int64_t> a;
    int64_t n;
    while (in) {
      in.get_int(n).expect('\n');
      a.push_back(n);
    }
    return a;
  }

  struct Node {
    const int64_t id = 0;
    const int64_t value = 0;

    Node* prev = nullptr;
    Node* next = nullptr;

    Node(const int64_t id, const int64_t value) : id(id), value(value) {}

    ~Node() {}  // TODO

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

  std::unordered_map<int64_t, Node*> make_nodes(const std::vector<int64_t>& a,
                                                const int64_t key = 1) {
    std::unordered_map<int64_t, Node*> nodes;
    auto first = new Node(0, a[0] * key);
    nodes[0] = first;
    auto prev = first;
    for (int i = 1; i < a.size(); ++i) {
      auto node = new Node(i, a[i] * key);
      nodes[i] = node;
      prev->next = node;
      node->prev = prev;
      prev = node;
    }
    prev->next = first;
    first->prev = prev;
    return nodes;
  }

  int64_t mix(std::vector<int64_t>& a, const int rounds = 1,
              const int64_t key = 1) {
    int64_t len = a.size() - 1;
    auto nodes = make_nodes(a, key);
    Node* zero = nullptr;
    for (int r = 0; r < rounds; ++r) {
      for (int64_t id = 0; id < a.size(); ++id) {
        auto n = nodes[id];
        if (n->value == 0) zero = n;
        auto after = n->prev;
        n->remove();
        after = after->nextn(n->value % len);
        n->insert(after);
      }
    }

    uint64_t sum = 0;
    zero = zero->nextn(1000 % a.size());
    sum += zero->value;
    zero = zero->nextn(1000 % a.size());
    sum += zero->value;
    zero = zero->nextn(1000 % a.size());
    sum += zero->value;
    return sum;
  }

 public:
  int64_t Part1(aoc::Input in) {
    auto a = parse(in);
    return mix(a);
  }

  int64_t Part2(aoc::Input in) {
    auto a = parse(in);
    return mix(a, 10, 811589153);
  };
};

TEST(Day20, Part1) { EXPECT_EQ(Day20().Part1(aoc::Input(2022, 20)), 3466); }

TEST(Day20, Part2) {
  EXPECT_EQ(Day20().Part2(aoc::Input(2022, 20)), 9995532008348);
}

}  // namespace aoc2022

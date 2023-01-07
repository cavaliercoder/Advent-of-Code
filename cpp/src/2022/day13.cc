#include "lib/aoc.h"

namespace aoc2022 {

class Day13 {
  struct Node {
    int value = 0;
    bool isnumber = false;
    std::vector<Node*> children;

    Node() = default;

    Node(const int n, const int depth = 0) {
      auto node = this;
      for (int i = 0; i < depth; ++i) {
        node->children.push_back(new Node());
        node = node->children.back();
      }
      node->value = n;
      node->isnumber = true;
      node->children = {node};
    }

    ~Node() {
      if (!isnumber)
        for (auto child : children) delete child;
    }

    int size() const { return isnumber ? 0 : children.size(); }

    friend int cmp(const std::vector<Node*>& a, const std::vector<Node*>& b) {
      for (int i = 0; i < a.size(); ++i) {
        if (i >= b.size()) return 1;
        if (a[i]->isnumber && b[i]->isnumber) {
          if (a[i]->value < b[i]->value) return -1;
          if (a[i]->value > b[i]->value) return 1;
        } else {
          auto n = cmp(a[i]->children, b[i]->children);
          if (n != 0) return n;
        }
      }
      if (b.size() > a.size()) return -1;
      return 0;
    }

    friend bool operator<(const Node& a, const Node& b) {
      return cmp(a.children, b.children) < 0;
    }

    friend bool operator==(const Node& a, const Node& b) {
      return cmp(a.children, b.children) == 0;
    }

    friend std::ostream& operator<<(std::ostream& os, const Node& node) {
      if (node.isnumber) return os << node.value;
      os << "[";
      for (int i = 0; i < node.size(); ++i) {
        if (i) os << ",";
        os << *node.children[i];
      }
      return os << "]";
    }

    static Node* parse(aoc::Input& in, const bool lf = true) {
      if (in.isdigit()) return new Node(in.get_uint<int>());
      auto node = new Node();
      in.expect('[');
      while (in) {
        if (in.branch(']')) break;
        node->children.push_back(parse(in, false));
        in.discard(',');
      }
      if (lf) in.expect('\n');
      return node;
    }
  };

 public:
  int Part1(aoc::Input in) {
    int i = 0, answer = 0;
    while (in) {
      auto a = Node::parse(in);
      auto b = Node::parse(in);
      in.expect('\n');
      if (*a < *b) answer += i + 1;
      delete a;
      delete b;
      ++i;
    }
    return answer;
  }

  int Part2(aoc::Input in) {
    auto div1 = new Node(2, 2), div2 = new Node(6, 2);
    std::vector<Node*> nodes = {div1, div2};
    while (in) {
      in.discard('\n');
      nodes.push_back(Node::parse(in));
    }
    std::sort(nodes.begin(), nodes.end(),
              [](const Node* a, const Node* b) -> bool { return *a < *b; });
    int answer = 1;
    for (int i = 0; i < nodes.size(); ++i) {
      if (nodes[i] == div1) answer *= i + 1;
      if (nodes[i] == div2) answer *= i + 1;
    }
    for (int i = 0; i < nodes.size(); ++i) delete nodes[i];
    return answer;
  }
};

TEST(Day13, Part1) { EXPECT_EQ(Day13().Part1(aoc::Input(2022, 13)), 5503); }
TEST(Day13, Part2) { EXPECT_EQ(Day13().Part2(aoc::Input(2022, 13)), 20952); }

}  // namespace aoc2022

#include <algorithm>
#include <unordered_map>

#include "lib/aoc.h"

namespace aoc2022 {

class Day07 {
  struct Node {
    Node* parent;
    std::string name;
    int size;
    bool is_dir;
    std::unordered_map<std::string, Node*> children;

    ~Node() {
      for (const auto& [_, child] : children) delete child;
    }

    void append(Node* child) {
      child->parent = this;
      children.insert({child->name, child});
      for (auto node = this; node != nullptr; node = node->parent)
        node->size += child->size;
    }

    static Node* parse(aoc::Input& in) {
      Node* root = new Node{nullptr, "", 0, true, {}};
      Node* cwd = nullptr;
      int iarg;
      std::string sarg;

      while (in) {
        if (!in.expect('$').expect(' ')) break;
        sarg = in.get_token();

        // cd <path>
        if (sarg == "cd") {
          in.get_token(sarg).expect('\n');
          if (sarg == "/") {
            cwd = root;
          } else if (sarg == "..") {
            cwd = cwd->parent;
          } else {
            cwd = cwd->children.at(sarg);
          }
          continue;
        }

        // ls
        if (sarg == "ls") {
          in.expect('\n');
          while (in && in.peek() != '$') {
            // <size> <name>
            if (in.isdigit()) {
              in.get_uint(iarg).get_token(sarg).expect('\n');
              cwd->append(new Node{cwd, sarg, iarg, false, {}});
            }

            // dir <name>
            if (in.branch('d')) {
              in.get_token(sarg).get_token(sarg).expect('\n');
              cwd->append(new Node{cwd, sarg, 0, true, {}});
              continue;
            }
          }
          continue;
        }
      }
      return root;
    }
  };

  int sum_dirs(const Node* n, const int limit = 100000) {
    if (!n->is_dir) return 0;
    int sum = 0;
    if (n->size < limit) sum += n->size;
    for (const auto& [_, child] : n->children) sum += sum_dirs(child);
    return sum;
  }

  int find_best(const Node* n, const int needed) {
    if (!n->is_dir) return 0;
    int best = n->size;
    for (const auto& [_, child] : n->children) {
      int child_best = find_best(child, needed);
      if (child_best > needed && child_best < best) best = child_best;
    }
    return best;
  }

 public:
  int Part1(aoc::Input in) {
    auto root = *Node::parse(in);
    return sum_dirs(&root);
  }

  int Part2(aoc::Input in) {
    auto root = Node::parse(in);
    int needed = 30000000 - (70000000 - root->size);
    return find_best(root, needed);
  }
};

TEST(Day07, Part1) { EXPECT_EQ(Day07().Part1(aoc::Input(2022, 7)), 1427048); }
TEST(Day07, Part2) { EXPECT_EQ(Day07().Part2(aoc::Input(2022, 7)), 2940614); }

}  // namespace aoc2022

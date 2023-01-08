#ifndef AOC_GRAPH_H
#define AOC_GRAPH_H

#include <unordered_map>

namespace aoc {

// Computes the shortest path weight between all vertices in a graph in Θ(|V|³).
//
// Edges are provided as an AdjacencyList keyed by Key, which should map source
// vertices to their connected vertices, to the weight of each connected edge.
// For example, edges[u][v] should return the weight of the edge from u to v.
//
// The inf value should be high enough that it will exceed the maximum weight of
// any path through the graph but not so high that 2*inf creates an integer
// overflow or sign change.
//
// The returned map can be used to find the shortest path between any two
// points.
//
// To find the the shortest distance between vertices v and u:
//
//     auto weights = floyd_warshall<Key>(edges);
//     auto best = weight[v][u];
//
template <typename Key, typename Weight = int,
          typename AdjacencyList =
              std::unordered_map<Key, std::unordered_map<Key, Weight>>>
std::unordered_map<Key, std::unordered_map<Key, Weight>> floyd_warshall(
    AdjacencyList& edges, const Weight inf = 1 << 30) {
  auto distance = std::unordered_map<Key, std::unordered_map<Key, Weight>>();
  for (const auto& [v, _] : edges)
    for (const auto& [u, _] : edges) distance[v][u] = inf;
  for (const auto& [v, ve] : edges) {
    distance[v][v] = 0;
    for (const auto& [u, w] : ve) distance[v][u] = w;
  }
  for (const auto& [k, _] : edges)
    for (const auto& [i, _] : edges)
      for (const auto& [j, _] : edges)
        if (distance[i][k] < inf && distance[k][j] < inf)
          distance[i][j] =
              std::min(distance[i][j], distance[i][k] + distance[k][j]);
  return distance;
}

}  // namespace aoc

#endif  // AOC_GRAPH_H

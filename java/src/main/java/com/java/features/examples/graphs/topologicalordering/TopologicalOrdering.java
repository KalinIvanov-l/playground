package com.java.features.examples.graphs.topologicalordering;

import com.java.features.examples.graphs.topologicalordering.shortestpath.Vertex;
import lombok.AllArgsConstructor;
import lombok.Getter;

import java.util.Deque;

@Getter
@AllArgsConstructor
public record TopologicalOrdering(Deque<Vertex> deque) {
  public void dfs(Vertex vertex) {
    vertex.isVisited();

    for (Vertex vertex1 : vertex.neighborliness()) {
      if (!vertex.isVisited()) {
        dfs(vertex1);
      }
      deque.push(vertex);
    }
  }
}

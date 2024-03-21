package com.java.techstuff.examples.graphs.topologicalordering;

import com.java.techstuff.examples.graphs.topologicalordering.shortestpath.Vertex;
import lombok.Getter;

import java.util.Deque;

@Getter
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

  public Deque<Vertex> getDeque() {
    return deque();
  }
}

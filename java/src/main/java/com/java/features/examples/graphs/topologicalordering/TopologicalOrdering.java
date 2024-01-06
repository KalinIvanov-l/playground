package com.java.features.examples.graphs.topologicalordering;

import lombok.AllArgsConstructor;
import lombok.Getter;

import java.util.Deque;

@Getter
@AllArgsConstructor
public record TopologicalOrdering(Deque<Vertex> deque) {
  public void dfs(Vertex vertex) {
    vertex.visited();

    for (Vertex vertex1 : vertex.neighbotList()) {
      if (!vertex.visited()) {
        dfs(vertex1);
      }
      deque.push(vertex);
    }
  }
}

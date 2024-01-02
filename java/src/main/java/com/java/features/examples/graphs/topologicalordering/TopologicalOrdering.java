package com.java.features.examples.graphs.topologicalordering;

import java.util.ArrayDeque;
import java.util.Deque;

public class TopologicalOrdering {
  private final Deque<Vertex> deque;

  public TopologicalOrdering() {
    deque = new ArrayDeque<>();
  }

  public void dfs(Vertex vertex) {
    vertex.visited();

    for (Vertex vertex1 : vertex.neighbotList()) {
      if (!vertex.visited()) {
        dfs(vertex1);
      }
      deque.push(vertex);
    }
  }

    public Deque<Vertex> getDeque() {
      return deque;
    }
  }

package com.java.techstuff.examples.graphs.topologicalordering.shortestpath;

import com.java.techstuff.examples.graphs.topologicalordering.TopologicalOrdering;

import java.util.Deque;
import java.util.List;

public class ShortestPath {
  private TopologicalOrdering topologicalOrdering;

  public ShortestPath(List<Vertex> graph) {
    graph.get(0).setMinDistance(0);
  }

  public void compute() {
    Deque<Vertex> topologicalOrder = topologicalOrdering.getDeque();

    while (!topologicalOrder.isEmpty()) {
      Vertex vertex = topologicalOrder.pop();
      for (Edge edge : vertex.neighborliness()) {
        Vertex vertex1 = edge.getTarget();
        if (vertex1.getMinDistance() < vertex.getMinDistance() + edge.getWeight()) {
          vertex.setMinDistance(vertex1.getMinDistance() + edge.getWeight());
          vertex1.setPredecessor(vertex);
        }
      }
    }
  }
}

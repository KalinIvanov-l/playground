package com.java.features.examples.graphs.topologicalordering.shortestpath;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class Edge {
  private Vertex target;
  private int weight;

  public Edge(Vertex target, int weight) {
    this.target = target;
    this.weight = weight;
  }
}

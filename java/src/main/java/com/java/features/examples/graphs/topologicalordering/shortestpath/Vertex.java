package com.java.features.examples.graphs.topologicalordering.shortestpath;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.util.ArrayList;
import java.util.List;

@Getter
@Setter
@ToString
public class Vertex extends Edge {
  private String name;
  @Getter
  private boolean visited;
  private int minDistance;
  private Vertex predecessor;
  private List<Edge> adjancencyList;
  private int weight;

  public Vertex(String name, Vertex predecessor, int weight) {
    super(predecessor, weight);
    this.name = name;
    minDistance = Integer.MAX_VALUE;
    adjancencyList = new ArrayList<>();
  }

  public Vertex[] neighborliness() {
    return new Vertex[0];
  }
}

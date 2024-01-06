package com.java.features.examples.graphs.topologicalordering.shortestpath;

import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.List;

@Getter
@Setter
public class Vertex {
  private String name;
  private boolean visited;
  private int minDistance;
  private Vertex predecessor;
  private List<Edge> adjancencyList;

  public Vertex(String name) {
    this.name = name;
    minDistance = Integer.MAX_VALUE;
    adjancencyList = new ArrayList<>();
  }

  @Override
  public String toString() {
    return name + "-" + predecessor;
  }
}

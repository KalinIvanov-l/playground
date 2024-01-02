package com.java.features.examples.graphs.topologicalordering;

import java.util.List;

public record Vertex(
        String name,
        boolean visited,
        List<Vertex> neighbotList
)
{
  public void addNeighbot(Vertex vertex) {
    neighbotList.add(vertex);
  }
}

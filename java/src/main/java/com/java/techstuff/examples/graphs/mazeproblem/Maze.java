package com.java.techstuff.examples.graphs.mazeproblem;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class Maze {
  private final int[][] labyrinth;
  private final int startRow;
  private final int startCol;
  private final boolean[][] visited;
  private final int[][] directions = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

  public Maze(int[][] maze, int startRow, int startCol) {
    this.labyrinth = maze;
    this.visited = new boolean[maze.length][maze[0].length];
    this.startRow = startRow;
    this.startCol = startCol;
  }

  public boolean findWayInMaze() {
    return explore(startRow, startCol);
  }

  private boolean explore(int row, int col) {
    if (row < 0 || col < 0 || row >= labyrinth.length
            || col >= labyrinth[0].length || labyrinth[row][col] == 0 || visited[row][col]) {
      return false;
    }
    if (labyrinth[row][col] == 2) {
      return true;
    }

    visited[row][col] = true;
    for (int[] direction : directions) {
      int nextRow = row + direction[0];
      int nextCol = col + direction[1];

      if (explore(nextRow, nextCol)) {
        return true;
      }
    }
    return false;
  }

  public static void main(String[] args) {
    int[][] map = {
            {1, 1, 1, 1, 1, 1},
            {2, 1, 0, 0, 0, 1},
            {0, 1, 0, 1, 0, 1},
            {0, 1, 0, 1, 0, 0},
            {0, 1, 0, 1, 1, 0},
            {0, 0, 0, 1, 1, 0}
    };
    var maze = new Maze(map, 1, 0);
    boolean pathFound = maze.findWayInMaze();
    log.info("Path found: {}", pathFound);
  }
}

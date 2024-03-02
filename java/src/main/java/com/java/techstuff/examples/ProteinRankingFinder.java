package com.java.techstuff.examples;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class ProteinRankingFinder {
  private static final Logger LOG = LoggerFactory.getLogger(ProteinRankingFinder.class);
  private static final String NOT_FOUND = "No major parent found";

  public static int[] findDualChildren(int[] arr) {
    int majorParent = findMajorParent(arr);
    Map<Integer, Integer> valueToIndex = new HashMap<>();
    for (int i = 0; i < arr.length; i++) {
      int complement = majorParent - arr[i];
      if (valueToIndex.containsKey(complement) && valueToIndex.get(complement) != i) {
        return new int[]{valueToIndex.get(complement), i};
      }
      valueToIndex.put(arr[i], i);
    }
    throw new IllegalArgumentException("Dual Child isn't found");
  }

  private static int findMajorParent(int[] arr) {
    Map<Integer, Integer> frequency = new HashMap<>();
    for (int value : arr) {
      frequency.put(value, frequency.getOrDefault(value, 0) + 1);
    }

    for (Map.Entry<Integer, Integer> entry : frequency.entrySet()) {
      if (entry.getValue() > 1) {
        return entry.getKey();
      }
    }
    return Arrays.stream(arr).max().orElseThrow(() -> new IllegalArgumentException(NOT_FOUND));
  }

  public static void main(String[] args) {
    int[] testArray1 = {3, 5, 1, 2, 8, 6, 9, 5, 10};
    LOG.info("Output 1: {}", Arrays.toString(findDualChildren(testArray1)));

    int[] testArray2 = {3, 11, 5, 1, 2, 8, 6, 9, 10, 4, 7};
    LOG.info("Output 2: {}", Arrays.toString(findDualChildren(testArray2)));
  }
}

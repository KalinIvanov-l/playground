package com.java.features.examples.advent_of_code;

import lombok.extern.slf4j.Slf4j;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;
import java.util.Map;
import java.util.Objects;

@Slf4j
public class SnowCalibration {
  private static final Map<String, String> numberWordsToDigits = Map.of(
          "one",   "one1one",
          "two",   "two2two",
          "three", "three3three",
          "four",  "four4four",
          "five",  "five5five",
          "six",   "six6six",
          "seven", "seven7seven",
          "eight", "eight8eight",
          "nine",  "nine9nine"
  );

  public static void main(String[] args) {
    var fileName = "java/src/main/resources/input.txt";
    var totalSum = 0;

    try (var reader = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = reader.readLine()) != null) {
        totalSum += extractNumberValue(line);
      }
    } catch (IOException e) {
      log.error("Error occurs while reading", e);
    }
    log.info("Total sum of calibration values: {}", totalSum);
  }

  private static int extractNumberValue(String line) {
    var numbers = Arrays.stream(line.split("\\W+"))
            .filter(word -> !word.isEmpty())
            .map(word -> word.matches("\\d+") ? word : numberWordsToDigits.get(word.toLowerCase()))
            .filter(Objects::nonNull)
            .mapToInt(Integer::parseInt)
            .toArray();

    if (numbers.length < 2) {
      return numbers.length == 1 ? numbers[0] : 0;
    }
    return numbers[0] + numbers[numbers.length - 1];
  }
}
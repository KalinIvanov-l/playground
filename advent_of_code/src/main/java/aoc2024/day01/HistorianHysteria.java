package aoc2024.day01;

import lombok.extern.slf4j.Slf4j;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

@Slf4j
public class HistorianHysteria {

    public static void main(String[] args) throws IOException {
        List<String> lines = Files.readAllLines(Path.of("advent_of_code/src/main/resources/input.txt"));
        process(lines);
    }

    private static void process(List<String> lines) {
        List<int[]> pairs = lines.stream()
                .map(line -> line.trim().split("\\s+"))
                .filter(parts -> parts.length == 2)
                .map(parts -> new int[]{Integer.parseInt(parts[0]), Integer.parseInt(parts[1])})
                .toList();

        List<Integer> list1 = pairs.stream().map(pair -> pair[0]).sorted().collect(Collectors.toList());
        List<Integer> list2 = pairs.stream().map(pair -> pair[1]).sorted().collect(Collectors.toList());

        Map<Integer, Long> counts2 = pairs.stream()
                .collect(Collectors.groupingBy(pair -> pair[1], Collectors.counting()));

        int part1 = calculatePart1(list1, list2);
        int part2 = calculatePart2(list1, counts2);

        log.info("Part 1: {}", part1);
        log.info("Part 2: {}", part2);
    }

    private static int calculatePart1(List<Integer> list1, List<Integer> list2) {
        return IntStream.range(0, list1.size())
                .map(i -> Math.abs(list2.get(i) - list1.get(i)))
                .sum();
    }

    private static int calculatePart2(List<Integer> list1, Map<Integer, Long> counts2) {
        return list1.stream()
                .mapToInt(n -> n * counts2.getOrDefault(n, 0L).intValue())
                .sum();
    }
}

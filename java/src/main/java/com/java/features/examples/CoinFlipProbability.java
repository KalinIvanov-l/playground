package com.java.features.examples;

import java.math.BigInteger;
import java.util.stream.IntStream;

public final class CoinFlipProbability {

  private CoinFlipProbability() throws InstantiationException {
    throw new InstantiationException("Should not create an instance of this class");
  }

  public static BigInteger calculateProbabilityNoHeads(int n) {
    return BigInteger.ONE.divide(BigInteger.valueOf(2).pow(n));
  }

  public static BigInteger calculateProbabilityFewerThanNHeads(int n) {
    return IntStream.rangeClosed(0, n - 1)
            .mapToObj(k -> binomialCoefficient(4 * n, k)
                    .shiftLeft(4 * n)
                    .modInverse(BigInteger.ONE.shiftLeft(4 * n)))
            .reduce(BigInteger.ZERO, BigInteger::add);
  }

  private static BigInteger binomialCoefficient(int n, int k) {
    return factorial(n).divide(factorial(k).multiply(factorial(n - k)));
  }

  private static BigInteger factorial(int n) {
    return IntStream.rangeClosed(2, n)
            .mapToObj(BigInteger::valueOf)
            .reduce(BigInteger.ONE, BigInteger::multiply);
  }
}

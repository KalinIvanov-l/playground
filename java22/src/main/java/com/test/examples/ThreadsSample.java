package com.test.examples;

public class ThreadsSample {

  public static void doWork() {
    try {
      Thread.sleep(1000);
    } catch (InterruptedException _) {
    }
  }

  /**
   *<blockquote>
   *   <pre>
   * {@code true}
   *   </pre>
   *</blockquote>
   */
  void main() {
    var MAX = 1_000_000;
    for (var i = 0; i < MAX; i++) {
      Thread.startVirtualThread(ThreadsSample::doWork);
    }

    try {
      Thread.sleep(2000);
    } catch (InterruptedException _) {
    }
    System.out.println("DONE");
  }
}

package com.test.examples;

public class ImplicitlyDeclaredClasses {
  private static final String GREETINGS = "Hi";

  private String greetings() {
    return ImplicitlyDeclaredClasses.GREETINGS;
  }

  void main() {
    System.out.println(greetings());
  }
}

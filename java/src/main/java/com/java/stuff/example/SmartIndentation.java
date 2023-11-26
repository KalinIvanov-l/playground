package com.java.stuff.example;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SmartIndentation {
  private static final Logger LOGGER = LoggerFactory.getLogger(SmartIndentation.class);
  public static String smartIndentation() {
    return """
            It is great
            when compilers care about conventions
            Makes our life easier
            """;
  }

  public static void main(String[] args) {
    LOGGER.info(smartIndentation());
  }
}

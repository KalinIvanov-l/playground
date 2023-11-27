package com.java.features.examples.textblock;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SmartIndentation {
  private static final Logger LOGGER = LoggerFactory.getLogger(SmartIndentation.class);
  private static final String MESSAGE = """
            It is great
            when compilers care about conventions
            Makes our life easier
            """;

  public static String smartIndentation() {
    return MESSAGE;
  }

  public static void main(String[] args) {
    LOGGER.info(smartIndentation());
  }
}

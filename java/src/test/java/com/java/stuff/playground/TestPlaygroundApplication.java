package com.java.stuff.playground;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.test.context.TestConfiguration;

@TestConfiguration(proxyBeanMethods = false)
public class TestPlaygroundApplication {

  public static void main(String[] args) {
    SpringApplication.from(PlaygroundApplication::main).with(TestPlaygroundApplication.class).run(args);
  }

}

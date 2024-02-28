package com.java.stuff.springboot;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
class ApplicationTests {

  @Test
  void contextLoads() {
    var args = new String[] {};
    Application.main(args);
    assertThat(Application.class).isNotNull();
  }
}

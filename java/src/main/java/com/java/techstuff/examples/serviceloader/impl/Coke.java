package com.java.techstuff.examples.serviceloader.impl;

import com.java.techstuff.examples.serviceloader.Drink;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class Coke implements Drink {
  public Coke() {
    log.info("creating {}", this);
  }

  @Override
  public String getName() {
    return "Coke";
  }

  @Override
  public int getSize() {
    return 400;
  }
}

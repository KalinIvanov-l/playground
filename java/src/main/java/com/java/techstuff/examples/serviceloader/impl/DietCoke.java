package com.java.techstuff.examples.serviceloader.impl;

import com.java.techstuff.examples.serviceloader.Drink;
import com.java.techstuff.examples.serviceloader.LowCalorie;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@LowCalorie
public class DietCoke implements Drink {
  public DietCoke() {
    log.info("creating {}", this);
  }

  @Override
  public String getName() {
    return "DietCoke";
  }

  @Override
  public int getSize() {
    return 400;
  }
}

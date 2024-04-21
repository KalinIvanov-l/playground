package com.java.techstuff.examples.serviceloader.impl;

import com.java.techstuff.examples.serviceloader.Drink;
import lombok.extern.slf4j.Slf4j;

import java.util.ServiceLoader;

@Slf4j
public class TakeOrder {
  public static void main(String[] args) {
    log.info("We're ready to take your order");
    log.info("What would you like?");

    var drinks = ServiceLoader.load(Drink.class);
    for (var drink : drinks) {
      log.info(drink.getInfo());
    }
    log.info("Please choose from the above");
  }
}

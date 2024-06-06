package com.test.examples;

import static java.lang.StringTemplate.STR;

public class StringTemplates {
  public void stringTemplate() {
    var name = "Kalin";
    var info = STR."My name is \{name}";
    System.out.println(info);
  }

  public static void main() {
    var strTemp = new StringTemplates();
    strTemp.stringTemplate();
  }
}

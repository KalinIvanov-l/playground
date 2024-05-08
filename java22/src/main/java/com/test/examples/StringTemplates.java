package com.test.examples;

public class StringTemplates {
  public void stringTemplate() {
    var name = "Kalin";
    var info = STR."My name is \{name}";
    System.out.println(info);
  }

  public static void main(String[] args) {
    var strTemp = new StringTemplates();
    strTemp.stringTemplate();
  }
}

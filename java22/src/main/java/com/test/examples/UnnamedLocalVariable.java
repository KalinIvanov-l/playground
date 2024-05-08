package com.test.examples;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class UnnamedLocalVariable {
  void checkConnection() {
    try (Connection _ = DriverManager.getConnection("example", "admin", "example")) {
      System.out.println("""
              DB Connection successful
              URL =example
              usr =admin
              pwd =example""");
    } catch (SQLException _) {
      // this is just an example of using `_` IS NOT A GOOD PRACTICE TO DO
    }
  }

  public static void main(String[] args) {
    UnnamedLocalVariable unnamedLocalVariable = new UnnamedLocalVariable();
    unnamedLocalVariable.checkConnection();
  }
}

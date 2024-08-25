package com.test.examples;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class UnnamedLocalVariable {

  /**
   * <pre>
   * @code
   * </pre>
   */
  void checkConnection() {
    try (Connection _ = DriverManager.getConnection("example", "admin", "example")) {
      System.out.println("""
              DB Connection successful
              URL =example
              usr =admin
              pwd =example
              """
      );
    } catch (SQLException _) {
      // this is just an example of using `_`. IT IS NOT A GOOD PRACTICE TO LEAVE THIS BLOCK EMPTY
    }
  }

  void main() {
    UnnamedLocalVariable unnamedLocalVariable = new UnnamedLocalVariable();
    unnamedLocalVariable.checkConnection();
  }
}

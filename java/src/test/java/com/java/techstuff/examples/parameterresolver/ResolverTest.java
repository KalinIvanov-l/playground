package com.java.techstuff.examples.parameterresolver;

import com.java.techstuff.examples.parameterresolverexceptionsolution.Book;
import com.java.techstuff.examples.parameterresolverexceptionsolution.TestParameterResolver;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

import static org.junit.jupiter.api.Assertions.assertEquals;

@ExtendWith(TestParameterResolver.class)
class ResolverTest {
  @Test
  void givenTitleAndAuthor_ShouldAsserExists(String title, String author) {
    Book book = new Book(title, author);
    assertEquals(0, book.getTitle().length());
  }
}

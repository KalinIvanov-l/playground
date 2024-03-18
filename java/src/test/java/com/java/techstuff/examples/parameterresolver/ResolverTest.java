package com.java.techstuff.examples.parameterresolver;

import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

import com.java.techstuff.examples.parameterresolverexceptionsolution.Book;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

@ExtendWith(TestParameterResolver.class)
class ResolverTest {
  @Test
  void givenTitleAndAuthor_ShouldAsserExists(String title, String author) {
    Book book = new Book(title, author);
    assertThat(book.getTitle()).isNotEmpty();
    assertThat(book.getAuthor()).isNotEmpty();
  }

  @Test
  void givenBook_whenConfiguringParameterResolver_thenParameterResolves(Book book) {
    assertThat(book.getTitle()).isEqualTo("Something");
  }
}

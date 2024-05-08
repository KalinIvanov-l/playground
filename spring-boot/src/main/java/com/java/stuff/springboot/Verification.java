package com.java.stuff.springboot;

import org.springframework.web.service.annotation.PostExchange;

public interface Verification {
  @PostExchange("/verify")
  void verify(Application application);
}

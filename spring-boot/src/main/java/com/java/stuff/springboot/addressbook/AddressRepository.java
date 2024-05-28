package com.java.stuff.springboot.addressbook;

import com.java.stuff.springboot.addressbook.model.Address;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AddressRepository extends JpaRepository<Address, Long> {
}

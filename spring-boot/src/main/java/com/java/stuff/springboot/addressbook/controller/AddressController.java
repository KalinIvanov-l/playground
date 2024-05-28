package com.java.stuff.springboot.addressbook.controller;

import com.java.stuff.springboot.addressbook.model.Address;
import com.java.stuff.springboot.addressbook.services.AddressService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/addresses")
public class AddressController {
  private final AddressService addressService;

  @Autowired
  public AddressController(AddressService addressService) {
    this.addressService = addressService;
  }

  @GetMapping
  public List<Address> getAllAddresses() {
    return addressService.getAllAddresses();
  }

  @PostMapping
  public Address createAddress(@RequestBody Address address) {
    return addressService.saveAddress(address);
  }

  @DeleteMapping("/{id}")
  public ResponseEntity<Void> deleteAddress(@PathVariable Long id) {
    addressService.deleteAddress(id);
    return ResponseEntity.noContent().build();
  }
}

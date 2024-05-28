package com.java.stuff.springboot.addressbook.services;

import com.java.stuff.springboot.addressbook.AddressRepository;
import com.java.stuff.springboot.addressbook.model.Address;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AddressService {
  private final AddressRepository addressRepository;

  @Autowired
  public AddressService(AddressRepository addressRepository) {
    validateAddress(addressRepository);
    this.addressRepository = addressRepository;
  }

  public List<Address> getAllAddresses() {
    return addressRepository.findAll();
  }

  public Address saveAddress(Address address) {
    return addressRepository.save(address);
  }

  public void deleteAddress(Long id) {
    addressRepository.deleteById(id);
  }

  private void validateAddress(AddressRepository address) {
    if (address == null) {
      throw new IllegalArgumentException("Address cannot be null");
    }
  }
}
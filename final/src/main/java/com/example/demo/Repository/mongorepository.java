package com.example.demo.Repository;

import com.example.demo.Entity.mongotable;
import org.springframework.data.mongodb.repository.MongoRepository;
public interface mongorepository extends MongoRepository<mongotable,Integer> {

    
}
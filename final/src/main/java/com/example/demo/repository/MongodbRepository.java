package com.example.demo.repository;
import com.example.demo.entity.MongoTable;
import org.springframework.data.mongodb.repository.MongoRepository;
public interface MongodbRepository extends MongoRepository<MongoTable,Integer> {

    
}

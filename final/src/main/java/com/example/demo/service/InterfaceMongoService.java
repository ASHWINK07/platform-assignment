package com.example.demo.service;
import com.example.demo.entity.MongoTable;

public interface InterfaceMongoService {
    MongoTable AddDetails(MongoTable mongo);
    MongoTable GetDetailsbyId(int _id);
    MongoTable UpdateDetailsMongodb(MongoTable mongo);
    String DeleteMongoDetails(int _id);
}

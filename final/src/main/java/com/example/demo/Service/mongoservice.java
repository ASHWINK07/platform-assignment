package com.example.demo.service;

import com.example.demo.entity.MongoTable;
import com.example.demo.repository.MongodbRepository;
import com.example.demo.service.InterfaceMongoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import java.util.Random;
@Service

public class MongoService implements InterfaceMongoService {
    @Autowired
    private MongodbRepository repo;

    public MongoTable AddDetails(MongoTable mongo){
        Random rand = new Random();
        mongo.set_id(rand.nextInt(999));
        return repo.save(mongo);
    }
    public MongoTable GetDetailsbyId(int _id){
        boolean exists = repo.existsById(_id);
        System.out.println(exists);
        if (!exists){
            MongoTable details = new MongoTable();
            details.setName("Details not");
            details.setDepartment("Found");
            return details;
        }
        
        return repo.findById(_id).get();
    }
    public MongoTable UpdateDetailsMongodb(MongoTable mongo){
        boolean exists = repo.existsById(mongo.get_id());
        System.out.println(exists);
        if (!exists){
            MongoTable data = new MongoTable();
            data.setName("Details not");
            data.setDepartment("Found");
            return data;
        }
        MongoTable existingmongodetails = GetDetailsbyId(mongo.get_id());
         //repo.findById(mongo.get_id()).get();
         System.out.println(existingmongodetails.getName());
        existingmongodetails.setDepartment(mongo.getDepartment());
        existingmongodetails.setName(mongo.getName());
        return repo.save(existingmongodetails);
    }
    public String DeleteMongoDetails(int _id){
        boolean exists = repo.existsById(_id);
        System.out.println(exists);
        if (!exists){
            return "details not found";
        }
        repo.deleteById(_id);
        return _id+" deletion successfull";
    }
}

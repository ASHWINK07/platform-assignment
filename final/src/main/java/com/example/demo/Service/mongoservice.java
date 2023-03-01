package com.example.demo.Service;

import com.example.demo.Entity.mongotable;
import com.example.demo.Repository.mongorepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import java.util.Random;
@Service

public class mongoservice {
    @Autowired
    private mongorepository repo;

    public mongotable adddetails(mongotable mongo){
        Random rand = new Random();
        mongo.set_id(rand.nextInt(999));
        return repo.save(mongo);
    }
    public mongotable getDetailsbyid(int _id){
        return repo.findById(_id).get();
    }
    public mongotable updatedetailsMongodb(mongotable mongo){
        mongotable existingmongodetails = getDetailsbyid(mongo.get_id());
         //repo.findById(mongo.get_id()).get();
         System.out.println(existingmongodetails.getName());
        existingmongodetails.setDepartment(mongo.getDepartment());
        existingmongodetails.setName(mongo.getName());
        return repo.save(existingmongodetails);
    }
    public String deletemongodetails(int _id){
        repo.deleteById(_id);
        return _id+" deletion successfull";
    }
}

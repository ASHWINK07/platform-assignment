package com.example.demo.Controller;

import com.example.demo.Entity.sqltable;
import com.example.demo.Entity.mongotable;
import com.example.demo.Service.sqlservice;
import com.example.demo.Service.mongoservice;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
@RestController
public class sqlcontroller {
    @Autowired
    private sqlservice service;

    @Autowired 
    private mongoservice mservice;

    @GetMapping("/testing")
    public String test(@RequestParam String name ){
        return "Testing :"+name;
    }



    // @GetMapping(value = "/records/{id}/")
    // public sqltable getdetails(@RequestParam String db,@PathVariable String id){
    //     //db=mysql
        
    //     if (db=="mongodb")
    //         System.out.println("mongo db work in progress");

    //     return service.getdetailsById(Integer.parseInt(id));
    //     //return db;
    // }
    @GetMapping(value = "/records/{id}/")
    public String getdetails(@RequestParam String db,@PathVariable String id){
        //db=mysql
        
        if (db.equals("mongodb")){
            mongotable result = new mongotable();
            result = mservice.getDetailsbyid(Integer.parseInt(id));
            System.out.println("mongo db work in progress");
            return result.getName()+" "+result.getDepartment();
        }
        
        sqltable p = new sqltable();
        p =  service.getdetailsById(Integer.parseInt(id));
        return p.getName() + " "+ p.getDepartment();
        //return db;
       // return "get successfull";
    }

    @PutMapping(value = "/records/{id}/")
    public String updatedetails(@RequestParam String db,@RequestParam("name") String name,@RequestParam("department") String department,@PathVariable String id){
        //db=mysql
        if (db.equals("mongodb")){
            mongotable data = new mongotable();
            data.set_id(Integer.parseInt(id));
            data.setDepartment(department);
            data.setName(name);
            mservice.updatedetailsMongodb(data);
            System.out.println("mongo db work in progress");
            return "update successfull";
        }
        sqltable Details = new sqltable();
        Details.setId(Integer.parseInt(id));
        Details.setName(name);
        Details.setDepartment(department);
        System.out.println(name+" "+department);
        service.updatedetails(Details);
        return "Update successfull";
    }
    @PostMapping(value = "/records/")
    public String insertdetails(@RequestParam String db,@RequestParam("name") String name,@RequestParam("department") String department){
        //mysql
        // if (db=="mysql"){
        if (db.equals("mongodb")){
            mongotable data = new mongotable();
            data.setDepartment(department);
            data.setName(name);
            mservice.adddetails(data);
            return "Insertion successfull";


        }
        sqltable Details = new sqltable();
        Details.setDepartment(department);
        Details.setName(name);
        service.savedetails(Details);
        System.out.println("working");
        return "Insertion successfull";
        // }
        // return "Insertion successfull";
        // mongodb mong = new mongodb();
        // mong.setDepartment(department);
        // mong.setName(name);
        // //mong.set_id(1);
        // mservice.adddetails(mong);
        // return "Insertion successfull";
        
    }

    @DeleteMapping(value = "/records/{id}/")
    public String deletesqldetails(@RequestParam String db,@PathVariable String id){
        //mysql 
        if (db.equals("mongodb")){
            return mservice.deletemongodetails(Integer.parseInt(id));
        }
        return service.deletedetails(Integer.parseInt(id));
        //return "deletion successfull";
        //return mservice.deletemongodetails(Integer.parseInt(id));
    }
}

package com.example.demo.controller;

import com.example.demo.entity.SqlTable;
import com.example.demo.service.MongoService;
import com.example.demo.service.SqlService;
import com.example.demo.entity.MongoTable;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
@RestController
public class SqlController {
    @Autowired
    private SqlService service;

    @Autowired
    private MongoService mservice;

    @GetMapping("/testing")
    public String Test(@RequestParam String name ){
        return "Testing :"+name;
    }

    @GetMapping(value = "/records/{id}/")
    public String GetDetails(@RequestParam String db,@PathVariable String id){
        //db=mysql
        
        if (db.equals("mongodb")){
            MongoTable result = new MongoTable();
            result = mservice.GetDetailsbyId(Integer.parseInt(id));
            return result.getName()+" "+result.getDepartment();
        }
        //db-mysql
        SqlTable details = new SqlTable();
        details =  service.GetDetailsById(Integer.parseInt(id));
        return details.getName() + " "+ details.getDepartment();
    }

    @PutMapping(value = "/records/{id}/")
    public String UpdateDetails(@RequestParam String db,@RequestParam("name") String name,@RequestParam("department") String department,@PathVariable String id){
        //db=mysql
        if (db.equals("mongodb")){
            MongoTable new_data = new MongoTable();
            new_data.set_id(Integer.parseInt(id));
            new_data.setDepartment(department);
            new_data.setName(name);
            new_data = mservice.UpdateDetailsMongodb(new_data);
            if (new_data.getDepartment().equals("Found") ) {
                return "Details not present in database";
            }
            return "update successfull";
        }
        SqlTable details = new SqlTable();
        details.setId(Integer.parseInt(id));
        details.setName(name);
        details.setDepartment(department);
        //System.out.println(name+" "+department);
        details = service.UpdateDetails(details);
        if (details.getDepartment().equals("Found")) {
            return "Details not present in database";
        }
        return "Update successfull";
    }
    @PostMapping(value = "/records/")
    public String InsertDetails(@RequestParam String db,@RequestParam("name") String name,@RequestParam("department") String department){

        if (db.equals("mongodb")){
            MongoTable data = new MongoTable();
            data.setDepartment(department);
            data.setName(name);
            mservice.AddDetails(data);
            
            return "Insertion successfull";


        }
        SqlTable details = new SqlTable();
        details.setDepartment(department);
        details.setName(name);
        service.SaveDetails(details);
        
        return "Insertion successfull";

        
    }

    @DeleteMapping(value = "/records/{id}/")
    public String DeleteDetails(@RequestParam String db,@PathVariable String id){
        
        if (db.equals("mongodb")){
            return mservice.DeleteMongoDetails(Integer.parseInt(id));
        }
        return service.DeleteDetails(Integer.parseInt(id));
    }
}

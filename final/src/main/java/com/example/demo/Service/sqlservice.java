package com.example.demo.service;
import com.example.demo.entity.SqlTable;
import com.example.demo.repository.SqlRepository;
import com.example.demo.service.InterfaceSqlService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
@Service
public class SqlService implements InterfaceSqlService {
    @Autowired
    private SqlRepository repository;

    public SqlTable SaveDetails(SqlTable details) {
        return repository.save(details);
    }
    public SqlTable GetDetailsById(int id) {
        boolean exists = repository.existsById(id);
        System.out.println(exists);
        if (!exists){
            SqlTable details = new SqlTable();
            details.setName("Details not");
            details.setDepartment("Found");
            return details;
        }
        System.out.println(repository.findAll());
        return repository.findById(id).orElse(null);
    }

    // public sqltable getdetailsByName(String name) {
    //     return repository.findByName(name);
    // }

    public String DeleteDetails(int id) {
        boolean exists = repository.existsById(id);
        System.out.println(exists);
        if (!exists){
            return "details not found";
        }
        repository.deleteById(id);
        return "details removed !! " + id;
    }

    public SqlTable UpdateDetails(SqlTable details) {
        boolean exists = repository.existsById(details.getId());
        System.out.println(exists);
        if (!exists){
            SqlTable data = new SqlTable();
            data.setName("Details not");
            data.setDepartment("Found");
            return data;
        }
        SqlTable existingdetails = repository.findById(details.getId()).orElse(null);
        existingdetails.setName(details.getName());
        existingdetails.setDepartment(details.getDepartment());
        return repository.save(existingdetails);
    }
}

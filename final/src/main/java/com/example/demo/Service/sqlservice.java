package com.example.demo.Service;
import com.example.demo.Entity.sqltable;
import com.example.demo.Repository.sqlrepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
@Service
public class sqlservice {
    @Autowired
    private sqlrepository repository;

    public sqltable savedetails(sqltable details) {
        return repository.save(details);
    }
    public sqltable getdetailsById(int id) {
        return repository.findById(id).orElse(null);
    }

    // public sqltable getdetailsByName(String name) {
    //     return repository.findByName(name);
    // }

    public String deletedetails(int id) {
        repository.deleteById(id);
        return "details removed !! " + id;
    }

    public sqltable updatedetails(sqltable details) {
        sqltable existingdetails = repository.findById(details.getId()).orElse(null);
        existingdetails.setName(details.getName());
        existingdetails.setDepartment(details.getDepartment());
        // existingdetails.setQuantity(details.getQuantity());
        // existingdetails.setPrice(details.getPrice());
        return repository.save(existingdetails);
    }
}

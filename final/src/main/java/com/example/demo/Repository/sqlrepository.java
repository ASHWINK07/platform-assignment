package com.example.demo.Repository;
import com.example.demo.Entity.sqltable;
import org.springframework.data.jpa.repository.JpaRepository;
public interface sqlrepository extends JpaRepository<sqltable,Integer> {
    
}

package com.example.demo.repository;
import com.example.demo.entity.SqlTable;
import org.springframework.data.jpa.repository.JpaRepository;
public interface SqlRepository extends JpaRepository<SqlTable,Integer> {
    
}

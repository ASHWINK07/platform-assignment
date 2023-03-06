package com.example.demo.entity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import jakarta.persistence.Id;
import jakarta.persistence.Table;
//import jakarta.persistence. GeneratedValue;
import jakarta.persistence.Entity;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "users")
public class SqlTable {
    
    @Id
    //@GeneratedValue
    private int id;  
    private String Name;
    private String Department;

    
    
}

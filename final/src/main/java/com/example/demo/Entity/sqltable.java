package com.example.demo.Entity;
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
public class sqltable {
    @Id
    //@GeneratedValue
    private int id;
    private String name;
    private String department;
}

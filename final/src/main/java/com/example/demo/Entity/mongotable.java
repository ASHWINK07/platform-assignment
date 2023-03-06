package com.example.demo.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
//import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
@Document(collection = "records")
@Data
@AllArgsConstructor
@NoArgsConstructor
public class MongoTable {
    //@Id
    private int _id;
    private String Name;
    private String Department;
}

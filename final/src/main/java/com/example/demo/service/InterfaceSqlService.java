package com.example.demo.service;
import com.example.demo.entity.SqlTable;

public interface InterfaceSqlService {
     SqlTable SaveDetails(SqlTable details);
     SqlTable GetDetailsById(int id);
     String DeleteDetails(int id);
     SqlTable UpdateDetails(SqlTable details) ;
}

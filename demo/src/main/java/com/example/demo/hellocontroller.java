package com.example.demo;

import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RequestMapping;
//import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.stereotype.Controller;
//import org.springframework.web.bind.annotation.RequestMapping;
@CrossOrigin(origins="*")
@Controller
public class hellocontroller {
    @ResponseBody
    @RequestMapping("/")
    public String hello(){
        return "Hello World!";
    }
}
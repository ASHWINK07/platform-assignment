package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestParam;
//import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
@RestController
@SpringBootApplication
public class DemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}
	@ResponseBody
    @GetMapping(value = "/mysql")
    public String getdetails(@RequestParam String name){
		//String name = 
		System.out.println("Get request successfull ");
        return name;
    }
	@ResponseBody
	@DeleteMapping(value = "/mysql")
	public String deletedetails(@RequestParam String name){
		System.out.println("Delete request successfull");
		return name;
	}
	@ResponseBody
	@PostMapping(value = "/mysql")
	public String insertdetails(@RequestParam("name") String name,@RequestParam("department") String department){
		System.out.println("Post Request Successfull");
		return "name: "+name+"&department:"+department;
	}
	@ResponseBody
	@PutMapping(value = "/mysql")
	public String updatedetails(@RequestParam("name") String name,@RequestParam("department") String department){
		System.out.println("Put Request Successfull");
		return "name: "+name+"&department:"+department;
	}



}

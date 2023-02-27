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
    public String getdetailssql(@RequestParam String name){
		//String name = 
		System.out.println("Get request successfull for mysql ");
        return name;
    }
	@ResponseBody
	@DeleteMapping(value = "/mysql")
	public String deletedetailssql(@RequestParam String name){
		System.out.println("Delete request successfull for mysql ");
		return name;
	}
	@ResponseBody
	@PostMapping(value = "/mysql")
	public String insertdetailssql(@RequestParam("name") String name,@RequestParam("department") String department){
		System.out.println("Post Request Successfull for mysql ");
		return "name: "+name+"&department:"+department;
	}
	@ResponseBody
	@PutMapping(value = "/mysql")
	public String updatedetailssql(@RequestParam("name") String name,@RequestParam("department") String department){
		System.out.println("Put Request Successfull for mysql ");
		return "name: "+name+"&department:"+department;
	}

	@ResponseBody
	@GetMapping(value = "/mongodb")
	public String getdetailsmongodb(@RequestParam String name){
		System.out.println("Get request successfull for mongodb ");
        return name;	
	}

	@ResponseBody
	@DeleteMapping(value="/mongodb")
	public String deletedetailsmongodb(@RequestParam String name){
		System.out.println("Delete Request successfull for mongodb");
		return name;
	}

	@ResponseBody
	@PostMapping(value = "/mongodb")
	public String insertdetailsmongodb(@RequestParam("name") String name,@RequestParam("department") String department){
		System.out.println("Post Request sucessfull for mongodb");
		return "name: "+name+"&department:"+department;
	}

	@ResponseBody
	@PutMapping(value = "/mongodb")
	public String updatedetailsmongodb(@RequestParam("name") String name,@RequestParam("department") String department){
		System.out.println("Put Request sucessfull for mongodb");
		return "name: "+name+"&department:"+department;
	}

}

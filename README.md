# platform-assignment
CRUD operation on mongdb and mysql using Golang and Java and loadbalancer
java spring boot runs in port number 8080 
Go httpserver runs in port number 8081
these are the apis 
check for the id path whether that users exists or not
Go
	-mysql
		curl http://127.0.0.1:8081/records/1/\?db\=mysql
		curl -X POST http://127.0.0.1:8081/records/\?db\=mysql -d "name=archit&department=frontend"
		curl -X POST http://127.0.0.1:8081/records/\?db\=mysql -d "name=ashwin&department=platform"
		curl -X POST http://127.0.0.1:8081/records/\?db\=mysql -d "name=gourav&department=AI"
		curl -X PUT http://127.0.0.1:8081/records/6/\?db\=mysql -d "name=ashwin&department=frontend"
		curl -X DELETE http://127.0.0.1:8081/records/1/\?db\=mysql

	- mongodb
		curl http://127.0.0.1:8081/records/3/\?db\=mongodb 
		curl -X POST http://127.0.0.1:8081/records/\?db\=mongodb -d "name=gourav&department=platform"
		curl -X POST http://127.0.0.1:8081/records/\?db\=mongodb -d "name=ashwin&department=platform"
		curl -X POST http://127.0.0.1:8081/records/\?db\=mongodb -d "name=prajwal&department=AI"
		curl -X PUT http://127.0.0.1:8081/records/31/\?db\=mongodb -d "name=gourav&department=frontend"
		curl -X DELETE http://127.0.0.1:8081/records/31/\?db\=mongodb

Java
	-mysql
		curl http://127.0.0.1:8080/records/2/\?db\=mysql
		curl -X POST http://127.0.0.1:8080/records/\?db\=mysql -d "name=archit&department=frontend"
		curl -X POST http://127.0.0.1:8080/records/\?db\=mysql -d "name=yashwal&department=self-servePIM"
		curl -X PUT http://127.0.0.1:8080/records/2/\?db\=mysql -d "name=ashwin&department=frontend"
		curl -X DELETE http://127.0.0.1:8080/records/2/\?db\=mysql
	- MongoDB
		curl http://127.0.0.1:8080/records/148/\?db\=mongodb
		curl -X POST http://127.0.0.1:8080/records/\?db\=mongodb -d "name=yashwal&department=self-serve-PIM"
		curl -X PUT http://127.0.0.1:8080/records/379/\?db\=mongodb -d "name=ashwin&department=frontend"
		curl -X DELETE http://127.0.0.1:8080/records/148/\?db\=mongodb
    
  to run the application first run the mongodb and mysql server . Here I am running both of these databases inside docker . I have exposed the port 
  command:
  inside final folder run the following command
  - mvn spring-boot:run
  inside go folder run the following command
  - go run httpserver.go
  at any terminal run the below command
  -docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 --name MYSQL mysql:latest
  -docker run -p 27017:27017 --name MONGODB mongo:latest
  
  run the HAproxy 
  sudo haproxy -f loadbalancer.cfg
  
  

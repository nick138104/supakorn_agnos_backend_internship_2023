# supakorn_agnos_backend_internship_2023

## Deploy on local

#### file .env

สร้าง .env file ที่มี CONNECTION_STRING ที่เอาไว้ต่อ postgresql
```
CONNECTION_STRING = "user=<username> password=<password> dbname=<database_name> sslmode=disable"
```
โดยใส่ username, password และ database_name ของ postgresql ที่จะใช้ในการเก็บ log ผลลัพท์

#### table in database

ทำการสร้าง table ที่มีชื่อ log ใน database
``` sql
CREATE TABLE log (
	init_password VARCHAR ( 40 ),
	num_of_steps int
);
```
init_password จะทำการเก็บ password ที่ส่งมาให้ตรวจสอบ และ num_of_steps เก็บผลลัพท์จำนวน action ที่ต้องทำ

## Unit Test

สำหรับการ run ตัว unit test สามารถใช้คำสั่ง
```go test```
จะทำการ run unit test ของ file password_validation_test.go ที่ทำการทดสอบ function password_validation ใน server.go
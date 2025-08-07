package config

import (
	"flow1/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CoonectDB(){
	dsn := "root:passwordb@tcp(127.0.0.1:3306)/flowdb?charset=utf8mb4&parseTime=True&loc=Local"
	database,err :=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("failed o connecto db",err)
	}
 fmt.Println("db connected")

 database.AutoMigrate(&models.User{})
 DB  = database
}
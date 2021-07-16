package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitMYSQL() {

	DB, err := gorm.Open(mysql.Open(os.Getenv("DSN")))
	if err != nil {
		log.Fatal(err.Error())
	}
	db = DB
	if err:=DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&TestDrive{});err !=nil{
		log.Fatal(err.Error())
	}
}
package database

import (
	"ChatApp/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Dns = "host=localhost user=postgres password=3242 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(postgres.Open(Dns), &gorm.Config{})
	if err != nil {
		panic("veritabanına bağlanamadın gevşek")
	}

	DB = database
	err = database.AutoMigrate(&models.User{}, models.Aboutme{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

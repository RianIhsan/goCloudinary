package database

import (
	"fmt"
	"log"

	"github.com/RianIhsan/goCloudinary/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/gocloudinary?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Koneksi database gagal!")
	} else {
		fmt.Println("Koneksi database berhasil!")
	}
}

func RunMigrate() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal migrasi table")
	} else {
		fmt.Println("Berhasil migrasi table")
	}
}

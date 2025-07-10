package database

import (
	"badminton-app/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/badminton?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	err = DB.AutoMigrate(
		&models.Kehadiran{},
		&models.Keuangan{},
		&models.User{},
	)
	if err != nil {
		log.Fatal("Gagal migrasi database: ", err)
	}

	fmt.Println("Koneksi database berhasil")
}

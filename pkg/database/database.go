package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "u317027005_golang:Adyfp24.@tcp(153.92.15.18)/u317027005_db_ngl?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("koneksi terhubung")
	DB = db
	
	return db, nil
}

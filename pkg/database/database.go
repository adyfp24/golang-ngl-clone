package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "root:OBtrbunNRFDXdqDmdeAfcQrZyqEyUNRc@tcp(roundhouse.proxy.rlwy.net:30579)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("koneksi terhubung")
	DB = db
	
	return db, nil
}

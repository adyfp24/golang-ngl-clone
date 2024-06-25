package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := viper.GetString("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("koneksi terhubung")
	DB = db
	
	return db, nil
}

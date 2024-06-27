package migrations

import (
	"fmt"
	"log"

	"github.com/adyfp24/golang-ngl-clone/app/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB){
	err := db.AutoMigrate(&models.User{}, &models.Chat{})
	if(err != nil){
		log.Print(err)
	}

	fmt.Println("migrasi berhasil dijalankan")
}
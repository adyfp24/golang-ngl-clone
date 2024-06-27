package migrations

import (
	"fmt"
	"log"

	"github.com/adyfp24/golang-ngl-clone/app/models"
	"github.com/adyfp24/golang-ngl-clone/pkg/database"
)

func RunMigration(){
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Chat{})
	if(err != nil){
		log.Print(err)
	}

	fmt.Println("migrasi berhasil dijalankan")
}
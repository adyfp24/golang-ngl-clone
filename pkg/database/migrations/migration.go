package migrations

import(
	"fmt"
	"github.com/adyfp24/gin-ngl-clone/pkg/database"
	"log"
)

func RunMigration(){
	db, err := database.InitDB()
	err = db.AutoMigrate()
	if(err != nil){
		log.Print(err)
	}

	fmt.Println("migrasi berhasil dijalankan")
}
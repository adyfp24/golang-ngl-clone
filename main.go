package main

import (
	"github.com/adyfp24/gin-ngl-clone/database/migrations"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	migrations.RunMigration()

	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "welcome to ngl clone")
	})
	
	err := app.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

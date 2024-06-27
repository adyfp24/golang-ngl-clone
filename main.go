package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adyfp24/golang-ngl-clone/app/routes"
	"github.com/adyfp24/golang-ngl-clone/config"
	"github.com/adyfp24/golang-ngl-clone/pkg/database"
	"github.com/adyfp24/golang-ngl-clone/pkg/database/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	db, err := database.InitDB()
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}
	migrations.RunMigration(db)
	app := gin.Default()

	app.StaticFS("/static", http.Dir("./web/static"))
	app.LoadHTMLGlob("./app/views/*html")
	
	app.Use(cors.Default())
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "welcome to ngl clone")
	})

	routes.InitRoute(app)

	err = app.Run(":" + viper.GetString("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}

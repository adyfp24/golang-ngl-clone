package main

import (
	"github.com/adyfp24/gin-ngl-clone/pkg/database/migrations"
	"github.com/adyfp24/gin-ngl-clone/internal/routes"
	"github.com/adyfp24/gin-ngl-clone/config"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/gin-contrib/cors"
	"net/http"
)

func main() {
	config.LoadConfig()
	migrations.RunMigration()
	app := gin.Default()

	app.StaticFS("/static", http.Dir("./web/static"))
	app.LoadHTMLGlob("./internal/views/*")
	
	app.Use(cors.Default())
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "welcome to ngl clone")
	})

	routes.InitRoute(app)

	err := app.Run(":" + viper.GetString("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}

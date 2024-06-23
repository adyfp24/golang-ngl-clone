package main

import (
	"github.com/adyfp24/gin-ngl-clone/database/migrations"
	"github.com/adyfp24/gin-ngl-clone/routes"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/gin-contrib/cors"
	"net/http"
)

func main() {
	migrations.RunMigration()
	app := gin.Default()

	app.StaticFS("/static", http.Dir("./web/static"))
	app.LoadHTMLGlob("./web/static/templates/pages/*")
	app.Use(cors.Default())
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "welcome to ngl clone")
	})

	routes.InitRoute(app)

	err := app.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

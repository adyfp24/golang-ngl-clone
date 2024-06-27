package handler

import (
	// "github.com/adyfp24/golang-ngl-clone/pkg/database/migrations"
	"github.com/adyfp24/golang-ngl-clone/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
)

var app *gin.Engine

func init() {
	// migrations.RunMigration()
	app = gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.StaticFS("/static", http.Dir("../web/static"))
	app.LoadHTMLGlob("../app/views/*")
	app.Use(cors.Default())

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "welcome to ngl clone")
	})

	routes.InitRoute(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

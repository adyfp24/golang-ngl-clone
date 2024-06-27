package handler

import (
	"html/template"
	"net/http"

	"github.com/adyfp24/golang-ngl-clone/app/routes"
	"github.com/adyfp24/golang-ngl-clone/app/views"
	"github.com/adyfp24/golang-ngl-clone/pkg/database/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	migrations.RunMigration()
    app = gin.New()
    app.Use(gin.Logger())
    app.Use(gin.Recovery())
    // app.StaticFS("/static", http.Dir("web/static"))
	
	_template, err := template.ParseFS(
		views.Default,
		"default/*.html",
	)
	if err != nil {
		panic(err)
	}
	app.SetHTMLTemplate(_template)

    app.Use(cors.Default())
    app.GET("/", func(ctx *gin.Context) {
        ctx.String(200, "welcome to ngl clone")
    })

    routes.InitRoute(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
    app.ServeHTTP(w, r)
}

package handler

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/adyfp24/golang-ngl-clone/app/routes"
	"github.com/adyfp24/golang-ngl-clone/app/views"
	"github.com/adyfp24/golang-ngl-clone/pkg/database"
	"github.com/adyfp24/golang-ngl-clone/web"

	// "github.com/adyfp24/golang-ngl-clone/pkg/database"
	// "github.com/adyfp24/golang-ngl-clone/pkg/database/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	database.InitDB()

    app = gin.New()
    app.Use(gin.Logger())
    app.Use(gin.Recovery())
	defaultFads, _ := fs.Sub(web.Static, "static")
	app.StaticFS("/static", http.FS(defaultFads))
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

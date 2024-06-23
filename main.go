package main

import(
	"log"
	"github.com/gin-gonic/gin"
)

func main(){
	app := gin.Default();
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "welcome to ngl clone");
	})
	err := app.Run(":3000");
	if(err != nil){
		log.Fatal(err);
	}
}
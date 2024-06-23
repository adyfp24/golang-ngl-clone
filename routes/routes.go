package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/adyfp24/gin-ngl-clone/handlers"
)

func InitRoute(route *gin.Engine){

	route.GET("/question", handlers.CreateQuestion)
	

}
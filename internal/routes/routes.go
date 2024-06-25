package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/adyfp24/gin-ngl-clone/internal/handlers"
)

func InitRoute(route *gin.Engine){

	route.GET("/question", handlers.CreateQuestion)
	

}
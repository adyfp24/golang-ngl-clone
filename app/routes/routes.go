package routes

import(
	"github.com/gin-gonic/gin"
	"github.com/adyfp24/golang-ngl-clone/app/handlers"
)

func InitRoute(route *gin.Engine){

	route.GET("/question", handlers.QuestionRender)
	route.POST("/create-question", handlers.CreateQuestion)

	route.GET("/question-all", handlers.ReadAllQuestion)
	route.GET("/question/:id", handlers.ReadQuestionById)

	route.POST("/answer/:id", handlers.CreateAnswer)
}
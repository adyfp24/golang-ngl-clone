package handlers

import(
	"github.com/gin-gonic/gin"
)

func AnswerRender(c *gin.Context){
	c.HTML(200, "answer.html", nil)
}

func CreateAnswer(c *gin.Context){

}
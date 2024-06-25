package handlers

import(
	"github.com/gin-gonic/gin"
) 

func QuestionRender(c *gin.Context){
	c.HTML(200, "create-question.html", nil)	
}
func CreateQuestion(c *gin.Context){
		
}

func ReadAllQuestion(c *gin.Context){
	c.HTML(200, "all-question.html", nil)
}

func ReadQuestionById(c *gin.Context){
	c.HTML(200, "question-detail.html", nil)
}

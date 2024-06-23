package handlers

import(
	"github.com/gin-gonic/gin"
) 

func CreateQuestion(c *gin.Context){
	c.HTML(200, "question.html", gin.H{
		"title":   "Question Page",
        "content": "Welcome to the question page!",
	})	
}
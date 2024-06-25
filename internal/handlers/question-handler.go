package handlers

import (
	"github.com/adyfp24/gin-ngl-clone/internal/models"
	"github.com/adyfp24/gin-ngl-clone/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QuestionRender(c *gin.Context) {
	c.HTML(200, "create-question.html", nil)
}

func CreateQuestion(c *gin.Context) {
	var input models.InputMessage

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := models.Chat{
		Question: input.Question,
		UserID: 1,
	}

	if err := repositories.CreateQuestion(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pertanyaan berhasil dibuat"})
}

func ReadAllQuestion(c *gin.Context) {
    questions, err := repositories.ReadAllQuestion(1) 
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.HTML(http.StatusOK, "all-question.html", gin.H{
        "questions": questions,
    })
}

func ReadQuestionById(c *gin.Context) {
	c.HTML(200, "question-detail.html", nil)
}

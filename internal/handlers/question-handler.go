package handlers

import (
	"net/http"
	"strconv"

	"github.com/adyfp24/gin-ngl-clone/internal/models"
	"github.com/adyfp24/gin-ngl-clone/internal/repositories"
	"github.com/gin-gonic/gin"
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
	messsageID := c.Param("id")
	parseID, err := strconv.ParseInt(messsageID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
        return	
	}
	question, err := repositories.ReadQuestionById(uint(parseID))
	
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if question == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
        return
    }

	c.HTML(200, "question-detail.html", gin.H{
		"question": question,
	})
}

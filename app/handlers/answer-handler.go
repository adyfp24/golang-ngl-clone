package handlers

import (
	"net/http"
	"strconv"

	"github.com/adyfp24/golang-ngl-clone/app/models"
	"github.com/adyfp24/golang-ngl-clone/app/repositories"
	"github.com/gin-gonic/gin"
)

func CreateAnswer(c *gin.Context) {
	var input models.InputAnswer

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid message ID"})
		return
	}

	err = repositories.CreateAnswer(uint(messageID), input.Answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"answer" : "answer created succesfull",
	})
}

func ShareAnswer(c * gin.Context) {
	messageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
	}

	question, err := repositories.ShareAnswer(uint(messageID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
	}

	if question == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : "pesan tidak ditemukan",
		})
	}

	c.HTML(200, "share-answer.html", gin.H{
		"question" : question,
	})
}

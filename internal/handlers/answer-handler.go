package handlers

import (
	"net/http"
	"strconv"

	"github.com/adyfp24/gin-ngl-clone/internal/models"
	"github.com/adyfp24/gin-ngl-clone/internal/repositories"
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

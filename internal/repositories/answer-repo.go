package repositories

import (
	"github.com/adyfp24/gin-ngl-clone/internal/models"
	"github.com/adyfp24/gin-ngl-clone/pkg/database"
)

func CreateAnswer(id uint, answer string) error {
	var chat models.Chat
	err := database.DB.First(&chat, id).Error
	if err != nil {
		return err
	}
	chat.Answer = answer
	return database.DB.Save(&chat).Error
}
package repositories

import (
	"github.com/adyfp24/gin-ngl-clone/internal/models"
	"github.com/adyfp24/gin-ngl-clone/pkg/database"
	"gorm.io/gorm"
)

func CreateQuestion(message *models.Chat) error {
	return database.DB.Create(message).Error
}

func ReadAllQuestion(userID uint) ([]models.Chat, error) {
	var chats []models.Chat
	if err := database.DB.Where("user_id = ?", userID).Find(&chats).Error; err != nil {
		return nil, err
	}
	return chats, nil
}

func ReadQuestionById(messageID uint) (*models.Chat, error) {
	var chat models.Chat
	err := database.DB.First(&chat, messageID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chat, err
}

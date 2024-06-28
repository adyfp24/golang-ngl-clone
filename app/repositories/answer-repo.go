package repositories

import (
	"github.com/adyfp24/golang-ngl-clone/app/models"
	"github.com/adyfp24/golang-ngl-clone/pkg/database"
	"gorm.io/gorm"
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

func ShareAnswer(id uint) (*models.Chat, error) {
	var chat models.Chat
	err := database.DB.First(&chat, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chat, nil
}
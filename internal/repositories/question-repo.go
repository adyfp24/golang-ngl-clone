package repositories

import(
	"github.com/adyfp24/gin-ngl-clone/pkg/database"
	"github.com/adyfp24/gin-ngl-clone/internal/models"
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
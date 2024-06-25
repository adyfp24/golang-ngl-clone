package repositories

import(
	"github.com/adyfp24/gin-ngl-clone/pkg/database"
	"github.com/adyfp24/gin-ngl-clone/internal/models"
)

func CreateChat(chat *models.Chat) error {
    return database.DB.Create(chat).Error
}
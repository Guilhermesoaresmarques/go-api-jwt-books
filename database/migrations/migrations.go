package migrations

import (
	"github.com/guilhermesoaresmarques/go-api-jwt-books/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}

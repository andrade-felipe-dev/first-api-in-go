package migrations

import (
	"github.com/andrade-felipe-dev/first-api-in-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}

package db

// todo продумать необходимость миграций
import (
	"app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Team{}, &models.PullRequest{}); err != nil {
		return nil, err
	}

	return db, nil
}

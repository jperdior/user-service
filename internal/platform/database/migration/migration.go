package migration

import (
	"gorm.io/gorm"
	"user-service/internal/user/domain"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
	)
}

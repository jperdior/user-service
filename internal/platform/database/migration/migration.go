package migration

import (
	"gorm.io/gorm"
	"user-service/internal/user/infrastructure/persistence"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&persistence.UserModel{},
	)
}

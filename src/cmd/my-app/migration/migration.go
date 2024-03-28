package migration

import (
	"fmt"

	"main.go/src/internal/constant"
	conf "main.go/src/internal/db"
	"main.go/src/internal/model"
)

// Migration defines the migration function
func Migration() error {
	if !constant.AUTO_MIGRATE {
		return nil // Skip migration if auto-migration is disabled
	}

	// Migrate schemas to database
	models := []interface{}{
		&[]model.User{
			{Username: "admin", Password: "admin"},
		},
		// Add other models for migration here (if applicable)
		&model.Account{},
		&model.Transaction{},
		&model.Entries{},
	}

	if err := conf.DB.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error auto-migrating models: %w", err)
	}

	for _, model := range models {
		if err := conf.DB.Create(model).Error; err != nil {
			return fmt.Errorf("error seeding models: %w", err)
		}
	}

	return nil
}

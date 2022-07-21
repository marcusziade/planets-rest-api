package database

import (
	"github.com/jinzhu/gorm"
	"github.com/marcusziade/planets-rest-api/planet"
)

// Migrates the database and creates the planet postgres table
func MigrateDatabase(database *gorm.DB) error {
	if result := database.AutoMigrate(&planet.Planet{}); result.Error != nil {
		return result.Error
	}

	return nil
}

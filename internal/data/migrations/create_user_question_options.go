package migrations

import (
	"spirit_quiz/internal/data/models"

	"gorm.io/gorm"
)

func CreatTables(db *gorm.DB) error {
	// Initialize the migrations
	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Question{}, &models.Option{}, &models.Result{})
	if err != nil {
		return err
	}
	return nil
}

func CreatResultTables(db *gorm.DB) error {
	// Initialize the migrations
	err := db.AutoMigrate(&models.Result{})
	if err != nil {
		return err
	}
	return nil
}

func CreatUserTables(db *gorm.DB) error {
	// Initialize the migrations
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

func DropTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Category{}, &models.Question{}, &models.Option{})
	if err != nil {
		return err
	}
	return nil
}

func DropResultTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Result{})
	if err != nil {
		return err
	}
	return nil
}

func DropUserTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

package database

import (
	"spirit_quiz/config"

	// "spirit_quiz/internal/data/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	config, err := config.GetConfigs()
	if err != nil {
		return nil, err
	}
	// Set up the database URL
	databaseUrl := config.DbUrl

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = migrations.CreatTables(db)
	// if err != nil {
	// 	return nil, err
	// }

	// err = migrations.DropTables(db)
	// if err != nil {
	// 	return nil, err
	// }

	// err = seed.Seeder(db)
	// if err != nil {
	// 	return nil, err
	// }

	// Return the database object
	return db, nil
}

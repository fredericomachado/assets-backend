package database

import (
	"assets-backend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"assets-backend/config"
)

var (
	DB  *gorm.DB
	err error
)

func MigrateDatabase() {

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.Env.DbHost,
		config.Env.DbUser,
		config.Env.DbPassword,
		config.Env.DbDatabase)
	DB, err = gorm.Open(postgres.Open(connectionString))

	DB.AutoMigrate(
		&models.Allocation{},
		&models.Asset{},
		&models.Category{},
		&models.Location{},
		&models.TargetAllocation{},
		&models.TargetMacroAllocation{},
		&models.User{},
		&models.Wallet{},
	)
}

func Connect() {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.Env.DbHost,
		config.Env.DbUser,
		config.Env.DbPassword,
		config.Env.DbDatabase)
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Database error: ", err)
	}
}

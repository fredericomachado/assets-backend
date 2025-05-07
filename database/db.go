package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"assets-backend/config"
)

type postgresDB struct {
	Db *sql.DB
}

var (
	DB  *gorm.DB
	err error
)

func NewPosgresDB() (*postgresDB, error) {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Env.DbUser, config.Env.DbPassword, config.Env.DbServer, config.Env.DbPort, config.Env.DbDatabase)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	return &postgresDB{Db: db}, nil
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

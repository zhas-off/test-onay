package database

import (
	"fmt"
	"os"

	"github.com/zhas-off/test-onay/internal/user"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {
	log.Info("Setting up new database connection")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DB"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to database: %w", err)
	}

	log.Println("running migrations")
	db.AutoMigrate(&user.User{})

	return &Database{
		Client: db,
	}, nil
}

package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhas-off/test-onay/internal/database"
	transportHTTP "github.com/zhas-off/test-onay/internal/transport/http"
	"github.com/zhas-off/test-onay/internal/user"
)

// Run - sets up our application
func Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting Up Our APP")

	var err error

	err = godotenv.Load("../.env")

	if err != nil {
		log.Error("Error loading .env file")
		return err
	}

	store, err := database.NewDatabase() //connecting to database
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}

	userService := user.NewService(store)
	handler := transportHTTP.NewHandler(userService)

	if err := handler.Serve(); err != nil {
		log.Error("failed to gracefully serve our application")
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}
}

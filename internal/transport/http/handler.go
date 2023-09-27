package http

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	App     *fiber.App
	Service UserService
}

func NewHandler(service UserService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.App = fiber.New()

	// Set up the routes
	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	// Создаем группу маршрутов "/api"
	api := h.App.Group("/api")

	// Применяем middleware к группе маршрутов
	api.Use(CheckAgeMiddleware)

	// Определяем маршруты внутри группы
	api.Post("/user", h.PostUser)
	api.Get("/users", h.GetUsers)
	api.Put("/user/:id", h.UpdateUser)
}

// Serve  serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.App.Listen(":8080"); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := h.App.ShutdownWithContext(ctx); err != nil {
		log.Error(err)
	}

	log.Println("shutting down gracefully")
	return nil
}

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

	// Sets up our middleware functions
	// h.App.Use(CheckAgeMiddleware)

	// Set up the routes
	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	h.App.Post("/api/v1/user", h.PostUser)
	h.App.Get("/api/user", h.GetUsers)
	h.App.Put("/api/user/:id", h.UpdateUser)
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
		panic(err)
	}

	log.Println("shutting down gracefully")
	return nil
}

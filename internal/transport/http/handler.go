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
	Service CommentService
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}

	h.App = fiber.New()

	// Sets up our middleware functions
	h.App.Use(CheckAgeMiddleware)

	// Set up the routes
	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	h.App.Post("/api/v1/comment", h.PostComment)
	h.App.Get("/api/v1/comment/:id", h.GetComments)
	h.App.Put("/api/v1/comment/:id", h.UpdateComment)
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

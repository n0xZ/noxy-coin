package server

import (
	"noxy/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	*fiber.App
	client database.Service
}

func New() *FiberServer {
	middlewares := cors.New()

	server := &FiberServer{
		App:    fiber.New(),
		client: database.New(),
	}
	server.Use(middlewares)
	return server
}

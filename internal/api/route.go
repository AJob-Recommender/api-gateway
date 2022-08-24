package api

import (
	"github.com/AJob-Recommender/base-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func route(fib *fiber.App, handler *handler.Handler) {
	fib.Post("/predict", handler.Predict)
}

package api

import (
	"github.com/AJob-Recommender/base-api/internal/config"
	"github.com/AJob-Recommender/base-api/internal/handler"
	"github.com/AJob-Recommender/base-api/internal/services/seer"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Serve(cfg *config.Config, log *zap.SugaredLogger) {
	fib := fiber.New()

	seerClient := seer.NewSeer(cfg, log)

	h := handler.NewHandler(cfg, log, seerClient)

	route(fib, h)
}

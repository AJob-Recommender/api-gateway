package api

import (
	"github.com/AJob-Recommender/base-api/internal/config"
	"github.com/AJob-Recommender/base-api/internal/handler"
	"github.com/AJob-Recommender/base-api/internal/services/seer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func Serve(cfg *config.Config, log *zap.SugaredLogger) {
	fib := fiber.New()

	corsConfig := cors.ConfigDefault
	corsConfig.AllowHeaders = "*"
	corsConfig.AllowCredentials = true
	fib.Use(cors.New(corsConfig))

	seerClient := seer.NewSeer(cfg, log)
	h := handler.NewHandler(cfg, log, seerClient)

	route(fib, h)

	if err := fib.Listen(cfg.API.Port); err != nil {
		log.Error(err.Error())
	}
}

package handler

import (
	"github.com/AJob-Recommender/base-api/internal/config"
	"github.com/AJob-Recommender/base-api/pkg/client"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	Config     *config.Config
	Log        *zap.SugaredLogger
	SeerClient client.Client
}

func NewHandler(cfg *config.Config, log *zap.SugaredLogger, seerClient client.Client) *Handler {
	return &Handler{
		Config:     cfg,
		Log:        log,
		SeerClient: seerClient,
	}
}

func (h *Handler) Predict(ctx *fiber.Ctx) error {

	return nil
}

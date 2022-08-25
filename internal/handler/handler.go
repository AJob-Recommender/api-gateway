package handler

import (
	"github.com/AJob-Recommender/base-api/internal/config"
	"github.com/AJob-Recommender/base-api/internal/services/seer"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Handler struct {
	Config     *config.Config
	Log        *zap.SugaredLogger
	SeerClient *seer.Seer
}

func NewHandler(cfg *config.Config, log *zap.SugaredLogger, seerClient *seer.Seer) *Handler {
	return &Handler{
		Config:     cfg,
		Log:        log,
		SeerClient: seerClient,
	}
}

func (h *Handler) Predict(ctx *fiber.Ctx) error {
	req := new(seer.Request)

	if err := ctx.BodyParser(req); err != nil {
		h.Log.Warn("request body is not valid", zap.Error(err), zap.Any("request", string(ctx.Body())))

		return errors.New("request is not valid")
	}

	res, err := h.SeerClient.Predict(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

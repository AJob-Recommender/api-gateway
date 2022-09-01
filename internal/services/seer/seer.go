package seer

import (
	"context"
	"github.com/AJob-Recommender/base-api/internal/config"
	"github.com/AJob-Recommender/base-api/pkg/client"
	"go.uber.org/zap"
	"math"
	"net/http"
)

const (
	endpoint = "/predict"
)

type Seer struct {
	Client client.Client
	Config *config.Config
	Log    *zap.SugaredLogger
}

type Service interface {
	Predict(ctx context.Context, req *Request) (res interface{}, err error)
}

func NewSeer(config *config.Config, log *zap.SugaredLogger) *Seer {
	seerCfg := config.Services.Seer
	seerClient := client.NewHTTP(client.Options{
		Timeout:            seerCfg.Timeout,
		MaxConsPerHost:     seerCfg.MaxConsPerHost,
		MaxIdleConsPerHost: seerCfg.MaxIdleConsPerHost,
		MaxIdleCons:        seerCfg.MaxIdleCons,
		MaxRetry:           seerCfg.MaxRetry,
		Logger:             log,
	})

	return &Seer{
		Client: seerClient,
		Config: config,
		Log:    log,
	}
}

func (s Seer) Predict(ctx context.Context, req *Request) (res Response, err error) {
	request := client.Request{
		Method: http.MethodPost,
		URL:    s.Config.Services.Seer.URL + endpoint,
		Body:   req,
	}

	if _, err = s.Client.DoWithRetry(ctx, request, &res); err != nil {
		return res, err
	}

	for idx, val := range res.Results {
		res.Results[idx].Confidence = roundNumber(val.Confidence)
	}

	return res, nil
}

func roundNumber(number float64) float64 {
	return math.Floor(number*100) / 100
}

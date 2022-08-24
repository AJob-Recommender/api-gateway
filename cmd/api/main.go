package main

import (
	"github.com/AJob-Recommender/base-api/internal/api"
	"github.com/AJob-Recommender/base-api/internal/config"
	"github.com/AJob-Recommender/base-api/pkg/logger"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		panic(err.Error())
	}

	log, err := logger.New()

	api.Serve(cfg, log)
}

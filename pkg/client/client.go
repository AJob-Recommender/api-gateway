package client

import (
	"context"
	"github.com/AJob-Recommender/base-api/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// Client represent client.
type Client interface {
	Do(req *http.Request) (res *http.Response, err error)
	DoWithRetry(ctx context.Context, request Request, responseData interface{}) (internal Response, err error)
}

type HTTP struct {
	Options Options
}

type Request struct {
	Body   interface{}
	URL    string
	Method string
}

type Options struct {
	Timeout            time.Duration
	MaxConsPerHost     int
	MaxIdleConsPerHost int
	MaxIdleCons        int
	MaxRetry           int
	Client             *http.Client
	Logger             logger.Logger
	LoggerEncoder      string
	LoggerOutputType   string
	LoggerLevel        string
	AuthUser           string
	AuthPass           string
}

// NewHTTP create new http client
// if logger is nil and creating it failed, this function will panic.
func NewHTTP(options Options) *HTTP {
	if options.Client == nil {
		options.Client = &http.Client{
			Timeout: options.Timeout,
			Transport: &http.Transport{
				MaxConnsPerHost:     options.MaxConsPerHost,
				MaxIdleConns:        options.MaxIdleCons,
				MaxIdleConnsPerHost: options.MaxIdleConsPerHost,
			},
		}
	}

	if options.Logger == nil {
		var err error
		options.Logger, err = logger.New(options.LoggerLevel, options.LoggerOutputType, options.LoggerEncoder)

		if err != nil {
			panic(err)
		}
	}

	return &HTTP{Options: options}
}

// Do call http client Do method.
func (h *HTTP) Do(req *http.Request) (res *http.Response, err error) {

	res, err = h.Options.Client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/AJob-Recommender/base-api/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
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
	Logger             *zap.SugaredLogger
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
		options.Logger, err = logger.New()

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

func (h *HTTP) DoWithRetry(ctx context.Context, request Request, responseData interface{}, // nolint:cyclop
) (internal Response, err error) {
	buff := &bytes.Buffer{}
	if err = json.NewEncoder(buff).Encode(request.Body); err != nil {
		h.Options.Logger.Warn("Encode Request Error", zap.Error(err), zap.Any("request", request))

		return internal, errors.WithStack(err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL, buff)
	if err != nil {
		h.Options.Logger.Warn("Init Request Error", zap.Error(err))

		return internal, errors.WithStack(err)
	}

	httpRequest.SetBasicAuth(h.Options.AuthUser, h.Options.AuthPass)
	httpRequest.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	for try := 0; try < h.Options.MaxRetry; try++ {
		resp, err = h.DoWithContextCancellation(ctx, httpRequest) // nolint:bodyclose // we are closing body in above code.
		if err == nil || errors.Is(err, context.Canceled) {
			break
		}

		if try == h.Options.MaxRetry-1 {
		}
	}

	if resp != nil {
		bodyByte, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			return internal, errors.WithStack(err)
		}

		internal = Response{Body: bodyByte, StatusCode: resp.StatusCode}
	}

	if err != nil {
		return internal, err
	}

	defer CloseIO(h.Options.Logger.Warn, resp.Body)

	if responseData != nil {
		if err = json.Unmarshal(internal.Body, &responseData); err != nil {
			h.Options.Logger.Warn("Response Parse Error", zap.Error(err), zap.ByteString("response", internal.Body))

			return internal, errors.WithStack(err)
		}
	}

	if responseData == nil {
		responseData = string(internal.Body)
	}

	return internal, nil
}

// DoWithContextCancellation is like Do with handling context cancellation (normally happened by user).
func (h *HTTP) DoWithContextCancellation(ctx context.Context, req *http.Request) (res *http.Response, err error) {
	var status int
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // nolint:wrapcheck
	default:
		defer func() {
			if err != nil {
				h.Options.Logger.Warn("Request Error",
					zap.String("url", req.URL.String()),
					zap.Int("max_retry", h.Options.MaxRetry),
					zap.Int("status", status),
					zap.Error(err),
				)
			}
		}()
		res, err = h.Do(req)

		if err != nil {
			return nil, err
		}

		if res != nil && !IsSuccessful(res.StatusCode) {
			status = res.StatusCode

			return res, errors.Errorf("http request failed with status of: %s", res.Status)
		}

		return res, nil
	}
}

// IsSuccessful indicated with 2xx status codes.
func IsSuccessful(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

// CloseIO close io closer.
func CloseIO(log func(args ...interface{}), closer io.ReadCloser) {
	if closer == nil {
		return
	}

	// to avoid memory leak when reusing http connection
	if _, err := io.Copy(ioutil.Discard, closer); err != nil {
		log(err)
	}

	if err := closer.Close(); err != nil {
		log(err)
	}
}

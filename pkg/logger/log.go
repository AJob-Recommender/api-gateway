package logger

import (
	"go.uber.org/zap"
)

// New function.
func New() (sugar *zap.SugaredLogger, err error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	defer func() {
		errs := logger.Sync()

		if errs != nil {
			logger.Error(errs.Error())

			return
		}
	}()

	sugar = logger.Sugar()

	return
}

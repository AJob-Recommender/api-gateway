package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// nolint
const (
	JSONEncoder    = "json"
	ConsoleEncoder = "console"

	STDOutOutput = "stdout"
	STDErrOutput = "stderr"
	NatsOutput   = "nats"

	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	DPanic     = "dpanic"
	Panic      = "panic"
	Fatal      = "fatal"
)

// New create a new zap Sugared Logger.
// outputType can be multi like : stdout|nats.
func New(level string, outputType string, encoderType string) (*zap.SugaredLogger, error) {
	encoder, err := createEncoder(encoderType)
	if err != nil {
		return nil, err
	}

	outputs, err := createOutputs(outputType)
	if err != nil {
		return nil, err
	}

	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cores := make([]zapcore.Core, 0, len(outputs))
	for _, writer := range outputs {
		cores = append(cores, zapcore.NewCore(encoder, writer, zapLevel))
	}

	return zap.New(zapcore.NewTee(cores...)).Sugar(), nil
}

func createEncoder(encoderType string) (zapcore.Encoder, error) { // nolint:ireturn
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	switch encoderType {
	case ConsoleEncoder:
		return zapcore.NewConsoleEncoder(encoderConfig), nil
	case JSONEncoder:
		return zapcore.NewJSONEncoder(encoderConfig), nil
	}

	return nil, errors.New("unsupported encoder type")
}

func createOutputs(outputType string) ([]zapcore.WriteSyncer, error) {
	parts := strings.Split(outputType, "|")
	writers := make([]zapcore.WriteSyncer, 0, len(parts))

	for _, part := range parts {
		switch part {
		case STDOutOutput:
			writers = append(writers, zapcore.AddSync(os.Stdout))
		case STDErrOutput:
			writers = append(writers, zapcore.AddSync(os.Stderr))
		case NatsOutput:
			panic("not implemented")
		default:
			return nil, errors.New(fmt.Sprintf("unsupported output type: %s", part))
		}
	}

	return writers, nil
}

// Logger is acceptable logger.
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	DPanicf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync() error
}

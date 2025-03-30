package logs

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(2))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	if IsTestEnvironment() {
		return
	}

	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	if IsTestEnvironment() {
		return
	}

	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	if IsTestEnvironment() {
		return
	}

	log.Error(message, fields...)
}

func IsTestEnvironment() bool {
	return os.Getenv("GO_TESTING_MODE") == "true"
}

func Println(a ...any) (n int, err error) {
	if IsTestEnvironment() {
		return
	}

	return fmt.Println(a...)
}

func Printf(format string, a ...any) (n int, err error) {
	if IsTestEnvironment() {
		return
	}

	return fmt.Printf(format, a...)
}

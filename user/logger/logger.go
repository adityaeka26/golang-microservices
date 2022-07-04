package logger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

type Logger interface {
	GetLogger() *zap.Logger
}

type LoggerImpl struct {
	logger *zap.Logger
}

func NewLog(appName string, appEnvironment string) Logger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger = logger.With(zap.String("app", appName)).With(zap.String("environment", appEnvironment))
	// count := 0
	// for {
	// 	if rand.Float32() > 0.8 {
	// 		logger.Error("oops...something is wrong",
	// 			zap.Int("count", count),
	// 			zap.Error(errors.New("error details")))
	// 	} else {
	// 		logger.Info("everything is fine",
	// 			zap.Int("count", count))
	// 	}
	// 	count++
	// 	time.Sleep(time.Second * 2)
	// }
	return &LoggerImpl{
		logger: logger,
	}
}

func (logger LoggerImpl) GetLogger() *zap.Logger {
	return logger.logger
}

package logger

import (
	"os"

	"go.uber.org/zap"
)

func ZapLog() *zap.Logger {

	_, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return nil
	}
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"./log.log"}
	cfg.ErrorOutputPaths = []string{"./log.log"}

	logger, _ := cfg.Build()
	return logger
}

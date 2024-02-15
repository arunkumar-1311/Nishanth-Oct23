package logger

import (
	"fmt"
	"os"
	"github.com/go-kit/log"
)

func GokitLogger(msg error) log.Logger {
	var logger log.Logger

	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Print("Can't create file ", err)
		return nil
	}
	logger = log.NewJSONLogger(file)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"message", msg,
		"time", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	return logger
}



package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

func ZeroLogger() (log *zerolog.Event) {

	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Print("Can't create file ", err)
		return
	}
	logger := zerolog.New(file).With().Caller().Timestamp().Logger()
	log = logger.Error()
	return
}

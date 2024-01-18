package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Create a log file to record the logs
func logFile() (file *os.File) {

	var Logerr error
	file, Logerr = os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if Logerr != nil {
		fmt.Print("Can't create file ", Logerr)
		return nil
	}
	return file
}

func Logging() (log *logrus.Logger) {
	file := logFile()
	log = &logrus.Logger{
		Out:          file,
		ReportCaller: true,
		Formatter:    new(logrus.JSONFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.DebugLevel,
	}
	return log
}

package applog

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var loggers = make(map[string]zerolog.Logger)

func Init(names ...string) {
	for _, name := range names {
		openFile, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("open log file err: %v\n", err)
			return
		}
		loggers[name] = zerolog.New(openFile)
	}
}
func InitV2(names ...string) {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	for _, name := range names {
		logFile, err := os.OpenFile(fmt.Sprintf("./logs/%s.log", name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Printf("open log file err: %v\n", err)
			return
		}
		multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
		logger := zerolog.New(multi).With().Timestamp().Logger()
		loggers[name] = logger
	}
}

func Logger(name string) *zerolog.Logger {
	instance := loggers[name]
	return &instance
}

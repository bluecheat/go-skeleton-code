package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"skeleton-code/config"
	"skeleton-code/logger"
)

type closer func()

func loggerInit(conf *config.Config) closer {
	var logOutput io.Writer
	var output closer = func() {}
	if conf.Log.Type == "file" {
		f, err := os.OpenFile("broker.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(fmt.Sprintf("error opening log file: %v", err))
			os.Exit(1)
		}
		logOutput = f
		output = func() {
			f.Close()
		}
	} else {
		logOutput = os.Stdout
	}
	level, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		panic(fmt.Sprintf("error log level : %v", err))
	}
	logger.Init(logOutput, level)
	return output
}

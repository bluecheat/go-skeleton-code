package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

var logger *logrus.Logger

func Root() *logrus.Logger {
	return logger
}

func Init(output io.Writer, level logrus.Level) *logrus.Logger {
	logger = logrus.New()
	logger.SetFormatter(&LogFormatter{})
	logger.AddHook(&CodeLineNumberHook{})
	logger.SetOutput(output)

	logger.SetLevel(level)
	return logger
}

func Http(c *gin.Context) logrus.FieldLogger {
	request, is := c.Get(LFRequestID)
	if !is {
		return logger
	}
	return logger.WithField(LFRequestID, request)
}

func Info(args ...interface{}) {
	logger.Infoln(args)
}

func Warn(args ...interface{}) {
	logger.Warnln(args)
}

func Error(args ...interface{}) {
	logger.Errorln(args)
}

func Debug(args ...interface{}) {
	logger.Debugln(args)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

// 2016-09-27 09:38:21.541541811 +0200 CEST
// 127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700]
// "GET /apache_pb.gif HTTP/1.0" 200 2326
// "http://www.example.com/start.html"
// "Mozilla/4.08 [en] (Win98; I ;Nav)"

var timeFormat = "2006-01-02 15:04:05"

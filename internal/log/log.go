package log

import "github.com/sirupsen/logrus"

// NewLogger Создать логгер
func NewLogger() *logrus.Logger {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "01.02.2006 15:04:05"

	logger := logrus.New()
	logger.SetFormatter(customFormatter)

	return logger
}

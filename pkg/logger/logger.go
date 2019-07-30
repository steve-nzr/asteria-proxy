package logger

import (
	"github.com/gookit/color"
)

var globalPrefix = "[KIT] "

// Debug log
func Debug(format string, args ...interface{}) {
	color.Cyan.Printf(prefixFormat(format), args...)
}

// Info log
func Info(format string, args ...interface{}) {
	color.Info.Printf(prefixFormat(format), args...)
}

// Warning log
func Warning(format string, args ...interface{}) {
	color.Warn.Printf(prefixFormat(format), args...)
}

// Error log
func Error(format string, args ...interface{}) {
	color.Error.Printf(prefixFormat(format), args...)
}

func prefixFormat(format string) string {
	return globalPrefix + " " + format + "\n"
}

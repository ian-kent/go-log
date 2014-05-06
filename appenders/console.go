package appenders

import (
	"fmt"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/levels"
)

type consoleAppender struct {
	Appender
	Layout layout.Layout
}

func Console() *consoleAppender {
	return &consoleAppender{
		Layout: layout.Default(),
	}
}

func (a *consoleAppender) Write(level levels.LogLevel, message string, args ...interface{}) {
	fmt.Println(a.Layout.Format(level, message, args...))
}

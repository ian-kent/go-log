package log

import (
	"fmt"
	"github.com/ian-kent/go-log/levels"
	"time"
)

type Logger struct {
	level   levels.LogLevel
	name    string
	Enabled map[levels.LogLevel]bool
}

var logger *Logger

func Global() *Logger {
	if logger == nil {
		logger = New(levels.DEBUG, ".")
	}
	return logger
}

func New(level levels.LogLevel, name string) *Logger {
	logger := &Logger{
		level:   level,
		name:    name,
		Enabled: make(map[levels.LogLevel]bool),
	}
	logger.SetLevel(level)
	return logger
}

func compose(level levels.LogLevel, args ...interface{}) (string, []interface{}) {
	msg := "[%s] [%s] " + args[0].(string) + "\n"
	args = args[1:]
	return msg, append([]interface{}{time.Now(), levels.LogLevels[level]}, args...)
}

func write(level levels.LogLevel, params ...interface{}) {
	msg, args := compose(level, params)
	fmt.Printf(msg, args...)
}

func unwrap(args ...interface{}) []interface{} {
	head := args[0]
	switch head.(type) {
	case func() (string, []interface{}):
		msg, args := head.(func() (string, []interface{}))()
		args = unwrap(args...)
		return append([]interface{}{msg}, args...)
	case func() []interface{}:
		args = unwrap(head.(func() []interface{})()...)
	case func(...interface{}) []interface{}:
		args = unwrap(head.(func(...interface{}) []interface{})(args[1:]...)...)
	}
	return args
}

func Log(level levels.LogLevel, params ...interface{}) {
	Global().Log(level, params...)
}

func Debug(params ...interface{})   { Log(levels.DEBUG, params...) }
func Info(params ...interface{})    { Log(levels.INFO, params...) }
func Warn(params ...interface{})    { Log(levels.WARN, params...) }
func Error(params ...interface{})   { Log(levels.ERROR, params...) }
func Trace(params ...interface{})   { Log(levels.TRACE, params...) }
func Printf(params ...interface{})  { Log(levels.INFO, params...) }
func Println(params ...interface{}) { Log(levels.INFO, params...) }
func Fatalf(params ...interface{})  { Log(levels.FATAL, params...) }

func (l *Logger) write(level levels.LogLevel, params ...interface{}) {
	write(level, params...)
}
func (l *Logger) Log(level levels.LogLevel, params ...interface{}) {
	if !l.Enabled[level] {
		return
	}
	l.write(level, unwrap(params...)...)
}
func (l *Logger) Level() levels.LogLevel {
	return l.level
}
func (l *Logger) Name() string {
	return l.name
}
func (l *Logger) SetLevel(level levels.LogLevel) {
	l.level = level
	for k, _ := range levels.LogLevels {
		if k <= level {
			l.Enabled[k] = true
		} else {
			l.Enabled[k] = false
		}
	}
}

func (l *Logger) Debug(params ...interface{})   { l.Log(levels.DEBUG, params...) }
func (l *Logger) Info(params ...interface{})    { l.Log(levels.INFO, params...) }
func (l *Logger) Warn(params ...interface{})    { l.Log(levels.WARN, params...) }
func (l *Logger) Error(params ...interface{})   { l.Log(levels.ERROR, params...) }
func (l *Logger) Trace(params ...interface{})   { l.Log(levels.TRACE, params...) }
func (l *Logger) Printf(params ...interface{})  { l.Log(levels.INFO, params...) }
func (l *Logger) Println(params ...interface{}) { l.Log(levels.INFO, params...) }
func (l *Logger) Fatalf(params ...interface{})  { l.Log(levels.FATAL, params...) }

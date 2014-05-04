package log

import (
	"fmt"
	"time"
)

type Logger struct {
	level   LogLevel
	Enabled map[LogLevel]bool
}

type LogLevel int

const (
	FATAL = iota
	ERROR
	INFO
	WARN
	DEBUG
	TRACE
)

var levels = map[LogLevel]string{
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	WARN:  "WARN",
	INFO:  "INFO",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

var logger *Logger

func Global() *Logger {
	if logger == nil {
		logger = New(DEBUG)
	}
	return logger
}

func New(level LogLevel) *Logger {
	logger := &Logger{
		level:   level,
		Enabled: make(map[LogLevel]bool),
	}
	logger.SetLevel(level)
	return logger
}

func compose(level LogLevel, args ...interface{}) (string, []interface{}) {
	msg := "[%s] [%s] " + args[0].(string) + "\n"
	args = args[1:]
	return msg, append([]interface{}{time.Now(), levels[level]}, args...)
}

func write(level LogLevel, params ...interface{}) {
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

func Log(level LogLevel, params ...interface{}) {
	Global().Log(level, params...)
}

func Debug(params ...interface{})   { Log(DEBUG, params...) }
func Info(params ...interface{})    { Log(INFO, params...) }
func Warn(params ...interface{})    { Log(WARN, params...) }
func Error(params ...interface{})   { Log(ERROR, params...) }
func Trace(params ...interface{})   { Log(TRACE, params...) }
func Printf(params ...interface{})  { Log(INFO, params...) }
func Println(params ...interface{}) { Log(INFO, params...) }
func Fatalf(params ...interface{})  { Log(FATAL, params...) }

func (l *Logger) write(level LogLevel, params ...interface{}) {
	write(level, params...)
}
func (l *Logger) Log(level LogLevel, params ...interface{}) {
	if !l.Enabled[level] {
		return
	}
	l.write(level, unwrap(params...)...)
}
func (l *Logger) Level() LogLevel {
	return l.level
}
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
	for k, _ := range levels {
		if k <= level {
			l.Enabled[k] = true
		} else {
			l.Enabled[k] = false
		}
	}
}

func (l *Logger) Debug(params ...interface{})   { l.Log(DEBUG, params...) }
func (l *Logger) Info(params ...interface{})    { l.Log(INFO, params...) }
func (l *Logger) Warn(params ...interface{})    { l.Log(WARN, params...) }
func (l *Logger) Error(params ...interface{})   { l.Log(ERROR, params...) }
func (l *Logger) Trace(params ...interface{})   { l.Log(TRACE, params...) }
func (l *Logger) Printf(params ...interface{})  { l.Log(INFO, params...) }
func (l *Logger) Println(params ...interface{}) { l.Log(INFO, params...) }
func (l *Logger) Fatalf(params ...interface{})  { l.Log(FATAL, params...) }

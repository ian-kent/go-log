package log

import(
	"fmt"
	"time"
)

type Logger struct {
	Level LogLevel
}

type LogLevel int
const (
	TRACE = iota
	DEBUG
	WARN
	INFO
	ERROR
	FATAL
)
var levels = map[LogLevel]string{
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	WARN: "WARN",
	INFO: "INFO",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

var logger *Logger

func New(level LogLevel) *Logger {
	return &Logger{
		Level: level,
	}
}

func Write(level LogLevel, message string, params ...interface{}) {
	p := append([]interface{}{time.Now(), levels[level]}, params...)
	fmt.Printf("[%s] [%s] " + message + "\n", p...)
}

func Unwrap(args ...interface{}) []interface{} {
	head := args[0]
	switch head.(type) {
	case func(...interface{})[]interface{}:
		args = head.(func(...interface{})[]interface{})(args[1:]...)
	}
	return args
}

func Log(level LogLevel, params ...interface{}) {
	if logger == nil { logger = New(DEBUG) }
	logger.Log(level, params...)
}
func Debug(params ...interface{}) {	Log(DEBUG, params...) }
func Info(params ...interface{}) { Log(INFO, params...) }
func Warn(params ...interface{}) { Log(WARN, params...) }
func Error(params ...interface{}) {	Log(ERROR, params...) }
func Trace(params ...interface{}) {	Log(TRACE, params...) }
func Printf(params ...interface{}) { Log(INFO, params...) }
func Println(params ...interface{}) { Log(INFO, params...) }
func Fatalf(params ...interface{}) { Log(FATAL, params...) }

func (l *Logger) Write(level LogLevel, message string, params ...interface{}) { 
	Write(level, message, params...) 
}
func (l *Logger) Log(level LogLevel, params ...interface{}) { 
	l.Write(level, params[0].(string), params[1:]...) 
}
func (l *Logger) Debug(params ...interface{}) { l.Log(DEBUG, Unwrap(params...)...) }
func (l *Logger) Info(params ...interface{}) { l.Log(INFO, Unwrap(params...)...) }
func (l *Logger) Warn(params ...interface{}) { l.Log(WARN, Unwrap(params...)...) }
func (l *Logger) Error(params ...interface{}) { l.Log(ERROR, Unwrap(params...)...) }
func (l *Logger) Trace(params ...interface{}) { l.Log(TRACE, Unwrap(params...)...) }
func (l *Logger) Printf(params ...interface{}) { l.Log(INFO, Unwrap(params...)...) }
func (l *Logger) Println(params ...interface{}) { l.Log(INFO, Unwrap(params...)...) }
func (l *Logger) Fatalf(params ...interface{}) { l.Log(FATAL, Unwrap(params...)...) }
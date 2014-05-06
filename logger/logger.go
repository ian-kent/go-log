package logger

import(
	"github.com/ian-kent/go-log/levels"
	"github.com/ian-kent/go-log/appenders"
	"strings"
)

type Logger interface {
	Level()            levels.LogLevel
	Name()             string
	FullName()         string
	Enabled()          map[levels.LogLevel]bool
	Appender()         Appender
	Children()         []Logger
	Parent()           Logger
	GetLogger(string)  Logger
	SetLevel(levels.LogLevel)
	Log(levels.LogLevel, ...interface{})
}

type logger struct {
	Logger
	level    levels.LogLevel
	name     string
	enabled  map[levels.LogLevel]bool
	appender Appender
	children []Logger
	parent   Logger
}

type Appender interface {
	Write(level levels.LogLevel, message string, args ...interface{})
}

func New(name string) Logger {
	l := Logger(&logger{
		level: levels.DEBUG,
		name: name,
		enabled: make(map[levels.LogLevel]bool),
		appender: appenders.Console(),
		children: make([]Logger, 0),
		parent: nil,
	})
	l.SetLevel(levels.DEBUG)
	return l
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

func (l *logger) GetLogger(name string) Logger {
	bits := strings.Split(name, ".")

	if l.name == bits[0] {
		if len(bits) == 1 {
			return l
		}
		
		child := bits[1]
		for _, c := range l.children {
			if c.Name() == child {
				return c.GetLogger(strings.Join(bits[1:], "."))
			}
		}

		lg := New(child)
		lg.(*logger).parent = l
		l.children = append(l.children, lg)
		return lg.GetLogger(strings.Join(bits[1:], "."))
	}
	lg := New(bits[0])
	lg.(*logger).parent = l
	return lg.GetLogger(name)
}

func (l *logger) write(level levels.LogLevel, params ...interface{}) {
	l.Appender().Write(level, params[0].(string), params[1:]...)
}

func (l *logger) Log(level levels.LogLevel, params ...interface{}) {
	if !l.enabled[level] {
		return
	}
	l.write(level, unwrap(params...)...)
}

func (l *logger) Level() levels.LogLevel {
	return l.level
}

func (l *logger) Enabled() map[levels.LogLevel]bool {
	return l.enabled
}

func (l *logger) Name() string {
	return l.name
}

func (l *logger) FullName() string {
	n := l.name
	if l.parent != nil {
		p := l.parent.FullName()
		if len(p) > 0 {
			n = l.parent.FullName() + "." + n
		}
	}
	return n
}

func (l *logger) SetLevel(level levels.LogLevel) {
	l.level = level
	for k, _ := range levels.LogLevelsToString {
		if k <= level {
			l.enabled[k] = true
		} else {
			l.enabled[k] = false
		}
	}
}

func (l *logger) Debug(params ...interface{})   { l.Log(levels.DEBUG, params...) }
func (l *logger) Info(params ...interface{})    { l.Log(levels.INFO, params...) }
func (l *logger) Warn(params ...interface{})    { l.Log(levels.WARN, params...) }
func (l *logger) Error(params ...interface{})   { l.Log(levels.ERROR, params...) }
func (l *logger) Trace(params ...interface{})   { l.Log(levels.TRACE, params...) }
func (l *logger) Printf(params ...interface{})  { l.Log(levels.INFO, params...) }
func (l *logger) Println(params ...interface{}) { l.Log(levels.INFO, params...) }
func (l *logger) Fatalf(params ...interface{})  { l.Log(levels.FATAL, params...) }

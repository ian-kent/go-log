Go-Log
======

[![Build Status](https://travis-ci.org/ian-kent/go-log.svg?branch=master)](https://travis-ci.org/ian-kent/go-log)

A logger, for Go!

It's sort of ```log``` and ```code.google.com/p/log4go``` compatible, so in most cases
can be used without any code changes.

### Getting started

Install go-log:

```
go get github.com/ian-kent/go-log/log
```

Use the logger in your application:

```
import(
  "github.com/ian-kent/go-log/log"
)

// Pass a log message and arguments directly
log.Debug("Example log message: %s", "example arg")

// Pass a function which returns a log message and arguments
log.Debug(func(){[]interface{}{"Example log message: %s", "example arg"}})
log.Debug(func(i ...interface{}){[]interface{}{"Example log message: %s", "example arg"}})
```

You can also get the logger instance:
```
logger := log.Logger()
logger.Debug("Yey!")
```

Or get a named logger instance:

```
logger := log.Logger("foo.bar")
```

### Log levels

The default log level is DEBUG.

To get the current log level:

```
level := logger.Level()
```

Or to set the log level:

```
// From a LogLevel
logger.SetLevel(levels.TRACE)

// From a string
logger.SetLevel(log.Stol("TRACE"))
```

### Log appenders

The default log appender is ```appenders.Console()```, which logs
the raw message to STDOUT.

To get the current log appender:

```
appender := logger.Appender()
```

If the appender is ```nil```, the parent loggers appender will be used
instead.

If the appender eventually resolves to ```nil```, log data will be
silently dropped.

You can set the log appender:

```
logger.SetAppender(appenders.Console())
```

#### Rolling file appender

Similar to log4j's rolling file appender, you can use

```
// Append to (or create) file
logger.SetAppender(appenders.RollingFile("filename.log", true))

// Truncate (or create) file
logger.SetAppender(appenders.RollingFile("filename.log", false))
```

You can also control the number of log files which are kept:
```
r := appenders.RollingFile("filename.log", true)
r.MaxBackupIndex = 2 // filename.log, filename.log.1, filename.log.2
```

And the maximum log file size (in bytes):
```
r := appenders.RollingFile("filename.log", true)
r.MaxFileSize = 1024 // 1KB, defaults to 100MB
```

#### Fluentd appender

The fluentd appender lets you write log data directly to fluentd:

```
logger.SetAppender(appenders.Fluentd(fluent.Config{}))
```

It uses ```github.com/t-k/fluent-logger-golang```.

The tag is currently fixed to 'go-log', and the data structure sent
to fluentd is simple:

```
{
  message: "<output from layout>"
}

```

### Layouts

Each appender has its own layout. This allows the log data to be transformed
as it is written to the appender.

The default layout is ```layout.Basic()```, which passes the log message
and its arguments through ```fmt.Sprintf```.

To get the current log appender layout:
```
appender := logger.Appender()
layout := appender.Layout()
```

To set the log appender layout:
```
appender.SetLayout(layout.Basic())
```

You can also use ```layout.Pattern(pattern string)```, which accepts a
pattern format similar to log4j:

| Code | Description
| ---- | -----------
| %c   | The package the log statement is in
| %C   | Currently also the package the log statement is in
| %d   | The current date/time, using ```time.Now().String()```
| %F   | The filename the log statement is in
| %l   | The location of the log statement, e.g. ```package/somefile.go:12```
| %L   | The line number the log statement is on
| %m   | The log message and its arguments formatted with ```fmt.Sprintf```
| %n   | A new-line character
| %p   | Priority - the log level
| %r   | ms since logger was created

### Logger inheritance

Loggers are namespaced with a ```.```, following similar rules to Log4j.

If you create a logger named ```foo```, it will automatically inherit the
log settings (levels and appender) of the root logger.

If you then create a logger named ```foo.bar```, it will inherit the log
settings of ```foo```, which in turn inherits the log settings from the
root logger.

You can break this by setting the log level or setting an appender on
a child logger, e.g.:

```
logger := log.Logger("foo.bar")
logger.SetLevel(levels.TRACE)
logger.SetAppender(appenders.Console())
```

If you then created a logger named ```foo.bar.qux```, it would inherit
the trace level and console appender of the ```foo.bar``` logger.

### Roadmap

* log4j configuration support
  * .properties
  * .xml
  * .json
* layouts
  * fixmes/todos in pattern layout
* appenders
  * add socket appender
  * fixmes/todos and tests for fluentd appender
* optimise logger creation
  * collapse loggers when parent namespace is unused
  * reorganise loggers when new child tree is created
* add godoc documentation

### Contributing

Before submitting a pull request:

  * Format your code: ```go fmt ./...```
  * Make sure tests pass: ```go test ./...```

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.

Go-Log
======

[![Build Status](https://travis-ci.org/ian-kent/go-log.svg?branch=master)](https://travis-ci.org/ian-kent/go-log)

A logger, for Go!

### Getting started

Install go-log:

```go get github.com/ian-kent/go-log/log```

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

### Contributing

Before submitting a pull request:

  * Format your code: ```go fmt ./...```
  * Make sure tests pass: ```go test ./...```

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.

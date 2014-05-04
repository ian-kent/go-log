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

You can set the log level:

```
// From a LogLevel
logger.SetLevel(levels.TRACE)

// From a string
logger.SetLevel(log.Stol("TRACE"))
```

### Contributing

Before submitting a pull request:

  * Format your code: ```go fmt ./...```
  * Make sure tests pass: ```go test ./...```

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.

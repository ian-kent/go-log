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
logger := log.Global()
logger.Debug("Yey!")
```

You can also create a local logger object:

```
logger := log.New(log.Level("DEBUG"))
```

### Log levels

The default log level is DEBUG.

To get the current log level:

```
level := logger.Level()
```

You can set the log level:

```
logger.SetLevel(log.Level("TRACE"))
```

### Contributing

Before submitting a pull request:

  * Format your code: ```go fmt ./...```
  * Make sure tests pass: ```go test ./...```

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.

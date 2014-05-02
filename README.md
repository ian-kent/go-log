Go-Log
======

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
log.Debug(func(i ...interface{}){[]interface{}{"Example log message: %s", "example arg"}})
```

The default log level is DEBUG. You can't (currently) change the default log
level on the global logger.

You can also create a local logger object:

```
logger := log.New(log.DEBUG)
```

You can then set the log level:

```
logger.Level = log.TRACE
```

### Contributing

Before submitting a pull request:

  * Format your code: ```go fmt ./...```
  * Make sure tests pass: ```go test ./...```

### Licence

Copyright ©‎ 2014, Ian Kent (http://www.iankent.eu).

Released under MIT license, see [LICENSE](LICENSE.md) for details.

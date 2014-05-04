package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New(Level("DEBUG"), ".")
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), Level("DEBUG"))
	assert.NotNil(t, logger.Name())
	assert.Equal(t, logger.Name(), ".")
}

func TestGlobal(t *testing.T) {
	logger := Global()
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), Level("DEBUG"))
}

func TestLevels(t *testing.T) {
	logger := Global()
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), Level("DEBUG"))
	assert.Equal(t, logger.Enabled[Level("TRACE")], false)
	assert.Equal(t, logger.Enabled[Level("DEBUG")], true)
	assert.Equal(t, logger.Enabled[Level("WARN")], true)
	assert.Equal(t, logger.Enabled[Level("ERROR")], true)
	assert.Equal(t, logger.Enabled[Level("INFO")], true)
	assert.Equal(t, logger.Enabled[Level("FATAL")], true)

	logger.SetLevel(Level("TRACE"))
	assert.Equal(t, logger.Level(), Level("TRACE"))
	assert.Equal(t, logger.Enabled[Level("TRACE")], true)
	assert.Equal(t, logger.Enabled[Level("DEBUG")], true)
	assert.Equal(t, logger.Enabled[Level("WARN")], true)
	assert.Equal(t, logger.Enabled[Level("ERROR")], true)
	assert.Equal(t, logger.Enabled[Level("INFO")], true)
	assert.Equal(t, logger.Enabled[Level("FATAL")], true)

	logger.SetLevel(Level("FATAL"))
	assert.Equal(t, logger.Level(), Level("FATAL"))
	assert.Equal(t, logger.Enabled[Level("TRACE")], false)
	assert.Equal(t, logger.Enabled[Level("DEBUG")], false)
	assert.Equal(t, logger.Enabled[Level("WARN")], false)
	assert.Equal(t, logger.Enabled[Level("ERROR")], false)
	assert.Equal(t, logger.Enabled[Level("INFO")], false)
	assert.Equal(t, logger.Enabled[Level("FATAL")], true)

	logger.SetLevel(Level("INFO"))
	assert.Equal(t, logger.Level(), Level("INFO"))
	assert.Equal(t, logger.Enabled[Level("TRACE")], false)
	assert.Equal(t, logger.Enabled[Level("DEBUG")], false)
	assert.Equal(t, logger.Enabled[Level("WARN")], false)
	assert.Equal(t, logger.Enabled[Level("ERROR")], true)
	assert.Equal(t, logger.Enabled[Level("INFO")], true)
	assert.Equal(t, logger.Enabled[Level("FATAL")], true)
}

func TestUnwrap(t *testing.T) {
	args := unwrap(func(args ...interface{}) []interface{} {
		return []interface{}{
			"example log message",
			"example args",
		}
	})
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 2)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")

	var passedArgs []interface{}
	args = unwrap(func(args ...interface{}) []interface{} {
		passedArgs = args
		return []interface{}{
			"example log message",
			"example args",
			"more example args",
		}
	}, "passed args", "more passed args")
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 3)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")
	assert.Equal(t, args[2], "more example args")
	assert.Equal(t, len(passedArgs), 2)
	assert.Equal(t, passedArgs[0], "passed args")
	assert.Equal(t, passedArgs[1], "more passed args")

	args = unwrap("example log message", "example args", "more args")
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 3)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")
	assert.Equal(t, args[2], "more args")

	args = unwrap(func() []interface{} {
		return []interface{}{
			"example log message",
			"example args",
			"more example args",
		}
	})
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 3)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")
	assert.Equal(t, args[2], "more example args")

	args = unwrap(func() (string, []interface{}) {
		return "example log message", []interface{}{
			"example args",
			"more example args",
		}
	})
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 3)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")
	assert.Equal(t, args[2], "more example args")
}

func TestCompose(t *testing.T) {
	msg, args := compose(Level("DEBUG"), "test message", "example arg")
	assert.NotNil(t, msg)
	assert.Equal(t, msg, "[%s] [%s] test message\n")
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 3)
	assert.Equal(t, args[1], "DEBUG")
	assert.Equal(t, args[2], "example arg")
}

func TestWrite(t *testing.T) {

}

func TestLog(t *testing.T) {

}

func TestGlobalFuncs(t *testing.T) {

}

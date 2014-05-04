package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New(DEBUG)
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), DEBUG)
}

func TestGlobal(t *testing.T) {
	logger := Global()
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), DEBUG)
}

func TestLevels(t *testing.T) {
	logger := Global()
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), DEBUG)
	assert.Equal(t, logger.Enabled[TRACE], false)
	assert.Equal(t, logger.Enabled[DEBUG], true)
	assert.Equal(t, logger.Enabled[WARN], true)
	assert.Equal(t, logger.Enabled[ERROR], true)
	assert.Equal(t, logger.Enabled[INFO], true)
	assert.Equal(t, logger.Enabled[FATAL], true)

	logger.SetLevel(TRACE)
	assert.Equal(t, logger.Level(), TRACE)
	assert.Equal(t, logger.Enabled[TRACE], true)
	assert.Equal(t, logger.Enabled[DEBUG], true)
	assert.Equal(t, logger.Enabled[WARN], true)
	assert.Equal(t, logger.Enabled[ERROR], true)
	assert.Equal(t, logger.Enabled[INFO], true)
	assert.Equal(t, logger.Enabled[FATAL], true)

	logger.SetLevel(FATAL)
	assert.Equal(t, logger.Level(), FATAL)
	assert.Equal(t, logger.Enabled[TRACE], false)
	assert.Equal(t, logger.Enabled[DEBUG], false)
	assert.Equal(t, logger.Enabled[WARN], false)
	assert.Equal(t, logger.Enabled[ERROR], false)
	assert.Equal(t, logger.Enabled[INFO], false)
	assert.Equal(t, logger.Enabled[FATAL], true)

	logger.SetLevel(INFO)
	assert.Equal(t, logger.Level(), INFO)
	assert.Equal(t, logger.Enabled[TRACE], false)
	assert.Equal(t, logger.Enabled[DEBUG], false)
	assert.Equal(t, logger.Enabled[WARN], false)
	assert.Equal(t, logger.Enabled[ERROR], true)
	assert.Equal(t, logger.Enabled[INFO], true)
	assert.Equal(t, logger.Enabled[FATAL], true)
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
}

func TestWrite(t *testing.T) {

}

func TestLog(t *testing.T) {

}

func TestGlobalFuncs(t *testing.T) {

}

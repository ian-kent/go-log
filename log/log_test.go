package log

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	logger := New(DEBUG)
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level, DEBUG)
}

func TestUnwrap(t *testing.T) {
	args := Unwrap(func(args ...interface{})[]interface{}{
		return []interface{}{
			"example log message",
			"example args",
		}
	});
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 2)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")

	var passedArgs []interface{}
	args = Unwrap(func(args ...interface{})[]interface{}{
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

	args = Unwrap("example log message", "example args", "more args")
	assert.NotNil(t, args)
	assert.Equal(t, len(args), 3)
	assert.Equal(t, args[0], "example log message")
	assert.Equal(t, args[1], "example args")
	assert.Equal(t, args[2], "more args")
}

func TestWrite(t *testing.T) {

}

func TestLog(t *testing.T) {

}

func TestLevels(t *testing.T) {

}

func TestGlobalFuncs(t *testing.T) {

}

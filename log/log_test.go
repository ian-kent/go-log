package log

import (
	"github.com/ian-kent/go-log/levels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := Logger()
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), Stol("DEBUG"))
	assert.NotNil(t, logger.Name())
	assert.Equal(t, logger.Name(), "")
	assert.Equal(t, logger.FullName(), "")

	logger = Logger("foo")
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), Stol("DEBUG"))
	assert.NotNil(t, logger.Name())
	assert.Equal(t, logger.Name(), "foo")
	assert.Equal(t, logger.FullName(), "foo")

	logger = Logger("foo.bar")
	assert.NotNil(t, logger)
	assert.Equal(t, logger.Level(), Stol("DEBUG"))
	assert.NotNil(t, logger.Name())
	assert.Equal(t, logger.Name(), "bar")
	assert.Equal(t, logger.FullName(), "foo.bar")
}

func TestLevel(t *testing.T) {
	for k, s := range levels.LogLevelsToString {
		assert.Equal(t, Stol(s), k)
	}
}

func TestLog(t *testing.T) {

}

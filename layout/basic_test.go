package layout

import (
	"github.com/ian-kent/go-log/levels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	b := Basic()
	assert.NotNil(t, b)

	assert.Equal(t, b.Format(levels.DEBUG, "Test message"), "Test message")
	assert.Equal(t, b.Format(levels.DEBUG, "Test message %s", "test"), "Test message test")
}

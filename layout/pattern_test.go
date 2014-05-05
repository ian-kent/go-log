package layout

import (
	"github.com/ian-kent/go-log/levels"
	"github.com/stretchr/testify/assert"
	"testing"
	//"os"
)

func TestPattern(t *testing.T) {
	p := Pattern("")
	assert.NotNil(t, p)

	assert.Equal(t, p.Format(levels.DEBUG, "Test message %s", "test"), "")
	assert.Equal(t, p.Format(levels.DEBUG, "%c %C %d"), "")

	p = Pattern("%c")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "layout")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "layout")

	p = Pattern("%C")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "layout")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "layout")

	p = Pattern("%d")
	assert.NotNil(t, p)
	// FIXME
	//assert.Equal(t, p.Format(levels.DEBUG, ""), time.Now().String())
	//assert.Equal(t, p.Format(levels.DEBUG, "foo"), time.Now().String())

	p = Pattern("%F")
	assert.NotNil(t, p)
	//assert.Equal(t, p.Format(levels.DEBUG, ""), os.Getenv("GOPATH") + "/src/github.com/ian-kent/go-log/layout/pattern_test.go")
	//assert.Equal(t, p.Format(levels.DEBUG, "foo"), os.Getenv("GOPATH") + "/src/github.com/ian-kent/go-log/layout/pattern_test.go")

	p = Pattern("%l")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "layout/pattern_test.go:40")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "layout/pattern_test.go:41")

	p = Pattern("%L")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "45")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "46")

	p = Pattern("%m")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, "test message"), "test message")
	assert.Equal(t, p.Format(levels.DEBUG, "test message %s", "test"), "test message test")
	assert.Equal(t, p.Format(levels.DEBUG, "test message %d", 2), "test message 2")

	p = Pattern("%n")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "\n")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "\n")

	p = Pattern("%p")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "DEBUG")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "DEBUG")

	p = Pattern("%r")
	assert.NotNil(t, p)
	// FIXME
	//assert.Equal(t, p.Format(levels.DEBUG, ""), "FIXME")
	//assert.Equal(t, p.Format(levels.DEBUG, "foo"), "FIXME")

	p = Pattern("%x")
	assert.NotNil(t, p)
	// FIXME
	//assert.Equal(t, p.Format(levels.DEBUG, ""), "FIXME")
	//assert.Equal(t, p.Format(levels.DEBUG, "foo"), "FIXME")

	p = Pattern("%X")
	assert.NotNil(t, p)
	// FIXME
	//assert.Equal(t, p.Format(levels.DEBUG, ""), "FIXME")
	//assert.Equal(t, p.Format(levels.DEBUG, "foo"), "FIXME")

	p = Pattern("%%")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "%")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "%")

	p = Pattern("%c %C %l %L %p %m%n")
	assert.NotNil(t, p)
	assert.Equal(t, p.Format(levels.DEBUG, ""), "layout layout layout/pattern_test.go:89 89 DEBUG \n")
	assert.Equal(t, p.Format(levels.DEBUG, "foo"), "layout layout layout/pattern_test.go:90 90 DEBUG foo\n")
	assert.Equal(t, p.Format(levels.DEBUG, "foo=%s", "bar"), "layout layout layout/pattern_test.go:91 91 DEBUG foo=bar\n")
}

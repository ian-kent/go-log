package appenders

import (
	"github.com/ian-kent/go-log/levels"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRollingFile(t *testing.T) {
	os.Remove("rollingfile_test.log")
	a := RollingFile("rollingfile_test.log", true)
	assert.NotNil(t, a)
	_, err := os.Stat("rollingfile_test.log")
	assert.Nil(t, err)

	a.Write(levels.DEBUG, "Test message")
	f, err := os.Open("rollingfile_test.log")
	assert.Nil(t, err)
	assert.NotNil(t, f)

	b := make([]byte, 13)
	n, err := f.Read(b)

	assert.Equal(t, n, 13)
	assert.Equal(t, string(b), "Test message\n")

	f.Close()
}

func TestRollingFileTruncate(t *testing.T) {
	os.Remove("rollingfile_test.log")
	a := RollingFile("rollingfile_test.log", true)
	assert.NotNil(t, a)
	_, err := os.Stat("rollingfile_test.log")
	assert.Nil(t, err)

	a.Write(levels.DEBUG, "Test message")
	f, err := os.Open("rollingfile_test.log")
	assert.Nil(t, err)
	assert.NotNil(t, f)

	b := make([]byte, 13)
	n, err := f.Read(b)

	assert.Equal(t, n, 13)
	assert.Equal(t, string(b), "Test message\n")

	f.Close()
	a.Close()

	a = RollingFile("rollingfile_test.log", false)
	assert.NotNil(t, a)
	_, err = os.Stat("rollingfile_test.log")
	assert.Nil(t, err)

	a.Write(levels.DEBUG, "Foo")
	f, err = os.Open("rollingfile_test.log")
	assert.Nil(t, err)
	assert.NotNil(t, f)

	b = make([]byte, 4)
	n, err = f.Read(b)

	assert.Equal(t, n, 4)
	assert.Equal(t, string(b), "Foo\n")

	f.Close()
	a.Close()
}

func TestRollingFileRotate(t *testing.T) {
	os.Remove("rollingfile_test.log")
	a := RollingFile("rollingfile_test.log", true)
	assert.NotNil(t, a)
	_, err := os.Stat("rollingfile_test.log")
	assert.Nil(t, err)

	a.MaxFileSize = 20

	a.Write(levels.DEBUG, "Test message")

	f, err := os.Open("rollingfile_test.log")
	assert.Nil(t, err)
	assert.NotNil(t, f)

	b := make([]byte, 13)
	n, err := f.Read(b)

	assert.Equal(t, n, 13)
	assert.Equal(t, string(b), "Test message\n")

	a.Write(levels.DEBUG, "Another test")
	b = make([]byte, 13)
	n, err = f.Read(b)

	assert.Equal(t, n, 13)
	assert.Equal(t, string(b), "Another test\n")

	a.Write(levels.DEBUG, "Yet another test")

	f.Close()
	f, err = os.Open("rollingfile_test.log")
	assert.Nil(t, err)
	assert.NotNil(t, f)

	b = make([]byte, 17)
	n, err = f.Read(b)

	assert.Equal(t, n, 17)
	assert.Equal(t, string(b), "Yet another test\n")

	f.Close()

	f, err = os.Open("rollingfile_test.log.1")
	assert.Nil(t, err)
	assert.NotNil(t, f)

	b = make([]byte, 26)
	n, err = f.Read(b)

	assert.Equal(t, n, 26)
	assert.Equal(t, string(b), "Test message\nAnother test\n")

	f.Close()
	a.Close()
}

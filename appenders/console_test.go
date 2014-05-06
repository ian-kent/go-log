package appenders

import (
	"bytes"
	"github.com/ian-kent/go-log/levels"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestConsole(t *testing.T) {
	a := Console()
	assert.NotNil(t, a)

	assert.Equal(t, captureOutput(t, func() { a.Write(levels.DEBUG, "Test message") }), "Test message\n")
	assert.Equal(t, captureOutput(t, func() { a.Write(levels.DEBUG, "Test message %s", "foo") }), "Test message foo\n")
}

func captureOutput(t *testing.T, f func()) string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = stdout

	out := <-outC

	return out
}

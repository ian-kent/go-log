package layout

import (
	"github.com/ian-kent/go-log/levels"
	"fmt"
)

type basicLayout struct {
	Layout
}

func Basic() *basicLayout {
	return &basicLayout{}
}

func (a *basicLayout) Format(level levels.LogLevel, message string, args ...interface{}) string {
	return fmt.Sprintf(message, args...)
}

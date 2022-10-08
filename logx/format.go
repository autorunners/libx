package logx

import (
	"fmt"
)

func newDefaultFormater() Formatter {
	return defaultFormatter{}
}

type Formatter interface {
	Format(prefix string, v ...interface{}) string
}

// defaultFormatter
type defaultFormatter struct {
	prefix string
}

var _ Formatter = defaultFormatter{}

func (f defaultFormatter) Format(prefix string, v ...interface{}) string {
	return fmt.Sprintln(append(v, prefix)...)
}

// jsonFormatter
type jsonFormatter struct {
}

func (j jsonFormatter) Format(prefix string, v ...interface{}) string {
	panic("implement me")
}

var _ Formatter = jsonFormatter{}

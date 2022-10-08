package logx

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Priority int

const (
	// Severity.
	// From: sys/syslog.h on Linux, BSD, and OS X.

	LogAlert Priority = iota
	LogErr
	LogWarning
	LogInfo
	LogDebug
)

var (
	defaultLogx Logx

	prioritys = map[string]Priority{
		"alert":   LogAlert,
		"err":     LogErr,
		"warning": LogWarning,
		"info":    LogInfo,
		"debug":   LogDebug,
	}
	defaultConfig = Config{
		Prefix: "[logx]",
		Level:  "info",
		Type:   "console",
		Path:   "logs",
	}
)

func init() {
	content, err := ioutil.ReadFile("./etc/logx.yaml")
	if err != nil {
		initLogx(defaultConfig)
		return
	}
	var c Config
	err = yaml.Unmarshal(content, &c)
	log.Println(c)
	if err != nil {
		log.Println(err)
		initLogx(defaultConfig)
		return
	}
	initLogx(c)
}

// Logx is a logx interface
type Logx interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Level() Priority
	Logger() Logger
	SetLevel(l Priority)
	SetLogger(logger Logger)
}

var _ Logx = logx{}

func InitLogx(c Config) error {
	_, ok := prioritys[c.Level]
	if !ok {
		return fmt.Errorf("not support Level")
	}

	initLogx(c)
	return nil
}

func initLogx(c Config) {
	priority := prioritys[c.Level]
	logger := newLogger(c)

	defaultLogx = logx{
		logger: logger,
		level:  priority,
	}
	return
}

type logx struct {
	logger Logger
	level  Priority
}

func (d logx) Logger() Logger {
	return d.logger
}

func (d logx) SetLogger(logger Logger) {
	panic("implement me")
}

func (d logx) Debug(v ...interface{}) {
	d.log(LogDebug, v...)
}

func (d logx) Debugf(format string, v ...interface{}) {
	d.logf(LogDebug, format, v...)
}

func (d logx) Error(v ...interface{}) {
	d.log(LogErr, v...)
}

func (d logx) Errorf(format string, v ...interface{}) {
	d.logf(LogErr, format, v...)
}

func (d logx) Info(v ...interface{}) {
	d.log(LogInfo, v...)
}

func (d logx) Infof(format string, v ...interface{}) {
	d.logf(LogInfo, format, v...)
}

func (d logx) Warn(v ...interface{}) {
	d.log(LogWarning, v...)
}

func (d logx) Warnf(format string, v ...interface{}) {
	d.logf(LogWarning, format, v)
}

func (d logx) Level() Priority {
	panic("implement me")
}

func (d logx) SetLevel(l Priority) {
	panic("implement me")
}

func (d logx) log(level Priority, v ...interface{}) {
	if level > d.level {
		return
	}
	d.logger.Print(v...)
}

func (d logx) logf(level Priority, format string, v ...interface{}) {
	if level > d.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	d.logger.Print(msg)
}

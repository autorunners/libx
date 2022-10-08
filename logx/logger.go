package logx

var (
	defaultLogger = newDefaultLogger()
)

func newDefaultLogger() Logger {
	return logger{
		formater: newDefaultFormater(),
		engine:   newDefaultEngine(),
	}
}

func newLogger(c Config) Logger {
	engine := newDefaultEngineFactory().Gen(c)
	return logger{
		formater: newDefaultFormater(), // todo
		engine:   engine,
	}
}

// Logger is a logger interface
type Logger interface {
	Print(v ...interface{})
}

// logger
type logger struct {
	formater Formatter
	engine   Engine
}

var _ Logger = logger{}

func (d logger) Print(v ...interface{}) {
	message := d.formater.Format("[logger format]", v...)
	d.engine.Output(message)
}

// complexLogger
type complexLogger struct {
	loggers []Logger
}

var _ Logger = complexLogger{}

func (c complexLogger) Print(v ...interface{}) {
	for _, l := range c.loggers {
		l.Print(v...)
	}
}

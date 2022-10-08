package logx

import (
	"log"
)

const (
	calldepth = 6
)

func newDefaultEngine() Engine {
	return noneEngine{}
}

type Engine interface {
	Output(message string)
}

// noneEngine
type noneEngine struct {}
var _ Engine = noneEngine{}

func (n noneEngine) Output(message string) {
}



// consoleEngine
type consoleEngine struct {
	engine *log.Logger
}
var _ Engine = consoleEngine{}

func (d consoleEngine) Output(message string) {
	d.engine.Output(calldepth, message)
}


// fileEngine
type fileEngine struct {
	engine *log.Logger
}
var _ Engine = fileEngine{}

func (d fileEngine) Output(message string) {
	d.engine.Output(calldepth, message)
}






package logx

import (
	"fmt"
	"log"
	"os"
)

var (
	engineFactorys = map[string]EngineFactory{
		"console": consoleEngineFactory{},
		"file": fileEngineFactory{},
	}
)

func newDefaultEngineFactory() EngineFactory {
	return engineFactory{}
}

type EngineFactory interface {
	Gen(Config) Engine
}

// engineFactory
type engineFactory struct {}
var _ EngineFactory = engineFactory{}

func (e engineFactory) Gen(config Config) Engine {
	factory, got := engineFactorys[config.Type]
	if got {
		return factory.Gen(config)
	}
	panic(fmt.Sprint("not support engine type: ", config.Type))
}


type consoleEngineFactory struct {}
var _ EngineFactory = consoleEngineFactory{}

func (c consoleEngineFactory) Gen(config Config) Engine {
	flag := log.Ldate | log.Ltime | log.Llongfile
	return consoleEngine{
		engine: log.New(os.Stderr, config.Prefix, flag),
	}
}


type fileEngineFactory struct {}
var _ EngineFactory = fileEngineFactory{}

func (f fileEngineFactory) Gen(config Config) Engine {
	flag := log.Ltime | log.Llongfile
	out, err := os.OpenFile(config.Path, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return fileEngine{
		engine: log.New(out, config.Prefix, flag),
	}
}




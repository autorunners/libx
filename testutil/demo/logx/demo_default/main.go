package main

import (
	"github.com/autorunners/libx/logx"
	"log"
)

type Config struct {
	Logger logx.Config `yaml:"logger"`
}

func init() {
	log.SetFlags(log.Llongfile | log.Ltime)
}

func main() {

	logx.Debug("111")
	logx.Info("111")
	logx.Infof("[format]%s", "111")

}

package main

import (
	"github.com/autorunners/libx/logx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Logger logx.Config `yaml:"logger"`
}

func init() {
	log.SetFlags(log.Llongfile | log.Ltime)
}

func main() {
	body, _ := ioutil.ReadFile("./demo/logx/etc/config.yaml")
	var c Config
	yaml.Unmarshal(body, &c)

	//logx.InitLogx(c.Logger)

	logx.Debug("111")
	logx.Info("111")
	//logx.Warn("111")
	//logx.Error("111")

}

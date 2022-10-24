package main

import (
	"github.com/autorunners/libx/mqttx"
	"time"
)

func main() {
	configFile := "./testutil/testdata/mqtt/config.yaml"
	mqttx.ConectDefault(configFile)
	time.Sleep(time.Hour)
}

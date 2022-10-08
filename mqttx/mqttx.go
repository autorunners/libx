package mqttx

import (
	"github.com/autorunners/libx/mqttx/client"
	"github.com/autorunners/libx/mqttx/config"
	"github.com/autorunners/libx/mqttx/handler"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Connect(c config.Config, connectHandler mqtt.OnConnectHandler, messageHandler mqtt.MessageHandler, connectLostHandler mqtt.ConnectionLostHandler) {
	client.New().Connect(c, connectHandler, messageHandler, connectLostHandler)
}

func ConectDefault(configFile string) {
	c := config.Parse(configFile)
	client.New().Connect(c, handler.DefaultConnectHandler, handler.DefaultMessageHandler, handler.DefaultConnectLostHandler)
}

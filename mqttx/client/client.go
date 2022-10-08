package client

import (
	"fmt"
	"github.com/autorunners/libx/mqttx/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

func New() *MqttClient {
	return new(MqttClient)
}

type MqttClient struct {
	client mqtt.Client
}

func (m *MqttClient) Connect(c config.Config, connectHandler mqtt.OnConnectHandler, messageHandler mqtt.MessageHandler, connectLostHandler mqtt.ConnectionLostHandler) {
	broker := fmt.Sprintf("%s://%s:%s", c.Protocol, c.Host, c.Port)
	log.Println(broker)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	opts.SetClientID(c.ClientId)
	opts.SetUsername(c.User)
	opts.SetPassword(c.Pass)
	opts.DefaultPublishHandler = messageHandler
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	//opts.WillEnabled = false
	opts.KeepAlive = c.KeepAlive                // 心跳时间间隔
	opts.MaxReconnectInterval = 1 * time.Minute // 重连时间间隔

	if c.Protocol == "ssl" {
		opts.SetTLSConfig(config.NewTLSConfig(c.CertFiles))
	}

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Println("connect error := ", token.Error())
		panic(token.Error())
	}
	m.client = client
	m.Subscribe(c.Topics, messageHandler)
}

func (m *MqttClient) Subscribe(topics []string, handler mqtt.MessageHandler) {
	for _, topic := range topics {
		m.client.Subscribe(topic, 0, handler)
	}
}

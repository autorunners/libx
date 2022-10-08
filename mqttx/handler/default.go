package handler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func DefaultMessageHandler(client mqtt.Client, msg mqtt.Message) {
	payload := msg.Payload()
	log.Println(string(payload))
}

func DefaultConnectHandler(client mqtt.Client) {
	log.Println("MQTT Connected")
}

func DefaultConnectLostHandler(client mqtt.Client, err error) {
	log.Println("MQTT Connect lost: ", err)
}

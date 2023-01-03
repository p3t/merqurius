package client

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func Connect() error {

	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	broker := "localhost"
	port := 1884
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(2 * time.Second)
	opts.SetClientID("merqurius@" + hostname)
	opts.SetAutoReconnect(false)
	// opts.SetMaxReconnectInterval()
	opts.SetPassword("mqtt")
	opts.SetUsername("mqtt")

	opts.SetDefaultPublishHandler(f)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func Disconnect() {
	client.Disconnect(10000)
}

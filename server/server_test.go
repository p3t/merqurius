package server

import (
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
)

func TestServerListnes(t *testing.T) {
	assert := assert.New(t)
	done := make(chan bool)

	opts := mqtt.NewClientOptions()
	opts.SetClientID("TestServerListens")
	opts.SetConnectRetry(false)
	opts.AddBroker("tcp://localhost:1883")

	connected := false
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		t.Log("Connected")
		connected = true
		done <- true
	})

	Start()

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Error() != nil {
		t.Errorf("token.Error: %s", token.Error())
	}

	select {
	case <-done:
		t.Log("Connected")
	case <-time.After(5 * time.Second):
		t.Fatalf("Timeout")
	}

	assert.True(connected)

}

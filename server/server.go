package server

import (
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
)

var server *mqtt.Server

func Start() {

	fmt.Println(aurora.Magenta("Mochi MQTT Server initializing..."), aurora.Cyan("TCP"))

	sOpts := mqtt.New().Options
	server = mqtt.NewServer(sOpts)

	tcp := listeners.NewTCP("t1", "localhost:1883")
	err := server.AddListener(tcp, nil)
	if err != nil {
		log.Fatal(err)
	}

	server.Events.OnConnect = func(client events.Client, pk events.Packet) {
		fmt.Printf("OnConnect> client %s: %+v\n", client.ID, err)
	}
	// Start the server
	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println(aurora.BgMagenta("  Started!  "))

}

func Stop() {
	server.Close()
	fmt.Println(aurora.BgGreen("  Server closed  "))
}

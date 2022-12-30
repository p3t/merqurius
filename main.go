package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"merqurius/client"
	"merqurius/server"

	"github.com/logrusorgru/aurora"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	server.Start()
	if err := client.Connect(); err != nil {
		fmt.Println(aurora.BgRed(fmt.Sprintf("Client connection error: %s", err)))
	}

	<-done
	fmt.Println()
	fmt.Println(aurora.BgRed("  Caught Signal  "))
	client.Disconnect()
	server.Stop()
	fmt.Println(aurora.BgGreen("  Finished  "))
}

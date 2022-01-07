package main

import (
	"os"
	"os/signal"
	"skillbox_diploma/pkg/config"
	"skillbox_diploma/pkg/server"
	"skillbox_diploma/pkg/simulator"
	"syscall"
)

func main() {
	config.GlobalConfig = config.NewConfig("config.yaml")

	go simulator.StartSimulatorServer()
	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}

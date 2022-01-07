package main

import (
	"github.com/astrviktor/skillbox_diploma/pkg/config"
	"github.com/astrviktor/skillbox_diploma/pkg/server"
	"github.com/astrviktor/skillbox_diploma/pkg/simulator"
	"os"
	"os/signal"
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

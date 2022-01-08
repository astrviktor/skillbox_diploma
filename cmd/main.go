package main

import (
	"github.com/astrviktor/skillbox_diploma/pkg/config"
	"github.com/astrviktor/skillbox_diploma/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.GlobalConfig = config.NewConfig("config.yaml")
	config.GlobalConfig = config.ForHerokuConfig(config.GlobalConfig)

	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}

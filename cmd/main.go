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

	// for heroku PORT
	port := os.Getenv("PORT")
	config.GlobalConfig.Addr = ":" + port
	config.GlobalConfig.MMSAddr = "http://127.0.0.1:" + port + "/mms"
	config.GlobalConfig.SupportAddr = "http://127.0.0.1:" + port + "/support"
	config.GlobalConfig.IncidentAddr = "http://127.0.0.1:" + port + "/incident"

	go server.StartServer()

	exit := make(chan os.Signal, 0)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
}

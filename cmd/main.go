package main

import (
	"fmt"
	"skillbox_diploma/pkg/result"
	"skillbox_diploma/pkg/simulator"
	"time"
)

func main() {
	go simulator.StartSimulatorServer()
	time.Sleep(5 * time.Second)

	//fmt.Println(result.GetResultData())
	result := result.GetResultData()

	fmt.Println("")
	fmt.Println("SMS")
	fmt.Printf("%+v\n", result.SMS)

	fmt.Println("")
	fmt.Println("MMS")
	fmt.Printf("%+v\n", result.MMS)

	fmt.Println("")
	fmt.Println("VoiceCall")
	fmt.Printf("%+v\n", result.VoiceCall)

	fmt.Println("")
	fmt.Println("Email")
	fmt.Printf("%+v\n", result.Email)

	fmt.Println("")
	fmt.Println("Billing")
	fmt.Printf("%+v\n", result.Billing)

	fmt.Println("")
	fmt.Println("Support")
	fmt.Printf("%+v\n", result.Support)

	fmt.Println("")
	fmt.Println("Incidents")
	fmt.Printf("%+v\n", result.Incidents)
}

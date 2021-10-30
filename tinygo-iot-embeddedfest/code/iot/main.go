package main

import ( // OMIT
	// OMIT
	"machine"
	"time" // OMIT

	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/mqtt" // OMIT
	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/wifi" // OMIT
	// OMIT
) // OMIT

func main() {
	wifi.Setup()
	wifi.ConnectToAP()

	client := mqtt.NewClient()
	// OMIT
	err := mqtt.Connect(client)
	if err != nil {
		println(err.Error())
		time.Sleep(time.Second)
		machine.ResetProcessor()
	}

	for {
		mqtt.PublishMessage(client)

		time.Sleep(time.Second)
	}
}

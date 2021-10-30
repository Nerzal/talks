package main

import ( // OMIT
	"time" // OMIT

	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/mqttclient"
) // OMIT
// OMIT
func main() {
	sensor := InitializeSensor()
	client := InitializeMQTTClient()

	for {
		temp, err := sensor.ReadTemperature()
		if err != nil {
			println(err.Error())
		}

		mqttclient.PublishMessage(client, temp)
		time.Sleep(time.Second)
	}
}

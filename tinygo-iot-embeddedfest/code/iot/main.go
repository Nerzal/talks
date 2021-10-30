package main

import ( // OMIT
	// OMIT
	"device/arm" // OMIT
	"machine"    // OMIT
	"time"       // OMIT

	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/mqtt" // OMIT
	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/wifi" // OMIT
	"tinygo.org/x/drivers/lsm6ds3"                                  // OMIT
	// OMIT
) // OMIT

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})
	sensor := lsm6ds3.New(machine.I2C0)

	wifi.Setup()
	wifi.ConnectToAP()

	client := mqtt.NewClient()
	// OMIT
	err := mqtt.Connect(client)
	if err != nil {
		println(err.Error())
		time.Sleep(time.Second)
		arm.SystemReset()
	}

	for {
		temp, err := sensor.ReadTemperature()
		if err != nil {
			println(err.Error())
		}

		mqtt.PublishMessage(client, temp)
		time.Sleep(time.Second)
	}
}

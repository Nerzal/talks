package main

import ( // OMIT
	"device/arm" // OMIT
	"machine"    // OMIT
	"time"       // OMIT

	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/mqttclient" // OMIT
	"github.com/Nerzal/talks/tinygo-iot-embeddedfest/code/iot/wifi"       // OMIT

	"tinygo.org/x/drivers/lsm6ds3"  // OMIT
	"tinygo.org/x/drivers/net/mqtt" // OMIT
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

// OMIT
func InitializeSensor() lsm6ds3.Device { // OMIT
	machine.I2C0.Configure(machine.I2CConfig{}) // OMIT
	return lsm6ds3.New(machine.I2C0)            // OMIT
} // OMIT
// OMIT
func InitializeMQTTClient() mqtt.Client { // OMIT
	wifi.Setup()       // OMIT
	wifi.ConnectToAP() // OMIT
	// OMIT
	client := mqttclient.NewClient() // OMIT
	// OMIT
	err := mqttclient.Connect(client) // OMIT
	if err != nil {                   // OMIT
		println(err.Error())    // OMIT
		time.Sleep(time.Second) // OMIT
		arm.SystemReset()       // OMIT
	} // OMIT
	// OMIT
	return client // OMIT
} // OMIT

package wifi

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/wifinina"
)

// access point info
const ssid = "NoobyGames"
const pass = "IchHasseLangeWlanZugangsDaten1312!"

var (
	spi     = machine.NINA_SPI
	adaptor *wifinina.Device
)

func Setup() {
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
		Mode:      0,
		LSBFirst:  false,
	})

	adaptor = wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()
	time.Sleep(2 * time.Second)
}

func ConnectToAP() {
	println("Connecting to " + ssid)

	adaptor.SetPassphrase(ssid, pass)

	for st, _ := adaptor.GetConnectionStatus(); st != wifinina.StatusConnected; {
		println("Connection status: " + st.String())
		time.Sleep(time.Second)
		st, _ = adaptor.GetConnectionStatus()

		if st == wifinina.StatusConnectFailed || st == wifinina.StatusDisconnected {
			adaptor.SetPassphrase(ssid, pass)
			time.Sleep(time.Second)
		}
	}

	println("Connected.")
}

package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.D2
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led2 := machine.D3
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	go func() {
		for {
			led2.Low()
			time.Sleep(time.Millisecond * 1000)

			led2.High()
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}

package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("nuit")
		led.Low()
		time.Sleep(time.Second * 1)

		println("jour")
		led.High()
		time.Sleep(time.Second * 2)
	}
}

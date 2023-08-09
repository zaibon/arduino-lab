package main

import (
	"machine"
	"time"
)

const pin = machine.D9

func main() {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		println("nuit")
		pin.Low()
		time.Sleep(time.Second * 1)

		println("jour")
		pin.High()
		time.Sleep(time.Second * 2)
	}
}

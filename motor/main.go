package main

import (
	"machine"
	"time"
)

const (
	motorA = machine.D9
	motorB = machine.D10
)

func main() {
	for _, m := range []machine.Pin{motorA, motorB} {
		m.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	forward()
	time.Sleep(time.Second * 2)
	reverse()
	time.Sleep(time.Second * 2)
	stop()
}

func forward() {
	motorA.High()
	motorB.Low()
}

func reverse() {
	motorA.Low()
	motorB.High()
}

func stop() {
	motorA.Low()
	motorB.Low()
}

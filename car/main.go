package main

import (
	"machine"
	"time"
)

const (
	right1 = machine.D5
	right2 = machine.D6
	left1  = machine.D9
	left2  = machine.D10
)

func main() {
	for _, m := range []machine.Pin{right1, right2, left1, left2} {
		m.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	forward()
	time.Sleep(time.Millisecond * 800)
	reverse()
	time.Sleep(time.Millisecond * 800)
	turnRight()
	time.Sleep(time.Millisecond * 800)
	turnLeft()
	time.Sleep(time.Millisecond * 800)

	stop()
}

func reverse() {
	right1.High()
	right2.Low()
	left1.Low()
	left2.High()
}

func forward() {
	right1.Low()
	right2.High()
	left1.High()
	left2.Low()
}

func turnRight() {
	right1.High()
	right2.Low()
	left1.High()
	left2.Low()
}

func turnLeft() {
	right1.Low()
	right2.High()
	left1.Low()
	left2.High()
}

func stop() {
	for _, m := range []machine.Pin{right1, right2, left1, left2} {
		m.Low()
	}
}

package main

import (
	"machine"
	"time"
)

const buzzerPin = machine.D9

var period uint64 = 1e9 / 500

func main() {
	buzzerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pwm := machine.Timer1
	// Configure the PWM with the given period.
	pwm.Configure(machine.PWMConfig{
		Period: period,
	})

	beep()
	// fadeOut(pwm)
}

func beep() {
	buzzerPin.High()
	time.Sleep(time.Millisecond * 50)
	buzzerPin.Low()
}

func fadeOut(pwm machine.PWM) {
	ch, err := pwm.Channel(buzzerPin)
	if err != nil {
		println(err.Error())
		return
	}
	for i := 1; i < 255; i++ {
		pwm.Set(ch, pwm.Top()/uint32(i))
		time.Sleep(time.Millisecond * 10)
	}
	pwm.Set(ch, 0)
}

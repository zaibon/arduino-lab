package main

import (
	"machine"
	"time"
)

var period uint64 = 1e9 / 500

func main() {
	pin := machine.D9
	pwm := machine.Timer1

	// Configure the PWM with the given period.
	pwm.Configure(machine.PWMConfig{
		Period: period,
	})

	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
		return
	}

	max := 128
	for {
		for i := 1; i < max; i++ {
			// This performs a stylish fade-out
			pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Millisecond * 5)
		}
		for i := max; i > 1; i-- {
			// This performs a stylish fade-in
			pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Millisecond * 5)
		}
	}
	// for {
	// 	pwm.Set(ch, pwm.Top()/uint32(1))
	// 	time.Sleep(time.Millisecond * 5)
	// 	pwm.Set(ch, pwm.Top()/uint32(100))
	// 	time.Sleep(time.Millisecond * 5)
	// }
}

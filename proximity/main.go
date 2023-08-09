package main

import (
	"machine"
	"time"
)

const (
	beepPin      = machine.D9
	proximityPin = machine.D2
)

var period uint64 = 1e9 / 500

func main() {
	beepPin.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	proximityPin.Configure(machine.PinConfig{
		Mode: machine.PinInput,
	})

machine.CPUFrequency()

	previous := true
	for {
		far := proximityPin.Get()
		println(far)
		if previous != far {
			previous = far
		}
		beepPin.Set(!far)
		// if !far {
		// 	beepPin.High()
		// } else {
		// 	beepPin.Low()
		// }

		time.Sleep(time.Millisecond)
	}

	// buzzerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// pwm := machine.Timer1
	// // Configure the PWM with the given period.
	// pwm.Configure(machine.PWMConfig{
	// 	Period: period,
	// })

	// beep()
	// fadeOut(pwm)

}

func beep() {
	beepPin.High()
	time.Sleep(time.Millisecond * 50)
	beepPin.Low()
}

// func fadeOut(pwm machine.PWM) {
// 	ch, err := pwm.Channel(buzzerPin)
// 	if err != nil {
// 		println(err.Error())
// 		return
// 	}
// 	for i := 1; i < 255; i++ {
// 		pwm.Set(ch, pwm.Top()/uint32(i))
// 		time.Sleep(time.Millisecond * 10)
// 	}
// 	pwm.Set(ch, 0)
// }

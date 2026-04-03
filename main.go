package main

import (
	"machine"
	"time"
)

func main() {
	ledPin := machine.D9
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pwm := machine.Timer1 // PB1 and PB2
	pwm.Configure(machine.PWMConfig{})
	pwm.SetPeriod(0)

	ch, err := pwm.Channel(ledPin)
	if err != nil {
		panic(err)
	}
	pwm.Set(ch, pwm.Top()) // D9 = PB1 => channel 0

	machine.InitADC()
	potPin := machine.ADC{Pin: machine.ADC5}
	potPin.Configure(machine.ADCConfig{})

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{
		BaudRate: 9600,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	})

	r := 255.0 / 65472.0 // map()  return (x - in_min) * (out_max - out_min) / (in_max - in_min) + out_min;

	for {
		v := potPin.Get() // 0 <= v <= 65472

		z := float64(v) * r // 0 <= z <= 255

		pwm.Set(ch, uint32(z))

		time.Sleep(time.Millisecond * 500)
	}
}

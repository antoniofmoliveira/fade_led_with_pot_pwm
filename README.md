# Arduino Uno SMD - Fade LED with pot PWM

## Fade LED with pot PWM in Arduino Uno SMD with Golang

PWM in TinyGo 0.40.1

 PWM | Channel | Pin | Port | Bits
 ----- | ----- | ----- | ----- | -----
 Timer0 | 1 | ~5 | PD5 | 8
 Timer0 | 0 | ~6 | PD6 | 8
 Timer1 | 0 | ~9 | PB1 | 16
 Timer1 | 1 | ~10 | PB2 | 16
 Timer2 | 0 | ~11 | PB3 | 8
 Timer2 | 1 | ~3 | PD3 | 8

* VS Code
* Golang 1.25.8
* TinyGo 0.40.1 [TinyGo](tinygo.org)
* TinyGo Toolkit for VS Code (amken.tinygo-toolkit)

## TinyGo Toolkit

* Connect board (arduino uno smd)
* View > Command Palette > TinyGo: Initialize TinyGo Project
* View > Command Palette > TinyGo: Configure Intellisense
* Project > Target (arduino uno)
* Project > Port (/dev/ttyUSB0)
* Project Baud Rate (9600)
* Project > Flash
* Project > Serial Monitor (optional)
* View > Command Palette > TinyGo: Stop Serial Monitor (optional)

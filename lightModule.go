package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

type ModulesStruct struct {
	module1 Module
	module2 Module
}

type Module struct {
	Pin           rpio.Pin
	State         bool
	StopBreathing chan bool
}

func (lm *Module) init() {
	lm.Pin.Mode(rpio.Pwm)
	lm.Pin.Freq(1000)
}

func (lm *Module) lightOn() {
	lm.Pin.DutyCycle(100, 100)
	lm.State = true
	fmt.Println("State", lm.State)
}

func (lm *Module) lightOff() {
	lm.Pin.DutyCycle(0, 100)
	lm.State = false
	fmt.Println("State", lm.State)
}

func (lm *Module) breathOn() {
	fmt.Println("Start breathing")
	fmt.Println("State", lm.State)
	intensityDown := lm.State
	var intensity uint32
	if lm.State {
		intensity = 100
	} else {
		intensity = 0
	}

	for {
		select {
		case <-lm.StopBreathing:
			fmt.Println("Stop infinite loop")
			return
		default:
			if intensityDown {
				intensity -= 1
			} else {
				intensity += 1
			}

			if intensity == 100 || intensity == 0 {
				intensityDown = !intensityDown
			}

			lm.Pin.DutyCycle(intensity, 100)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (lm *Module) breathOff() {
	fmt.Println("Stop breathing")
	fmt.Println("State", lm.State)
	lm.StopBreathing <- true

	var intensity uint32
	if lm.State {
		intensity = 100
	} else {
		intensity = 0
	}

	lm.Pin.DutyCycle(intensity, 100)
}

func (lm *Module) state() {
	fmt.Println("State", lm.State)
}

func InitRaspberryPins() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
}

var Modules ModulesStruct

func SetupLight() {
	InitRaspberryPins()

	// Add one Module
	module1 := Module{rpio.Pin(13), false, make(chan bool)}
	module1.init()

	module2 := Module{rpio.Pin(19), false, make(chan bool)}
	module2.init()

	// Add Modules to Global Variable
	Modules.module1 = module1
	Modules.module2 = module2
}

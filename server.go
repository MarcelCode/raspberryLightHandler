package main

import (
	"github.com/labstack/echo"
	"github.com/stianeikeland/go-rpio"
)

func main(){
	SetupLight()

	e := echo.New()
	e.GET("/light/on", LightOn)
	e.GET("/light/off", LightOff)
	e.GET("/light/breath-on", BreathOn)
	e.GET("/light/breath-off", BreathOff)
	e.GET("/light/state", State)

	e.Logger.Fatal(e.Start(":3000"))
	defer rpio.Close()
}
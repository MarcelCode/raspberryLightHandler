package main

import (
	"github.com/labstack/echo"
	"github.com/stianeikeland/go-rpio"
)

func main() {
	SetupLight()

	e := echo.New()
	e.GET("/light/:id/on", LightOn)
	e.GET("/light/:id/off", LightOff)
	e.GET("/light/:id/breath-on", BreathOn)
	e.GET("/light/:id/breath-off", BreathOff)
	e.GET("/light/:id/state", State)

	e.Logger.Fatal(e.Start(":3000"))
	defer rpio.Close()
}

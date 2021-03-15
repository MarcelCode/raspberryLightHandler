package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func LightOn(c echo.Context) (err error){
	go Modules.module1.lightOn()
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error){
	go Modules.module1.lightOff()
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error){
	go Modules.module1.breathOn()
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error){
	go Modules.module1.breathOff()
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error){
	fmt.Println(Modules.module1.State)
	return c.JSON(http.StatusOK, "state")
}

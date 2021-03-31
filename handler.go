package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func getModule(id string) (module Module) {
	switch id {
	case "1":
		module = Modules.module1
	case "2":
		module = Modules.module2
	default:
		fmt.Println("Module not implemented")
	}
	return
}

func LightOn(c echo.Context) (err error) {
	id := c.Param("id")
	module := getModule(id)
	module.lightOn()
	return c.JSON(http.StatusOK, "light on")
}

func LightOff(c echo.Context) (err error) {
	id := c.Param("id")
	module := getModule(id)
	module.lightOff()
	return c.JSON(http.StatusOK, "light off")
}

func BreathOn(c echo.Context) (err error) {
	id := c.Param("id")
	module := getModule(id)
	module.breathOn()
	return c.JSON(http.StatusOK, "breath on")
}

func BreathOff(c echo.Context) (err error) {
	id := c.Param("id")
	module := getModule(id)
	module.breathOff()
	return c.JSON(http.StatusOK, "breath off")
}

func State(c echo.Context) (err error) {
	id := c.Param("id")
	module := getModule(id)
	fmt.Println(module.State)
	return c.JSON(http.StatusOK, "state")
}

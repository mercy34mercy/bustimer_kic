package main

import (
	"net/http"
	"practice-colly/controller"

	"github.com/labstack/echo"
)

var e = echo.New()

func Routing() {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Busdes! Clean Architecture API</h1>")
	})

	e.GET("/timetable",func(c echo.Context) error{
		busstop := "円町"
		destination := "立命館大学行き"
		busstoptourlCtrl := controller.BusstoToUrlController{}
		url,err := busstoptourlCtrl.FindURL(busstop,destination)
		if err != nil {
		}
		timetablecontroller := controller.TimetableController{}
		timetable := timetablecontroller.FindTimetable(url)
		return c.JSON(http.StatusOK,timetable)
	})
}


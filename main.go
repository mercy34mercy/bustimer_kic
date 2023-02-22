package main

import (
	"bytes"
	"net/http"
	"os"
	"bustimerkic/controller"
	"bustimerkic/domain/model"
	"bustimerkic/infra"
	"bustimerkic/infra/localcache"
	"github.com/labstack/echo"
)


func main() {
	infra.Init("gorm.db")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	localcache.Init()
	// go localcache.CreateTimetableCache()
	// utils.Dbcreate()

	router := Routing()
	router.Debug = true
	router.Logger.Fatal(router.Start(":" + port))
	

}

func Routing() *echo.Echo{
	var e = echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Busdes! Clean Architecture API</h1>")
	})

	e.GET("/system/info", func(e echo.Context) error {
		var systemInfo = &model.SystemInfo{
			Status:  200,
			Message: "",
			Time:    "",
		}
		return e.JSON(http.StatusOK, systemInfo)
	})

	e.GET("/timetable", func(c echo.Context) error {

		busstop := c.QueryParam("fr")
		var destination bytes.Buffer
		destination.WriteString(c.QueryParam("to"))
		destination.WriteString("行き")

		l := localcache.GetGoChache()
		if x, found := l.Get(busstop + destination.String()); found {
			return c.JSON(http.StatusOK, x.(model.TimeTable))
		}
		timetablecontroller := controller.TimetableController{}
		timetable,err := timetablecontroller.FindTimetable(busstop, destination.String()); 
		if err != nil{
			return c.HTML(http.StatusNotFound,"<h1>404 NOT FOUND<h1>")
		}
		localcache.CreateCachefromTimetable(busstop, destination.String(), timetable)

		return c.JSON(http.StatusOK, timetable)
	})

	e.GET("/timetable/multi", func(c echo.Context) error {
		busstop := c.QueryParam("fr")
		destination := c.QueryParam("to")
		var query model.Query = model.Query{
			Destination: destination,
			Busstop:     busstop,
		}
		destinationlist, busstoplist := query.SplitDestination()
		timetableCtrl := controller.TimetableFromBusstopController{}
		timetable,err := timetableCtrl.FindTimetable(busstoplist, destinationlist)
		if err != nil{
			return c.HTML(http.StatusNotFound,"<h1>404 NOT FOUND<h1>")
		}
		return c.JSON(http.StatusOK, timetable)
	})

	e.GET("/nextbus", func(c echo.Context) error {
		busstop := c.QueryParam("fr")
		var destination bytes.Buffer
		destination.WriteString(c.QueryParam("to"))
		destination.WriteString("行き")

		l := localcache.GetGoChache()

		if x, found := l.Get(busstop + destination.String()); found {
			var time model.TimeTable = x.(model.TimeTable)
			approachInfoCtrl := controller.ApproachInfoFromTimeTableController{}
			approachInfo := approachInfoCtrl.FindApproachInfoFromTimeTable(time, busstop, destination.String())
			return c.JSON(http.StatusOK, approachInfo)
		}

		timetablecontroller := controller.TimetableController{}
		timetable,_ := timetablecontroller.FindTimetable(busstop, destination.String())

		localcache.CreateCachefromTimetable(busstop, destination.String(), timetable)

		approachInfoCtrl := controller.ApproachInfoFromTimeTableController{}
		approachinfo := approachInfoCtrl.FindApproachInfoFromTimeTable(timetable, busstop, destination.String())
		return c.JSON(http.StatusOK, approachinfo)
	})
	return e
}

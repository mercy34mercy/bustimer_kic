package route

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewBusstopRouter(e *echo.Echo) {
	e.GET("/busstop",func(ctx echo.Context) error {
		return returnHTML(ctx)
	})
	e.POST("/busstop",func(ctx echo.Context) error {
		return returnHTML(ctx)
	})
}

func returnHTML(e echo.Context) error{
	return e.HTML(http.StatusOK,"<h1> Busdes </h1>")
}
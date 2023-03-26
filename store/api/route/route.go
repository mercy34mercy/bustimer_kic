package route

import "github.com/labstack/echo"

func Setup() *echo.Echo {
	e := echo.New()
	NewBusstopRouter(e)
	return e
}
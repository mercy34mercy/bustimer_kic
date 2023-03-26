package domain

import (

	"github.com/labstack/echo"
)

type Busstop struct {
	Name string `json:"name"`
}

type BusstopUsecase interface {
	GetByID(c echo.Context, userID string)
	Fetch(c echo.Context, userID string, busstop string)
	Delete(c echo.Context, userID string, busstop string)
}

type BusstopRepository interface {
	GetByID(c echo.Context, userID string)([]Busstop,error)
	Fetch(c echo.Context, userID string, busstop string) error
	Delete(c echo.Context, userID string, busstop string) error
}
package domain

import (
	"github.com/labstack/echo"
)

const (
	CollectionBusstop = "busstop"
)

type Busstop struct {
	Name string `json:"name"`
}

type BusstopUsecase interface {
	GetByID(c echo.Context, userID string) ([]Busstop, error)
	Fetch(c echo.Context, userID string, busstop Busstop) error
	Delete(c echo.Context, userID string, busstop Busstop) error
}

type BusstopRepository interface {
	GetByID(c echo.Context, userID string) ([]Busstop, error)
	Fetch(c echo.Context, userID string, busstop Busstop) error
	Delete(c echo.Context, userID string, busstop Busstop) error
}

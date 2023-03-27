package route

import (
	"cloud.google.com/go/firestore"
	"github.com/mercy34mercy/bustimer_kic/store/bootstrap"
	"net/http"

	"github.com/labstack/echo"
)

func NewBusstopRouter(e *echo.Echo, env *bootstrap.Env, db *firestore.Client) {
	e.GET("/busstop", func(ctx echo.Context) error {
		return returnHTML(ctx)
	})
	e.POST("/busstop", func(ctx echo.Context) error {
		return returnHTML(ctx)
	})
}

func returnHTML(e echo.Context) error {
	return e.HTML(http.StatusOK, "<h1> Busdes </h1>")
}

package route

import (
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/bootstrap"
	"github.com/mercy34mercy/bustimer_kic/store/firestore"
)

func Setup(e *echo.Echo, env *bootstrap.Env, db *firestore.Client) *echo.Echo {
	NewBusstopRouter(e, env, db)
	return e
}

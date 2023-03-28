package route

import (
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/api/controller"
	"github.com/mercy34mercy/bustimer_kic/store/bootstrap"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
	"github.com/mercy34mercy/bustimer_kic/store/firestore"
	"github.com/mercy34mercy/bustimer_kic/store/repository"
	"github.com/mercy34mercy/bustimer_kic/store/usecase"
)

func NewBusstopRouter(e *echo.Echo, env *bootstrap.Env, db firestore.Client) {
	e.GET("/busstop/:userID", func(ctx echo.Context) error {
		br := repository.NewBusstopRespository(db, domain.CollectionBusstop)
		tc := &controller.BusstopController{
			BusstopUsecase: usecase.NewBusstopRepository(br),
		}
		return tc.Fetch(ctx)
	})
	e.POST("/busstop", func(ctx echo.Context) error {
		br := repository.NewBusstopRespository(db, domain.CollectionBusstop)
		tc := &controller.BusstopController{
			BusstopUsecase: usecase.NewBusstopRepository(br),
		}
		return tc.Fetch(ctx)
	})
	e.DELETE("busstop/:userID/:busstop", func(ctx echo.Context) error {
		br := repository.NewBusstopRespository(db, domain.CollectionBusstop)
		tc := &controller.BusstopController{
			BusstopUsecase: usecase.NewBusstopRepository(br),
		}
		return tc.Delete(ctx)
	})
}

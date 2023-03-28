package controller

import (
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
)

type BusstopController struct {
	BusstopUsecase domain.BusstopUsecase
}

func (bc *BusstopController) GetByID(ctx echo.Context) ([]domain.Busstop, error) {
	userID := ctx.Param("userID")

	busstop, err := bc.BusstopUsecase.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return busstop, nil
}

func (bc *BusstopController) Fetch(ctx echo.Context) error {
	var busStop domain.Busstop

	userID := ctx.Param("userID")
	if err := ctx.Bind(&busStop); err != nil {
		return err
	}

	if err := bc.BusstopUsecase.Fetch(ctx, userID, busStop); err != nil {
		return err
	}

	return nil
}

func (bc *BusstopController) Delete(ctx echo.Context) error {
	userID := ctx.Param("userID")
	busStopName := ctx.Param("busstop")

	busStop := domain.Busstop{
		Name: busStopName,
	}

	if err := bc.BusstopUsecase.Delete(ctx, userID, busStop); err != nil {
		return err
	}

	return nil
}

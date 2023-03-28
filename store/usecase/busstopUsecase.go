package usecase

import (
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
)

type busstopUsecase struct {
	busstopRepository domain.BusstopRepository
}

func NewBusstopRepository(busstopRepository domain.BusstopRepository) domain.BusstopUsecase {
	return &busstopUsecase{
		busstopRepository: busstopRepository,
	}
}

func (bu *busstopUsecase) GetByID(ctx echo.Context, userID string) ([]domain.Busstop, error) {
	return bu.busstopRepository.GetByID(ctx, userID)
}

func (bu *busstopUsecase) Fetch(ctx echo.Context, userID string, busstop domain.Busstop) error {
	return bu.busstopRepository.Fetch(ctx, userID, busstop)
}
func (bu *busstopUsecase) Delete(ctx echo.Context, userID string, busstop domain.Busstop) error {
	return bu.busstopRepository.Delete(ctx, userID, busstop)
}

package controller

import (
	repositoryimpl "github.com/mercy34mercy/bustimer_kic/bustimer/infra/repositoryImpl"
	bustimersqlc "github.com/mercy34mercy/bustimer_kic/bustimer/sqlc/gen"
	"github.com/mercy34mercy/bustimer_kic/bustimer/usecase"
)

type FindBusstopListController struct{}

func (ctrl FindBusstopListController) FindBusstopList(busname string) ([]bustimersqlc.GetBusstopAndDestinationRow, error) {
	busstoptotimetableRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result, err := usecase.NewGetBusstopListUseCaseImpl(busname, busstoptotimetableRepository).FindBusstopList()
	return result, err
}

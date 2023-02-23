package controller

import (
	repositoryimpl "bustimerkic/infra/repositoryImpl"
	bustimersqlc "bustimerkic/sqlc/gen"
	"bustimerkic/usecase"
)

type FindBusstopListController struct{}

func (ctrl FindBusstopListController) FindBusstopList(busname string)([]bustimersqlc.GetBusstopAndDestinationRow,error){
	busstoptotimetableRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result,err := usecase.NewGetBusstopListUseCaseImpl(busname,busstoptotimetableRepository).FindBusstopList()
	return result,err
}
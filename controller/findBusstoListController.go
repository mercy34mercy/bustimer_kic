package controller

import (
	"bustimerkic/domain/model"
	repositoryimpl "bustimerkic/infra/repositoryImpl"
	"bustimerkic/usecase"
)

type FindBusstopListController struct{}

func (ctrl FindBusstopListController) FindBusstopList(busname string)([]model.Busstop,error){
	busstoptotimetableRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result,err := usecase.NewGetBusstopListUseCaseImpl(busname,busstoptotimetableRepository).FindBusstopList()
	return result,err
}
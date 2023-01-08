package controller

import (
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type FindBusstopListController struct{}

func (ctrl FindBusstopListController) FindBusstopList(busname string)([]model.Busstop,error){
	busstoptotimetableRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result,err := usecase.NewGetBusstopListUseCaseImpl(busname,busstoptotimetableRepository).FindBusstopList()
	return result,err
}
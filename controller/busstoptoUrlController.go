package controller

import (
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type BusstoToUrlController struct{}

func (ctrl BusstoToUrlController) FindURL(busstop string,destination string)([]string,error){
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result,err := usecase.NewGetBusstopUrlUseCaseImpl(busstop,destination,busstoptourlRepository).FindURL()
	return result,err
}
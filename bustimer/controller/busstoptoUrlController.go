package controller

import (
	repositoryimpl "bustimerkic/infra/repositoryImpl"
	"bustimerkic/usecase"
)

type BusstoToUrlController struct{}

func (ctrl BusstoToUrlController) FindURL(busstop string,destination string)([]string,error){
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result,err := usecase.NewGetBusstopUrlUseCaseImpl(busstop,destination,busstoptourlRepository).FindURL()
	return result,err
}
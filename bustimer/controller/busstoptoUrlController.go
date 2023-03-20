package controller

import (
	repositoryimpl "github.com/mercy34mercy/bustimer_kic/bustimer/infra/repositoryImpl"
	"github.com/mercy34mercy/bustimer_kic/bustimer/usecase"
)

type BusstoToUrlController struct{}

func (ctrl BusstoToUrlController) FindURL(busstop string, destination string) ([]string, error) {
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	result, err := usecase.NewGetBusstopUrlUseCaseImpl(busstop, destination, busstoptourlRepository).FindURL()
	return result, err
}

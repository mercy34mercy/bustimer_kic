package controller

import (
	"bustimerkic/domain/model"
	repositoryImpl "bustimerkic/infra/repositoryImpl"
	"bustimerkic/usecase"
)

type TimetableFromBusstopController struct{}

func (ctrl TimetableFromBusstopController) FindTimetable(busstop []string,destination []string)(model.MultiTimeTable,error){
	busstoptotimetableRepository := repositoryImpl.NewBusstopToUrlRepositoryImpl()
	timetable,err := usecase.NewGetUrlFromBusstopUseCaseImpl(busstop,destination,busstoptotimetableRepository).FindURLFromBusstop()
	return timetable,err
}
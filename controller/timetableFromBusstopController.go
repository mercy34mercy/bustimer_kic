package controller

import (
	"practice-colly/domain/model"
	repositoryImpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type TimetableFromBusstopController struct{}

func (ctrl TimetableFromBusstopController) FindTimetable(busstop []string,destination []string)(model.MultiTimeTable,error){
	busstoptotimetableRepository := repositoryImpl.NewBusstopToUrlRepositoryImpl()
	timetable,err := usecase.NewGetUrlFromBusstopUseCaseImpl(busstop,destination,busstoptotimetableRepository).FindURLFromBusstop()
	return timetable,err
}
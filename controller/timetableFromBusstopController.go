package controller

import (
	"practice-colly/domain/model"
	repositoryImpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type TimetableFromBusstopController struct{}

func (ctrl TimetableFromBusstopController) FindTimetable(busstop string,destination []string)(model.MultiTimeTable){
	busstoptotimetableRepository := repositoryImpl.NewBusstopToUrlRepositoryImpl()
	timetable := usecase.NewGetUrlFromBusstopUseCaseImpl(busstop,destination,busstoptotimetableRepository).FindURLFromBusstop()
	return timetable
}
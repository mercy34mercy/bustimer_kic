package controller

import (
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type TimetableController struct{}


func (ctrl TimetableController) FindTimetable(Busstop string,Destination string)(model.TimeTable,error){
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	timetable,err := usecase.NewGetTimetableUseCaseImpl(Busstop,Destination,busstoptourlRepository).FindTimetable()
	return timetable,err
}
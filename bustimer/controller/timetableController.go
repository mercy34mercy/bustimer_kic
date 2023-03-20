package controller

import (
	"bustimerkic/domain/model"
	repositoryimpl "bustimerkic/infra/repositoryImpl"
	"bustimerkic/usecase"
)

type TimetableController struct{}


func (ctrl TimetableController) FindTimetable(Busstop string,Destination string)(model.TimeTable,error){
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	timetable,err := usecase.NewGetTimetableUseCaseImpl(Busstop,Destination,busstoptourlRepository).FindTimetable()
	return timetable,err
}
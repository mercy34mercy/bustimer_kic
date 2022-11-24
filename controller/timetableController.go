package controller

import (
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type TimetableController struct{}


func (ctrl TimetableController) FindTimetable(URL []string)(model.TimeTable){
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	timetable := usecase.NewGetTimetableUseCaseImpl(URL,busstoptourlRepository).FindTimetable()
	return timetable
}
package controller

import (
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type CacheTimetableController struct{}


func (ctrl CacheTimetableController) FindCacheTimetable(businfo model.Busstop)(model.TimeTable,string,string){
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	timetable,busstop,destination := usecase.NewCreateCacheUseCaseImpl(businfo,busstoptourlRepository).FindCacheTimetable()
	return timetable,busstop,destination
}
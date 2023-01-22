package controller

import (
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type ApproachInfoFromTimeTableController struct{}

func (ctrl ApproachInfoFromTimeTableController) FindApproachInfoFromTimeTable(timetable model.TimeTable, busstop string, via string) model.ApproachInfos {
	approachInfoFromTimeTableRepository := repositoryimpl.NewApproachInfoRepositoryImpl()
	result := usecase.NewGetApproachInfoFromTimeTableUseCaseImpl(timetable,busstop,via,approachInfoFromTimeTableRepository).FindApproachInfoFromTimeTable()
	return result
}

package controller

import (
	"bustimerkic/domain/model"
	repositoryimpl "bustimerkic/infra/repositoryImpl"
	"bustimerkic/usecase"
)

type ApproachInfoFromTimeTableController struct{}

func (ctrl ApproachInfoFromTimeTableController) FindApproachInfoFromTimeTable(timetable model.TimeTable, busstop string, via string) model.ApproachInfos {
	approachInfoFromTimeTableRepository := repositoryimpl.NewApproachInfoRepositoryImpl()
	result := usecase.NewGetApproachInfoFromTimeTableUseCaseImpl(timetable,busstop,via,approachInfoFromTimeTableRepository).FindApproachInfoFromTimeTable()
	return result
}

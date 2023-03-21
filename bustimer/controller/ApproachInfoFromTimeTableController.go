package controller

import (
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	repositoryimpl "github.com/mercy34mercy/bustimer_kic/bustimer/infra/repositoryImpl"
	"github.com/mercy34mercy/bustimer_kic/bustimer/usecase"
)

type ApproachInfoFromTimeTableController struct{}

func (ctrl ApproachInfoFromTimeTableController) FindApproachInfoFromTimeTable(timetable model.TimeTable, busstop string, via string) model.ApproachInfos {
	approachInfoFromTimeTableRepository := repositoryimpl.NewApproachInfoRepositoryImpl()
	result := usecase.NewGetApproachInfoFromTimeTableUseCaseImpl(timetable, busstop, via, approachInfoFromTimeTableRepository).FindApproachInfoFromTimeTable()
	return result
}

package controller

import (
	"fmt"
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type ApproachInfoFromTimeTableController struct{}

func (ctrl ApproachInfoFromTimeTableController) FindApproachInfoFromTimeTable(timetable model.TimeTable, busstop string, via string) model.ApproachInfos {
	approachInfoFromTimeTableRepository := repositoryimpl.NewApproachInfoRepositoryImpl()
	fmt.Printf("%v",timetable)
	result := usecase.NewGetApproachInfoFromTimeTableUseCaseImpl(timetable,busstop,via,approachInfoFromTimeTableRepository).FindApproachInfoFromTimeTable()
	return result
}

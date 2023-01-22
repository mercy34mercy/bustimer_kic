package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type getApproachInfoFromTimetableUseCaseImpl struct {
	TimeTable              model.TimeTable
	Bustop                 string
	Via                    string
	ApproachInforepository repository.ApproachInfoRepository
}

type getApproachInfoFromTimetableUseCase interface {
	FindApproachInfoFromTimeTable() model.ApproachInfos
}

func NewGetApproachInfoFromTimeTableUseCaseImpl(timetable model.TimeTable, busstop string, via string,ApproachInforepository repository.ApproachInfoRepository) getApproachInfoFromTimetableUseCase {
	return getApproachInfoFromTimetableUseCaseImpl{
		TimeTable: timetable,
		Bustop:    busstop,
		Via:       via,
		ApproachInforepository: ApproachInforepository,
	}
}

func (impl getApproachInfoFromTimetableUseCaseImpl) FindApproachInfoFromTimeTable() model.ApproachInfos {
	approachInfo := impl.ApproachInforepository.FindApproachInfoFromTimeTable(impl.TimeTable,impl.Bustop,impl.Via)
	return approachInfo
}

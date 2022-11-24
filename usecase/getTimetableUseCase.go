package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type getTimetableUseCaseImpl struct {
	URL []string
	Timetablerepository repository.BusstopToTimetableRepository
}

type getTimetableUseCase interface {
	FindTimetable()(model.TimeTable)
}

func NewGetTimetableUseCaseImpl(URL []string,Timetablereposiotry repository.BusstopToTimetableRepository) getTimetableUseCase{
	return getTimetableUseCaseImpl{
		URL:URL,
		Timetablerepository: Timetablereposiotry,
	}
}

func (impl getTimetableUseCaseImpl)FindTimetable()(model.TimeTable){
	timetable := impl.Timetablerepository.FindTimetable(impl.URL)
	return timetable
}
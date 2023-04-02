package usecase

import (
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/repository"
)

type getTimetableUseCaseImpl struct {
	Busstop             string
	Destination         string
	Timetablerepository repository.BusstopToTimetableRepository
}

type getTimetableUseCase interface {
	FindTimetable() (model.TimeTable, error)
}

func NewGetTimetableUseCaseImpl(bussop string, destination string, Timetablereposiotry repository.BusstopToTimetableRepository) getTimetableUseCase {
	return getTimetableUseCaseImpl{
		Busstop:             bussop,
		Destination:         destination,
		Timetablerepository: Timetablereposiotry,
	}
}

func (impl getTimetableUseCaseImpl) FindTimetable() (model.TimeTable, error) {
	var timetable model.TimeTable
	url, err := impl.Timetablerepository.FindURLFromBusstop(impl.Busstop, impl.Destination)
	if err != nil {
		return timetable, err
	}
	timetable, err = impl.Timetablerepository.FindTimetableParallel(url)
	return timetable, err
}

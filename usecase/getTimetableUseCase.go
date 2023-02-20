package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type getTimetableUseCaseImpl struct {
	Busstop string
	Destination string
	Timetablerepository repository.BusstopToTimetableRepository
}

type getTimetableUseCase interface {
	FindTimetable()(model.TimeTable,error)
}

func NewGetTimetableUseCaseImpl(bussop string,destination string,Timetablereposiotry repository.BusstopToTimetableRepository) getTimetableUseCase{
	return getTimetableUseCaseImpl{
		Busstop:bussop,
		Destination: destination,
		Timetablerepository: Timetablereposiotry,
	}
}

func (impl getTimetableUseCaseImpl)FindTimetable()(model.TimeTable,error){
	var timetable model.TimeTable
	url,err := impl.Timetablerepository.FindURLFromBusstop(impl.Busstop,impl.Destination)
	if err != nil{
		return timetable,err
	}
	timetable,err = impl.Timetablerepository.FindTimetable(url)
	return timetable,err
}
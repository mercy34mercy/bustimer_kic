package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type getUrlFromBusstopUseCaseImpl struct {
	Busstop string
	Destination []string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type getUrlFromBusstopUseCase interface {
	FindURLFromBusstop() (model.MultiTimeTable)
}

func NewGetUrlFromBusstopUseCaseImpl(busstop string,destination []string,busstoptotimetableRepository repository.BusstopToTimetableRepository) getUrlFromBusstopUseCase {
	return getUrlFromBusstopUseCaseImpl{
		Busstop: busstop,
		Destination: destination,
		BusstopToTimetableRepository: busstoptotimetableRepository,
	}
}

func (impl getUrlFromBusstopUseCaseImpl) FindURLFromBusstop()(model.MultiTimeTable) {
	timetableanddestination := []model.TimeTableandDestination{}
	for _,des := range impl.Destination{
		url := impl.BusstopToTimetableRepository.FindURLFromBusstop(impl.Busstop,des)
		timetable := impl.BusstopToTimetableRepository.FindTimetable(url)
		timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
			TimeTable: timetable,
			Destination: des,
		})
	}

	multitimetable := impl.BusstopToTimetableRepository.CreateMultiTimetable(timetableanddestination,impl.Destination)

	return multitimetable
}
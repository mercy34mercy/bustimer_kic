package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type createCacheUseCaseImpl struct {
	Businfo                      model.Busstop
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type createCacheUseCase interface {
	FindCacheTimetable() (model.TimeTable, string, string)
}

func NewCreateCacheUseCaseImpl(businfo model.Busstop, busstoptotimetableRepository repository.BusstopToTimetableRepository) createCacheUseCase {
	return createCacheUseCaseImpl{
		Businfo:                      businfo,
		BusstopToTimetableRepository: busstoptotimetableRepository,
	}
}

func (impl createCacheUseCaseImpl) FindCacheTimetable() (model.TimeTable, string, string) {
	destination := impl.BusstopToTimetableRepository.EncodeDestination(impl.Businfo.Busstop, impl.Businfo.Destination)
	url, err := impl.BusstopToTimetableRepository.FindURL(impl.Businfo.Busstop, destination)
	if err != nil {
	}
	
	timetable := impl.BusstopToTimetableRepository.FindTimetable(url)
	return timetable, impl.Businfo.Busstop, destination
}


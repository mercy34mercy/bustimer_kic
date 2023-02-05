package usecase

import (
	"fmt"
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
	"practice-colly/infra/localcache"
)

type getUrlFromBusstopUseCaseImpl struct {
	Busstop                      []string
	Destination                  []string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type getUrlFromBusstopUseCase interface {
	FindURLFromBusstop() model.MultiTimeTable
}

func NewGetUrlFromBusstopUseCaseImpl(busstop []string, destination []string, busstoptotimetableRepository repository.BusstopToTimetableRepository) getUrlFromBusstopUseCase {
	return getUrlFromBusstopUseCaseImpl{
		Busstop:                      busstop,
		Destination:                  destination,
		BusstopToTimetableRepository: busstoptotimetableRepository,
	}
}

func (impl getUrlFromBusstopUseCaseImpl) FindURLFromBusstop() model.MultiTimeTable {
	l := localcache.GetGoChache()
	timetableanddestination := []model.TimeTableandDestination{}
	if len(impl.Busstop) == 1 {
		for _, des := range impl.Destination {
			if x, found := l.Get(impl.Busstop[0] + des); found {
				fmt.Println("cache exist")
				var timetable model.TimeTable = x.(model.TimeTable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: des,
				})
			} else {
				var url []string
				url = impl.BusstopToTimetableRepository.FindURLFromBusstop(impl.Busstop[0], des)
				timetable := impl.BusstopToTimetableRepository.FindTimetable(url)
				localcache.CreateCachefromTimetable(impl.Busstop[0]+"行き", des, timetable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: des,
				})
			}
		}

	} else {
		for _, bus := range impl.Busstop {
			if x, found := l.Get(bus + impl.Destination[0]); found {
				fmt.Println("cache exist")
				var timetable model.TimeTable = x.(model.TimeTable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: bus,
				})
			} else {
				var url []string
				url, _ = impl.BusstopToTimetableRepository.FindURL(bus, impl.Destination[0]+"行き")
				timetable := impl.BusstopToTimetableRepository.FindTimetable(url)
				localcache.CreateCachefromTimetable(bus, impl.Destination[0]+"行き", timetable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: bus,
				})
			}
		}
	}

	multitimetable := impl.BusstopToTimetableRepository.CreateMultiTimetable(timetableanddestination, impl.Destination)

	return multitimetable
}

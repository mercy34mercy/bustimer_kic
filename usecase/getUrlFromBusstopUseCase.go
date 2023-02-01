package usecase

import (
	"fmt"
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
	"practice-colly/infra/localcache"
)

type getUrlFromBusstopUseCaseImpl struct {
	Busstop                      string
	Destination                  []string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type getUrlFromBusstopUseCase interface {
	FindURLFromBusstop() model.MultiTimeTable
}

func NewGetUrlFromBusstopUseCaseImpl(busstop string, destination []string, busstoptotimetableRepository repository.BusstopToTimetableRepository) getUrlFromBusstopUseCase {
	return getUrlFromBusstopUseCaseImpl{
		Busstop:                      busstop,
		Destination:                  destination,
		BusstopToTimetableRepository: busstoptotimetableRepository,
	}
}

func (impl getUrlFromBusstopUseCaseImpl) FindURLFromBusstop() model.MultiTimeTable {
	l := localcache.GetGoChache()
	timetableanddestination := []model.TimeTableandDestination{}
	for _, des := range impl.Destination {
		if x, found := l.Get(impl.Busstop + des); found {
			fmt.Println("cache exist")
			var timetable model.TimeTable = x.(model.TimeTable)
			timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
				TimeTable:   timetable,
				Destination: des,
			})
		} else {
			var url []string
			if des == "立命館大学" {
				url,_ = impl.BusstopToTimetableRepository.FindURL(impl.Busstop,des + "行き")
			} else {
				url = impl.BusstopToTimetableRepository.FindURLFromBusstop(impl.Busstop, des)
			}
			timetable := impl.BusstopToTimetableRepository.FindTimetable(url)
			localcache.CreateCachefromTimetable(impl.Busstop, des, timetable)
			timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
				TimeTable:   timetable,
				Destination: des,
			})
		}
	}

	multitimetable := impl.BusstopToTimetableRepository.CreateMultiTimetable(timetableanddestination, impl.Destination)

	return multitimetable
}

package usecase

import (
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/repository"
	"github.com/mercy34mercy/bustimer_kic/bustimer/infra/localcache"
	redisclient "github.com/mercy34mercy/bustimer_kic/bustimer/redis"
)

type getUrlFromBusstopUseCaseImpl struct {
	Busstop                      []string
	Destination                  []string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type getUrlFromBusstopUseCase interface {
	FindURLFromBusstop() (model.MultiTimeTable, error)
}

func NewGetUrlFromBusstopUseCaseImpl(busstop []string, destination []string, busstoptotimetableRepository repository.BusstopToTimetableRepository) getUrlFromBusstopUseCase {
	return getUrlFromBusstopUseCaseImpl{
		Busstop:                      busstop,
		Destination:                  destination,
		BusstopToTimetableRepository: busstoptotimetableRepository,
	}
}

func (impl getUrlFromBusstopUseCaseImpl) FindURLFromBusstop() (model.MultiTimeTable, error) {
	l := localcache.GetGoChache()
	var multitimetable model.MultiTimeTable
	timetableanddestination := []model.TimeTableandDestination{}
	if impl.Busstop[0] == "立命館大学前" {

		for _, des := range impl.Destination {
			t,err :=  redisclient.Get(impl.Busstop[0] + des + "行き")
			if err != nil {
				var timetable = t.(model.TimeTable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   t.(model.TimeTable)
				}
			}
			if x, found := l.Get(impl.Busstop[0] + des + "行き"); found {
				var timetable model.TimeTable = x.(model.TimeTable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: des,
				})
			} else {
				url, err := impl.BusstopToTimetableRepository.FindURLFromBusstop(impl.Busstop[0], des)
				if err != nil {
					return multitimetable, err
				}
				timetable, err := impl.BusstopToTimetableRepository.FindTimetableParallel(url)
				if err != nil {
					return multitimetable, err
				}
				localcache.CreateCachefromTimetable(impl.Busstop[0], des+"行き", timetable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: des,
				})
				multitimetable = impl.BusstopToTimetableRepository.CreateMultiTimetable(timetableanddestination, impl.Destination)
				return multitimetable, err
			}
		}
	} else {
		for _, bus := range impl.Busstop {
			if x, found := l.Get(bus + impl.Destination[0] + "行き"); found {
				var timetable model.TimeTable = x.(model.TimeTable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: bus,
				})
			} else {
				url, err := impl.BusstopToTimetableRepository.FindURLFromBusstop(bus, impl.Destination[0])
				if err != nil {
					return multitimetable, err
				}
				timetable, err := impl.BusstopToTimetableRepository.FindTimetableParallel(url)
				if err != nil {
					return multitimetable, err
				}
				localcache.CreateCachefromTimetable(bus, impl.Destination[0]+"行き", timetable)
				timetableanddestination = append(timetableanddestination, model.TimeTableandDestination{
					TimeTable:   timetable,
					Destination: bus,
				})
				multitimetable = impl.BusstopToTimetableRepository.CreateMultiTimetable(timetableanddestination, impl.Destination)
				return multitimetable, err
			}
		}
	}
	multitimetable = impl.BusstopToTimetableRepository.CreateMultiTimetable(timetableanddestination, impl.Destination)
	var err error
	return multitimetable, err
}
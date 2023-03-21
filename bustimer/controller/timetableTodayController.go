package controller

import (
	"time"

	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	repositoryImpl "github.com/mercy34mercy/bustimer_kic/bustimer/infra/repositoryImpl"
	"github.com/mercy34mercy/bustimer_kic/bustimer/usecase"
)

type TimetableTodayController struct{}

func (ctrl TimetableTodayController) Get(busstop []string, destination []string) (model.MultiTimeTable, error) {
	busstoptotimetableRepository := repositoryImpl.NewBusstopToUrlRepositoryImpl()
	timetable, err := usecase.NewGetUrlFromBusstopUseCaseImpl(busstop, destination, busstoptotimetableRepository).FindURLFromBusstop()

	day := time.Now().Weekday()
	if day == time.Saturday || day == time.Sunday {
		for _, v := range timetable.TimeTable {
			for i := 5; i <= 24; i++ {
				v.Weekdays[i] = make([]model.OneBusTime, 0)
			}
		}
	} else {
		for _, v := range timetable.TimeTable {
			for i := 5; i <= 24; i++ {
				v.Holidays[i] = make([]model.OneBusTime, 0)
			}
		}
	}
	return timetable, err
}

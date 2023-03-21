package controller

import (
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	repositoryImpl "github.com/mercy34mercy/bustimer_kic/bustimer/infra/repositoryImpl"
	"github.com/mercy34mercy/bustimer_kic/bustimer/usecase"
)

type TimetableFromBusstopController struct{}

func (ctrl TimetableFromBusstopController) FindTimetable(busstop []string, destination []string) (model.MultiTimeTable, error) {
	busstoptotimetableRepository := repositoryImpl.NewBusstopToUrlRepositoryImpl()
	timetable, err := usecase.NewGetUrlFromBusstopUseCaseImpl(busstop, destination, busstoptotimetableRepository).FindURLFromBusstop()
	return timetable, err
}

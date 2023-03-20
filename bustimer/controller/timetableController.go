package controller

import (
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	repositoryimpl "github.com/mercy34mercy/bustimer_kic/bustimer/infra/repositoryImpl"
	"github.com/mercy34mercy/bustimer_kic/bustimer/usecase"
)

type TimetableController struct{}

func (ctrl TimetableController) FindTimetable(Busstop string, Destination string) (model.TimeTable, error) {
	busstoptourlRepository := repositoryimpl.NewBusstopToUrlRepositoryImpl()
	timetable, err := usecase.NewGetTimetableUseCaseImpl(Busstop, Destination, busstoptourlRepository).FindTimetable()
	return timetable, err
}

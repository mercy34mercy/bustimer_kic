package repository

import (
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	bustimersqlc "github.com/mercy34mercy/bustimer_kic/bustimer/sqlc/gen"
)

type BusstopToTimetableRepository interface {
	FindURL(busstop string, destination string) ([]string, error)
	FindTimetable(url []string) (model.TimeTable, error)
	FindBusstopList(busname string) ([]bustimersqlc.GetBusstopAndDestinationRow, error)
	EncodeDestination(destination string) (wrapdestination string)
	FindURLFromBusstop(busstop string, destination string) ([]string, error)
	CreateMultiTimetable(timetable []model.TimeTableandDestination, destinationlist []string) model.MultiTimeTable
}

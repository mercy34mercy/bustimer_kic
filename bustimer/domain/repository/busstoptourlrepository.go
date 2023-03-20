package repository

import (
	"bustimerkic/domain/model"
	bustimersqlc "bustimerkic/sqlc/gen"
)

type BusstopToTimetableRepository interface {
	FindURL(busstop string,destination string)([]string,error)
	FindTimetable(url []string)(model.TimeTable,error)
	FindBusstopList(busname string)([]bustimersqlc.GetBusstopAndDestinationRow,error)
	EncodeDestination(destination string)(wrapdestination string)
	FindURLFromBusstop(busstop string,destination string)([]string,error)
	CreateMultiTimetable(timetable []model.TimeTableandDestination,destinationlist []string)(model.MultiTimeTable)
}
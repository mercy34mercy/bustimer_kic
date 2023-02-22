package repository

import "bustimerkic/domain/model"

type BusstopToTimetableRepository interface {
	FindURL(busstop string,destination string)([]string,error)
	FindTimetable(url []string)(model.TimeTable,error)
	FindBusstopList(busname string)([]model.Busstop,error)
	EncodeDestination(destination string)(wrapdestination string)
	FindURLFromBusstop(busstop string,destination string)([]string,error)
	CreateMultiTimetable(timetable []model.TimeTableandDestination,destinationlist []string)(model.MultiTimeTable)
}
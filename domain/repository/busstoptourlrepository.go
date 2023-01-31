package repository

import "practice-colly/domain/model"

type BusstopToTimetableRepository interface {
	FindURL(busstop string,destination string)([]string,error)
	FindTimetable(url []string)(model.TimeTable)
	FindBusstopList(busname string)([]model.Busstop,error)
	EncodeDestination(busstop string,destination string)(wrapdestination string)
	FindURLFromBusstop(busstop string,destination string)([]string)
	CreateMultiTimetable(timetable []model.TimeTableandDestination,destinationlist []string)(model.MultiTimeTable)
}
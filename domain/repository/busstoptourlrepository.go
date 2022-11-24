package repository

import "practice-colly/domain/model"

type BusstopToTimetableRepository interface {
	FindURL(busstop string,destination string)([]string,error)
	FindTimetable(url []string)(model.TimeTable)
}
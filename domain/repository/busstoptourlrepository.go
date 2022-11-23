package repository

type BusstopToTimetableRepository interface {
	FindURL(busstop string,destination string,busname string)(string,error)
	FindHolidaysTimetable(url string)()
	FindSaturdaysTimetable(url string)()
	FindWeekendsTimetable(url string)()
}
package model

import (
	"regexp"
	"sort"
	"strconv"
)

type TimeTable struct {
	Weekdays 	map[int][]OneBusTime `json:"weekdays"`
	Saturdays 	map[int][]OneBusTime `json:"saturdays"`
	Holidays 	map[int][]OneBusTime `json:"holidays"`
}

type OneBusTime struct {
	Via 	string `json:"via"`
	Min 	string `json:"min"`
	BusStop string `json:"bus_stop"`
}


func CreateNewTimeTable() TimeTable {
	// 初期化
	timetable := TimeTable{
		Weekdays: make(map[int][]OneBusTime),
		Saturdays: make(map[int][]OneBusTime),
		Holidays: make(map[int][]OneBusTime),
	}

	// 時刻表にあるデータを埋める
	for i := 5; i <= 24; i++ {
		timetable.Weekdays[i] = make([]OneBusTime, 0)
		timetable.Saturdays[i] = make([]OneBusTime, 0)
		timetable.Holidays[i] = make([]OneBusTime, 0)
	}
	return timetable
}

func (timetable TimeTable) SortOneBusTime() {
	for _,oneBusTimeHolidayList := range timetable.Holidays{
		sortByMin(oneBusTimeHolidayList)
	}

	for _,oneBusTImeSaturdayList := range timetable.Saturdays{
		sortByMin(oneBusTImeSaturdayList)
	}

	for _,oneBusTimeWeekdaysList := range timetable.Weekdays{
		sortByMin(oneBusTimeWeekdaysList)
	}
}


func sortByMin(oneBusTimeList []OneBusTime) {
	rex := regexp.MustCompile("[0-9]+")	
	sort.Slice(oneBusTimeList, func(i, j int) bool {
		iMin, _ := strconv.Atoi(rex.FindString(oneBusTimeList[i].Min))
		jMin, _ := strconv.Atoi(rex.FindString(oneBusTimeList[j].Min) )
		return iMin < jMin
	})
}
package model

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type MultiTimeTable struct {
	TimeTable map[string]TimeTable `json:"timetable"`
}

type TimeTableandDestination struct {
	TimeTable   TimeTable
	Destination string
}

type TimeTable struct {
	Weekdays map[int][]OneBusTime `json:"weekdays"`
	Holidays map[int][]OneBusTime `json:"holidays"`
}

type OneBusTime struct {
	BusName string `json:"bus_name"`
	Min     string `json:"min"`
	BusStop string `json:"bus_stop"`
}

type Busstop struct {
	Busstop     string
	Destination string
}

type Query struct {
	Destination string
	Busstop     string
}

func CreateNewMultiTimeTable(destinationList []string) MultiTimeTable {
	multiTimeTable := MultiTimeTable{
		TimeTable: make(map[string]TimeTable),
	}

	return multiTimeTable
}

func (query Query) SplitDestination() ([]string, []string) {
	destinationlist := strings.Split(query.Destination, ",")
	busstopolist := strings.Split(query.Busstop, ",")
	return destinationlist, busstopolist
}

func CreateNewTimeTable() TimeTable {
	// 初期化
	timetable := TimeTable{
		Weekdays: make(map[int][]OneBusTime),
		Holidays: make(map[int][]OneBusTime),
	}

	// 時刻表にあるデータを埋める
	for i := 5; i <= 24; i++ {
		timetable.Weekdays[i] = make([]OneBusTime, 0)
		timetable.Holidays[i] = make([]OneBusTime, 0)
	}
	return timetable
}

func (timetable TimeTable) SortOneBusTime() {
	for _, oneBusTimeHolidayList := range timetable.Holidays {
		sortByMin(oneBusTimeHolidayList)
	}

	for _, oneBusTimeWeekdaysList := range timetable.Weekdays {
		sortByMin(oneBusTimeWeekdaysList)
	}
}

func sortByMin(oneBusTimeList []OneBusTime) {
	rex := regexp.MustCompile("[0-9]+")
	sort.Slice(oneBusTimeList, func(i, j int) bool {
		iMin, _ := strconv.Atoi(rex.FindString(oneBusTimeList[i].Min))
		jMin, _ := strconv.Atoi(rex.FindString(oneBusTimeList[j].Min))
		return iMin < jMin
	})
}

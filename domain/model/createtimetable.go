package model

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
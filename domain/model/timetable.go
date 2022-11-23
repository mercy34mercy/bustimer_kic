package model

type TimeTable struct {
	Busnumber int `json:"busnumber"`
	Weekdays 	map[int][]OneBusTime `json:"weekdays"`
	Saturdays 	map[int][]OneBusTime `json:"saturdays"`
	Holidays 	map[int][]OneBusTime `json:"holidays"`
}

type OneBusTime struct {
	Via 	string `json:"via"`
	Min 	string `json:"min"`
	BusStop string `json:"bus_stop"`
}
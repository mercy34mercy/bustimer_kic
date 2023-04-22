package domain

type TimeTable struct {
	Weekdays map[int][]OneBusTime `json:"weekdays"`
	Holidays map[int][]OneBusTime `json:"holidays"`
}

type OneBusTime struct {
	BusName string `json:"bus_name"`
	Min     string `json:"min"`
	BusStop string `json:"bus_stop"`
}

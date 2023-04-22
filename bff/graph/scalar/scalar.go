package scalar

import "io"

type Timetable struct {
	Weekdays map[int]OneBusTime
	Holidays map[int]OneBusTime
}

type OneBusTime struct {
	BusName string `json:"bus_name"`
	Min     string `json:"min"`
	BusStop string `json:"bus_stop"`
}

func (t *Timetable) UnmarshalGQL(v interface{}) error {
	panic("implement me")
}

func (t Timetable) MarshalGQL(w io.Writer) {

}

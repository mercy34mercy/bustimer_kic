package dto

import (
	"github.com/mercy34mercy/bustimer_kic/bff/domain"
	"github.com/mercy34mercy/bustimer_kic/bff/graph/model"
)

// domain.TimeTableをmodel.TimeTableに変換する
func TimetableDto(timetable domain.TimeTable) model.TimeTable {
	var Timetable model.TimeTable
	for k, v := range timetable.Weekdays {
		var oneBusTime []*model.OneBusTime
		for _, bus := range v {
			oneBusTime = append(oneBusTime, &model.OneBusTime{
				BusName: bus.BusName,
				Min:     bus.Min,
				BusStop: bus.BusStop,
			})
		}
		Timetable.Weekdays = append(Timetable.Weekdays, &model.HourBusTime{
			Bus:  oneBusTime,
			Hour: k,
		})
	}

	for k, v := range timetable.Holidays {
		var oneBusTime []*model.OneBusTime
		for _, bus := range v {
			oneBusTime = append(oneBusTime, &model.OneBusTime{
				BusName: bus.BusName,
				Min:     bus.Min,
				BusStop: bus.BusStop,
			})
		}
		Timetable.Holidays = append(Timetable.Holidays, &model.HourBusTime{
			Bus:  oneBusTime,
			Hour: k,
		})
	}

	return Timetable
}

// domain.ApproachInfosをmodel.ApproachInfosに変換する
func ApproachInfosDto(approachInfos domain.ApproachInfos) model.ApproachInfos {
	var ApproachInfos model.ApproachInfos
	for _, v := range approachInfos.ApproachInfo {
		ApproachInfos.ApproachInfo = append(ApproachInfos.ApproachInfo, &model.ApproachInfo{
			MoreMin:        v.MoreMin,
			RealARivalTime: v.RealARivalTime,
			Direction:      v.Direction,
			ScheduledTime:  v.ScheduledTime,
			Delay:          v.Delay,
			BusStop:        v.BusStop,
			BusName:        v.BusName,
			RequiredTime:   v.RequiredTime,
		})
	}
	return ApproachInfos
}

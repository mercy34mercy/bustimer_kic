package repositoryimpl

import (
	"practice-colly/config"
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
	"strconv"
	"time"
)

type ApproachInfoRepositoryImpl struct{}

func NewApproachInfoRepositoryImpl() repository.ApproachInfoRepository {
	return &ApproachInfoRepositoryImpl{}
}

func (repository *ApproachInfoRepositoryImpl) FindApproachInfoFromTimeTable(timetable model.TimeTable, via string, busstop string) model.ApproachInfos {
	approachinfo := model.CreateApproachInfos()

	approachinfo = getApproachInfoFromTimetable(approachinfo, timetable, via, busstop)

	fastThree := approachinfo.GetFastThree()

	return fastThree
}


func getApproachInfoFromTimetable(approachInfos model.ApproachInfos, timeTable model.TimeTable, via string, busstop string) model.ApproachInfos {
	Via := via
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	today := time.Now().In(jst)
	weekday := today.Weekday()
	if weekday == 6 {
		for hour, times := range timeTable.Saturdays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if (hour == today.Hour() && min > today.Minute() || hour > today.Hour()) && (hour-2 < today.Hour()) {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						RealArrivalTime:formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)) ,
						Direction:       busstop,
						BusName: 		 time.Via,	
						ScheduledTime:   formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    config.GetRequiredeTime(Via,busstop,time.Via),
					})
				}
			}
		}
	} else if weekday == 1 {
		for hour, times := range timeTable.Weekdays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if (hour == today.Hour() && min > today.Minute() || hour > today.Hour()) && (hour-2 < today.Hour())  {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						RealArrivalTime:formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)) ,
						Direction:       busstop,
						BusName: 		 time.Via,		
						ScheduledTime:   formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    config.GetRequiredeTime(Via,busstop,time.Via),
					})
				}
			}
		}
	} else {
		for hour, times := range timeTable.Holidays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if (hour == today.Hour() && min > today.Minute() || hour > today.Hour()) && (hour-2 < today.Hour())  {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						RealArrivalTime:formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)) ,
						Direction:       busstop,
						BusName: 		 time.Via,	
						ScheduledTime:   formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    config.GetRequiredeTime(Via,busstop,time.Via),
					})
				}
			}
		}
	}
	return approachInfos
}

func formatMin(min string) string{
	var mmMin = min
	if len(min) == 1 {
		mmMin = "0" + min
	}
	return mmMin
}

func formatHour(hour string) string{
	var mmHour = hour
	if len(hour) == 1 {
		mmHour = "0" + hour
	}
	return mmHour
}
package repositoryimpl

import (
	"os"
	"strconv"
	"time"

	"github.com/mercy34mercy/bustimer_kic/bustimer/config"
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/repository"
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
	if weekday == 6 || weekday == 0 {
		for hour, times := range timeTable.Holidays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if (hour == today.Hour() && min > today.Minute() || hour > today.Hour()) && (hour-2 < today.Hour()) {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						RealArrivalTime: formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						MoreMin:         "約n分後に到着",
						Direction:       busstop,
						BusName:         time.BusName,
						ScheduledTime:   formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    config.GetRequiredeTime(Via, busstop, time.BusName),
					})
				}
			}
		}
	} else {
		for hour, times := range timeTable.Weekdays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if (hour == today.Hour() && min > today.Minute() || hour > today.Hour()) && (hour-2 < today.Hour()) {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						RealArrivalTime: formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						MoreMin:         "約n分後に到着",
						Direction:       busstop,
						BusName:         time.BusName,
						ScheduledTime:   formatHour(strconv.FormatInt(int64(hour), 10)) + ":" + formatMin(strconv.FormatInt(toInt64(time.Min), 10)),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    config.GetRequiredeTime(Via, busstop, time.BusName),
					})
				}
			}
		}
	}

	if (len(approachInfos.ApproachInfo) < 1 && os.Getenv("GO_ENV") == "dev"){
		return getMockdata(today)
	}

	return approachInfos
}

func formatMin(min string) string {
	var mmMin = min
	if len(min) == 1 {
		mmMin = "0" + min
	}
	return mmMin
}

func formatHour(hour string) string {
	var mmHour = hour
	if len(hour) == 1 {
		mmHour = "0" + hour
	}
	return mmHour
}

func getMockdata(t time.Time) model.ApproachInfos {
	t1 := t.Add(10 * time.Minute)
	t2 := t.Add(15 * time.Minute)
	t3 := t.Add(20 * time.Minute)
	approachinfo := []model.ApproachInfo{
		{
			RealArrivalTime: formatHour(strconv.FormatInt(int64(t1.Hour()), 10)) + ":" + formatMin(strconv.FormatInt(int64(t1.Minute()), 10)),
			MoreMin:         "約n分後に到着",
			Direction:       "mock1停留所",
			BusName:         "mock1号系統",
			ScheduledTime:   formatHour(strconv.FormatInt(int64(t1.Hour()), 10)) + ":" + formatMin(strconv.FormatInt(int64(t1.Minute()), 10)),
			Delay:           "定時運行",
			BusStop:         "1",
			RequiredTime:    10,
		},
		{
			RealArrivalTime: formatHour(strconv.FormatInt(int64(t2.Hour()), 10)) + ":" + formatMin(strconv.FormatInt(int64(t2.Minute()), 10)),
			MoreMin:         "約n分後に到着",
			Direction:       "mock2停留所",
			BusName:         "mock2号系統",
			ScheduledTime:   formatHour(strconv.FormatInt(int64(t2.Hour()), 10)) + ":" + formatMin(strconv.FormatInt(int64(t2.Minute()), 10)),
			Delay:           "定時運行",
			BusStop:         "1",
			RequiredTime:    10,
		},
		{
			RealArrivalTime: formatHour(strconv.FormatInt(int64(t3.Hour()), 10)) + ":" + formatMin(strconv.FormatInt(int64(t3.Minute()), 10)),
			MoreMin:         "約n分後に到着",
			Direction:       "mock3停留所",
			BusName:         "mock3号系統",
			ScheduledTime:   formatHour(strconv.FormatInt(int64(t3.Hour()), 10)) + ":" + formatMin(strconv.FormatInt(int64(t3.Minute()), 10)),
			Delay:           "定時運行",
			BusStop:         "1",
			RequiredTime:    10,
		},
	}

	approachinfos := model.ApproachInfos{
		ApproachInfo: approachinfo,
	}
	return approachinfos
}

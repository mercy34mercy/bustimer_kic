package repositoryimpl

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
	"strconv"
	"strings"
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
func (repository *ApproachInfoRepositoryImpl) FindApproachInfo(url []string) model.ApproachInfos {
	approachinfo := model.CreateApproachInfos()
	for _, u := range url {
		scrapdata, via, busstop := scrapHTML(u)
		approachinfo = getApproachInfo(approachinfo, scrapdata, via, busstop)
	}

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
				if hour == today.Hour() && min > today.Minute() || hour > today.Hour() {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						MoreMin:         "約n分後に到着",
						RealArrivalTime: strconv.FormatInt(int64(hour), 10) + ":" + strconv.FormatInt(toInt64(time.Min), 10),
						Direction:       Via,
						Via:             Via,
						ScheduledTime:   strconv.FormatInt(int64(hour), 10) + ":" + strconv.FormatInt(toInt64(time.Min), 10),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    20,
					})
				}
			}
		}
	} else if weekday == 1 {
		for hour, times := range timeTable.Weekdays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if hour == today.Hour() && min > today.Minute() || hour > today.Hour() {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						MoreMin:         "約n分後に到着",
						RealArrivalTime: strconv.FormatInt(int64(hour), 10) + ":" + strconv.FormatInt(toInt64(time.Min), 10),
						Direction:       Via,
						Via:             Via,
						ScheduledTime:   strconv.FormatInt(int64(hour), 10) + ":" + strconv.FormatInt(toInt64(time.Min), 10),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    20,
					})
				}
			}
		}
	} else {
		for hour, times := range timeTable.Holidays {
			for _, time := range times {
				min, _ := strconv.Atoi(time.Min)
				if hour == today.Hour() && min > today.Minute() || hour > today.Hour() {
					approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
						MoreMin:         "約n分後に到着",
						RealArrivalTime: strconv.FormatInt(int64(hour), 10) + ":" + strconv.FormatInt(toInt64(time.Min), 10),
						Direction:       Via,
						Via:             Via,
						ScheduledTime:   strconv.FormatInt(int64(hour), 10) + ":" + strconv.FormatInt(toInt64(time.Min), 10),
						Delay:           "定時運行",
						BusStop:         "1",
						RequiredTime:    20,
					})
				}
			}
		}
	}
	return approachInfos
}

func getApproachInfo(approachInfos model.ApproachInfos, scrapedata []string, via string, busstop string) model.ApproachInfos {

	Via := via
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	today := time.Now().In(jst)
	weekday := today.Weekday()

	//土曜日
	if weekday == 6 {
		//日曜日
		for i := 0; i < len(scrapedata)/3; i++ {
			holidaylist := strings.Split(scrapedata[i*3+1], " ")
			for b := 0; b < len(holidaylist); b++ {
				min, _ := strconv.Atoi(holidaylist[b])
				if i+5 == today.Hour() && min > today.Minute() || i+5 > today.Hour() {
					hour := converthour(i + 5)
					if holidaylist[b] != "" {
						approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
							MoreMin:         "約n分後に到着",
							RealArrivalTime: hour + ":" + strconv.FormatInt(toInt64(holidaylist[b]), 10),
							Direction:       Via,
							Via:             Via,
							ScheduledTime:   hour + ":" + strconv.FormatInt(toInt64(holidaylist[b]), 10),
							Delay:           "定時運行",
							BusStop:         "1",
							RequiredTime:    20,
						})
					}
				}
			}
		}

	} else if weekday == 1 {
		for i := 0; i < len(scrapedata)/3; i++ {
			saturdaylist := strings.Split(scrapedata[i*3+2], " ")
			for a := 0; a < len(saturdaylist); a++ {
				min, _ := strconv.Atoi(saturdaylist[a])
				if i+5 == today.Hour() && min > today.Minute() || i+5 > today.Hour() {
					hour := converthour(i + 5)
					if saturdaylist[a] != "" {
						approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
							MoreMin:         "約n分後に到着",
							RealArrivalTime: hour + ":" + strconv.FormatInt(toInt64(saturdaylist[a]), 10),
							Direction:       busstop,
							Via:             Via,
							ScheduledTime:   hour + ":" + strconv.FormatInt(toInt64(saturdaylist[a]), 10),
							Delay:           "定時運行",
							BusStop:         "1",
							RequiredTime:    20,
						})
					}
				}
			}
		}
	} else {
		for i := 0; i < len(scrapedata)/3; i++ {
			weekdaylist := strings.Split(scrapedata[i*3], " ")
			for a := 0; a < len(weekdaylist); a++ {
				min, _ := strconv.Atoi(weekdaylist[a])
				if i+5 == today.Hour() && min > today.Minute() || i+5 > today.Hour() {
					hour := converthour(i + 5)
					if weekdaylist[a] != "" {
						approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, model.ApproachInfo{
							MoreMin:         "約n分後に到着",
							RealArrivalTime: hour + ":" + strconv.FormatInt(toInt64(weekdaylist[a]), 10),
							Direction:       busstop,
							Via:             Via,
							ScheduledTime:   hour + ":" + strconv.FormatInt(toInt64(weekdaylist[a]), 10),
							Delay:           "定時運行",
							BusStop:         "1",
							RequiredTime:    20,
						})
					}
				}
			}
		}
	}
	return approachInfos
}

func time2str(t time.Time) string {
	// レシーバーtを、"YYYY-MM-DDTHH-MM-SSZZZZ"という形の文字列に変換する
	return t.Format("2006-01-02T15:04:05Z07:00")
}

func converthour(h int) string {
	if h < 10 {
		return "0" + strconv.Itoa(h)
	} else {
		return strconv.Itoa(h)
	}
}

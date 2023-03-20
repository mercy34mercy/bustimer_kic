package repository

import "github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"

type ApproachInfoRepository interface {
	FindApproachInfoFromTimeTable(timetable model.TimeTable, via string, busstop string) model.ApproachInfos
}

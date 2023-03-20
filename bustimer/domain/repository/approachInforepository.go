package repository

import "bustimerkic/domain/model"

type ApproachInfoRepository interface{
	FindApproachInfoFromTimeTable(timetable model.TimeTable,via string,busstop string)(model.ApproachInfos)
}
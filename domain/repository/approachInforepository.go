package repository

import "practice-colly/domain/model"

type ApproachInfoRepository interface{
	FindApproachInfoFromTimeTable(timetable model.TimeTable,via string,busstop string)(model.ApproachInfos)
}
package repository

import "practice-colly/domain/model"

type ApproachInfoRepository interface{
	FindApproachInfo(url []string)(model.ApproachInfos)
}
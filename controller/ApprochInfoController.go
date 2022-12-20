package controller

import (
	"practice-colly/domain/model"
	repositoryimpl "practice-colly/infra/repositoryImpl"
	"practice-colly/usecase"
)

type ApproachInfoController struct{}

func (ctrl ApproachInfoController) FindApproachInfo(url []string)(model.ApproachInfos) {
	approachInfoRepository := repositoryimpl.NewApproachInfoRepositoryImpl()
	result := usecase.NewGetApproachInfoUseCaseImpl(url,approachInfoRepository).FindApproachInfo()
	return result
}
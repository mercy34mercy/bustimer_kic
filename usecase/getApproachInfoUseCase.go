package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type getApproachInfoUseCaseImpl struct {
	URL []string
	ApproachInforepository repository.ApproachInfoRepository
}

type getApproachInfoUseCase interface{
	FindApproachInfo()(model.ApproachInfos)
}

func NewGetApproachInfoUseCaseImpl(url []string,ApproachInforepository repository.ApproachInfoRepository) getApproachInfoUseCase{
	return getApproachInfoUseCaseImpl{
		URL: url,
		ApproachInforepository: ApproachInforepository,
	}
}

func (impl getApproachInfoUseCaseImpl)FindApproachInfo()(model.ApproachInfos){
	approachInfo := impl.ApproachInforepository.FindApproachInfo(impl.URL)
	return approachInfo
}
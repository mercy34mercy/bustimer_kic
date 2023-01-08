package usecase

import (
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
)

type getBusstopListUseCaseImpl struct {
	Busname string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type getBusstopListUseCase interface{
	FindBusstopList()([]model.Busstop,error)
}

func NewGetBusstopListUseCaseImpl(Busname string,busstoptotimetablerepository repository.BusstopToTimetableRepository) getBusstopListUseCase {
	return getBusstopListUseCaseImpl{
		Busname: Busname,
		BusstopToTimetableRepository: busstoptotimetablerepository,
	}
}

func (impl getBusstopListUseCaseImpl) FindBusstopList()([]model.Busstop,error){
	busstoplist,err := impl.BusstopToTimetableRepository.FindBusstopList(impl.Busname)
	return busstoplist,err
}
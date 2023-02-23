package usecase

import (
	"bustimerkic/domain/repository"
	bustimersqlc "bustimerkic/sqlc/gen"
)

type getBusstopListUseCaseImpl struct {
	Busname string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository
}

type getBusstopListUseCase interface{
	FindBusstopList()([]bustimersqlc.GetBusstopAndDestinationRow,error)
}

func NewGetBusstopListUseCaseImpl(Busname string,busstoptotimetablerepository repository.BusstopToTimetableRepository) getBusstopListUseCase {
	return getBusstopListUseCaseImpl{
		Busname: Busname,
		BusstopToTimetableRepository: busstoptotimetablerepository,
	}
}

func (impl getBusstopListUseCaseImpl) FindBusstopList()([]bustimersqlc.GetBusstopAndDestinationRow,error){
	busstoplist,err := impl.BusstopToTimetableRepository.FindBusstopList(impl.Busname)
	return busstoplist,err
}
package usecase

import "bustimerkic/domain/repository"

type getBusstopUrlUseCaseImpl struct {
	Busstop string
	Destination string
	BusstopToTimetableRepository repository.BusstopToTimetableRepository

}

type getBusstopUrlUseCase interface {
	FindURL()([]string,error)
}

func NewGetBusstopUrlUseCaseImpl(Busstop string,Destination string,busstoptotimetablerepository repository.BusstopToTimetableRepository) getBusstopUrlUseCase {
	return getBusstopUrlUseCaseImpl{
		Busstop: Busstop,
		Destination: Destination,
		BusstopToTimetableRepository: busstoptotimetablerepository,
	}
}

func (impl getBusstopUrlUseCaseImpl) FindURL()([]string,error){
	url,err := impl.BusstopToTimetableRepository.FindURL(impl.Busstop,impl.Destination)
	return url,err
}


package repositoryimpl

import "practice-colly/domain/repository"

type BusstopToTimetableRepositoryImpl struct{}

func NewBusstopToUrlRepository() repository.BusstopToTimetableRepository{
	return &BusstopToTimetableRepositoryImpl{}
}

func (repository *BusstopToTimetableRepositoryImpl) FindURL(busstop string,destination string,busname string)(string,error){

}

func (repository *BusstopToTimetableRepositoryImpl) FindHolidaysTimetable(url string)(){
	
}
func (repository *BusstopToTimetableRepositoryImpl) FindSaturdaysTimetable(url string)(){

}
func (repository *BusstopToTimetableRepositoryImpl) FindWeekendsTimetable(url string)(){

}
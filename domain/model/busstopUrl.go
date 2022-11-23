package model

type BusstopUrl struct{
	ID int `gorm:"AUTO_INCREMENT"`
	Busstop string `gorm:"type:varchar(255)"`
	Busname string `gorm:"type:varchar(255)"`
	Destination string `gorm:"type:varchar(255)"`
	URL string `gorm:"type:varchar(255);unique"`
}
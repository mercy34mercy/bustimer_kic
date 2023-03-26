package repository

import (
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
	"github.com/mercy34mercy/bustimer_kic/store/firestore"
)

type busstopRepository struct {
	client     firestore.Database
	collection string
}

func NewBusstopRespository(client firestore.Database, collection string) domain.BusstopRepository {
	return &busstopRepository{
		client:     client,
		collection: collection,
	}
}

func (ur *busstopRepository) GetByID(c echo.Context, userID string) ([]domain.Busstop, error) {
	return
}

func (ur *busstopRepository) Fetch(c echo.Context, userID string, busstop string) error {
	return
}

func (ur *busstopRepository) Delete(c echo.Context, userID string, busstop string) error {
	return
}

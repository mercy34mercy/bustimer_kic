package repository

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
	"github.com/mercy34mercy/bustimer_kic/store/firestore"
)

type busstopRepository struct {
	database   firestore.Database
	collection string
}

func NewBusstopRespository(database firestore.Database, collection string) domain.BusstopRepository {
	return &busstopRepository{
		database:   database,
		collection: collection,
	}
}

func (ur *busstopRepository) GetByID(c echo.Context, userID string) ([]domain.Busstop, error) {
	collection := ur.database.Collection("busstop")
	busstops, err := collection.FindByID(c, userID)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return busstops, nil
}

func (ur *busstopRepository) Fetch(c echo.Context, userID string, busstop string) error {
	return
}

func (ur *busstopRepository) Delete(c echo.Context, userID string, busstop string) error {
	return
}

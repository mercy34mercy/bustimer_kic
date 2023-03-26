package firestore

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
)

type Database interface {
	Collection(string) Collection
	Client() firestore.Client
}

type Collection interface {
	Find(echo.Context, string, string) (domain.Busstop, error)
	Delete(echo.Context, string, string) error
	Fetch(echo.Context, string, domain.Busstop) error
}

type firebaseCollection struct {
	fc firestore.CollectionRef
}

func NewCollection() (*firestore.CollectionRef, error) {
	ctx := context.Background()
	err := godotenv.Load(fmt.Sprintf("../env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		return nil, fmt.Errorf("faild open envfile: %v", err)
	}

	opt := option.WithCredentialsFile("../credential/busdes-firestore.json")
	client, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_CODE"), opt)
	if err != nil {
		return nil, fmt.Errorf("error get data: %v", err)
	}

	collection := client.Collection("busstop")
	return collection, nil
}

func (fc *firebaseCollection) GetDoc(ctx context.Context, path string) (*firestore.DocumentSnapshot, error) {
	datasnap, err := fc.fc.Doc(path).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("do not find %s : %v", path, err)
	}
	return datasnap, nil
}

func (fc *firebaseCollection) Get(ctx context.Context, path string) ([]domain.Busstop, error) {
	var busStops []domain.Busstop
	datasnap, err := fc.fc.Doc(path).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("do not find %s : %v", path, err)
	}
	if err := datasnap.DataTo(&busStops); err != nil {
		return nil, fmt.Errorf("faild parse to bussstop : %v", err)
	}
	return busStops, nil
}

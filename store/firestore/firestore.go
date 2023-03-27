package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
	"google.golang.org/api/option"
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

type Client interface {
	Database(string) Database
}

type firestoreDoc struct {
	fd firestore.DocumentSnapshot
}

type firebaseCollection struct {
	fc firestore.CollectionRef
}

func NewClient(ctx context.Context, projectID string, opt option.ClientOption) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		return nil, fmt.Errorf("can not connect to firestore : %v", err)
	}
	return client, nil
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

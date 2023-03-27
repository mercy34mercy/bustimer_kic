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
	FindByID(echo.Context, string) ([]domain.Busstop, error)
	Delete(echo.Context, string, string) error
	Fetch(echo.Context, string, domain.Busstop) error
}

type Client interface {
	Database(string) Database
}

type firestoreDatabase struct {
	fd firestore.Client
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

func (fd *firestoreDatabase) Collection(ctx context.Context, path string) (*firestore.CollectionRef, error) {
	collectionRef := fd.fd.Collection(path)
	return collectionRef, nil
}

func (fc *firebaseCollection) FindByID(ctx echo.Context, userID string) ([]domain.Busstop, error) {
	datasnap, err := fc.fc.Doc(userID).Get(ctx.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("do not find %s : %v", userID, err)
	}
	var busStops []domain.Busstop
	if err := datasnap.DataTo(&busStops); err != nil {
		return nil, fmt.Errorf("faild parse to bussstop : %v", err)
	}
	return busStops, nil
}

package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/domain"
	"google.golang.org/api/option"
)

type Client interface {
	Collection(string) Collection
}

type Collection interface {
	FindByID(echo.Context, string) ([]domain.Busstop, error)
	Delete(echo.Context, string, domain.Busstop) error
	Fetch(echo.Context, string, domain.Busstop) error
}

type firestoreClient struct {
	db *firestore.Client
}

type firestoreCollection struct {
	fc *firestore.CollectionRef
}

func NewClient(ctx context.Context, projectID string, opt option.ClientOption) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		return nil, fmt.Errorf("can not connect to firestore : %v", err)
	}
	return client, nil
}

func (fc *firestoreClient) Collection(path string) (*firestore.CollectionRef, error) {
	collectionRef := fc.db.Collection(path)
	return collectionRef, nil
}

func (fc *firestoreCollection) Delete(ctx echo.Context, userID string, busstop domain.Busstop) error {
	//fc.fc.Doc(userID).Delete(ctx.Request().Context(), busstop)
	return nil
}

func (fc *firestoreCollection) Fetch(ctx echo.Context, userID string, busstop domain.Busstop) error {
	_, err := fc.fc.Doc(userID).Create(ctx.Request().Context(), busstop)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (fc *firestoreCollection) FindByID(ctx echo.Context, userID string) ([]domain.Busstop, error) {
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

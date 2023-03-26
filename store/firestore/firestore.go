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

type firestoreClient struct {
	fl firestore.Client
}

func NewClient() (*firestore.Client, error) {
	ctx := context.Background()
	err := godotenv.Load(fmt.Sprintf("../env/%s.env", os.Getenv("GO_ENV")))

	opt := option.WithCredentialsFile("../credential/busdes-firestore.json")
	client, err := firestore.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_CODE"), opt)
	if err != nil {
		return client, fmt.Errorf("error get data: %v", err)
	}

	return client, nil
}

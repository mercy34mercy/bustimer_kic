package firestore

import (
	"context"
	"fmt"
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
	Find(echo.Context, string,string)(domain.Busstop,error)
	Delete(echo.Context, string,string)(error)
	Fetch(echo.Context, string,domain.Busstop)error
}

type firestoreClient struct {
	fl firestore.Client
}

func NewClient()(firestore.Client,error){
	ctx := context.Background()
	err := godotenv.Load(fmt.Sprintf("../env/%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        // .env読めなかった場合の処理
    }
	// opt := option.WithCredentialsFile("key.json")
	client,err := firestore.NewClient(ctx,"")
	if err != nil {
		fmt.Printf("error get data: %v", err)
	}
}
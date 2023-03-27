package bootstrap

import (
	cloudFirestore "cloud.google.com/go/firestore"
	"context"
	"github.com/mercy34mercy/bustimer_kic/store/firestore"
	"log"
)

func NewFirestoreDatabase(env *Env) *cloudFirestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, env.FirebaseProjectID, env.Option)
	if err != nil {
		log.Fatal("can not get firestore client", err)
	}
	return client
}

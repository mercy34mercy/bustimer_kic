package bootstrap

import (
	"cloud.google.com/go/firestore"
	"context"

	"log"
)

func NewFirestoreDatabase(env *Env) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, env.FirebaseProjectID, env.Option)
	if err != nil {
		log.Fatal("can not get firestore client", err)
	}
	return client
}

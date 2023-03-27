package boostrap

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

func NewFIrestoreDatabase(env *Env) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, env.FirebaseProjectID, env.Option)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

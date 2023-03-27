package firestore_test

import (
	"context"
	"github.com/mercy34mercy/bustimer_kic/store/bootstrap"
	"github.com/mercy34mercy/bustimer_kic/store/firestore"
	"testing"
)

func TestNewClient(t *testing.T) {
	app := bootstrap.App()
	env := app.Env
	ctx := context.Background()

	_, err := firestore.NewClient(ctx, env.FirebaseProjectID, env.Option)
	if err != nil {
		t.Errorf("can not get client : %v", err)
	}

}

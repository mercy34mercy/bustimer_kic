package bootstrap

import "cloud.google.com/go/firestore"

type Application struct {
	Env       *Env
	Firestore *firestore.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Firestore = NewFirestoreDatabase(app.Env)
	return *app
}

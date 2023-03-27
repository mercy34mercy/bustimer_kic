package boostrap

import "cloud.google.com/go/firestore"

type Application struct {
	Env       *Env
	firestore *firestore.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.firestore = NewFIrestoreDatabase(app.Env)
	return *app
}

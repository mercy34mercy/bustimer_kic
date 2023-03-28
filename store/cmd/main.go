package main

import (
	"github.com/labstack/echo"
	"github.com/mercy34mercy/bustimer_kic/store/bootstrap"
	"os"

	"github.com/mercy34mercy/bustimer_kic/store/api/route"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	firestore := app.Firestore

	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := route.Setup(e, env, firestore)
	router.Debug = true
	router.Logger.Fatal(router.Start(":" + port))
}

package main

import (
	"os"

	"github.com/mercy34mercy/bustimer_kic/store/api/route"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := route.Setup()
	router.Debug = true
	router.Logger.Fatal(router.Start(":" + port))

}
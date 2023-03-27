package boostrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
)

type Env struct {
	FirebaseProjectID string
	Option            option.ClientOption
}

func NewEnv() *Env {
	err := godotenv.Load(fmt.Sprintf("../env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Environment can't be loaded", err)
	}
	env := &Env{
		FirebaseProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
		Option:            option.WithCredentialsFile("../credential/busdes-firestore.json"),
	}
	return env
}

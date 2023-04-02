package booststrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	RedisDatabaseURL string
)

func NewEnv() {
	err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Environment can't be loaded", err)
	}
	RedisDatabaseURL = os.Getenv("REDIS_DATABASE_URL")
}

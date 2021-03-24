package Config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load(env string) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

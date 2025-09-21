package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	path := ".env"
	counter := 0

	err := godotenv.Load(path)

	for err != nil {
		counter++

		path = "../" + path

		err = godotenv.Load(path)

		if counter == 10 {
			break
		}
	}

	if err == nil {
		log.Println(".env loaded")
	} else {
		log.Println(".env is not loaded")
	}
}

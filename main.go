package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/one-d-plate/one-svc.git/cmd"
	"github.com/one-d-plate/one-svc.git/src/pkg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pkg.SignalInit()

	cmd.Execute()
}

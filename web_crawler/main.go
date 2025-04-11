package main

import (
	"log"

	"web_crawler/stages"

	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Need to do it sequentially, because subsequent stages will likely depend on the previous ones
	// stages.RunStage0()
	// stages.RunStage1()
	stages.RunStage2()

}

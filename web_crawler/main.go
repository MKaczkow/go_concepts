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

	// Comment / note for future self regarding the stages
	//
	// stages are basicall 'scripts' executing different steps of the pipeline,
	// so no real way to test them without mocking the webpage, whichi is a lot of boring work
	//
	// STAGE 0 -> recon, get 'sitemap', i.e. list of places to scrape
	// STAGE 1 -> scrape the pages, get the data and save them to JSON files
	// STAGE 2 -> scrape the pages, get the data and save them to database (MongoDB)
	//
	// so typical workflow would be like: stage 0 -> stage 2 or stage 0 -> stage 1
	// (just uncomment the lines below)

	// Need to do it sequentially, because subsequent stages will likely depend on the previous ones
	// stages.RunStage0()
	// stages.RunStage1()
	// stages.RunStage2()
	stages.RunStage3()
}

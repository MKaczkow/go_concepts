package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

func temp_main() {
	// Create a file to save the scraped data
	f, err := os.Create("places.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer f.Close()

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a CSV writer to write data to the file
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Write the CSV header
	err = writer.Write([]string{"place_url", "found_on_page"})
	if err != nil {
		log.Fatal("Cannot write header to file", err)
	}

	// Read allowed domain from environment variable
	allowedDomain := os.Getenv("ALLOWED_DOMAIN")
	if allowedDomain == "" {
		log.Fatal("ALLOWED_DOMAIN environment variable is required")
	}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomain),
	)

	// Begin scrapping
	fmt.Println("Starting scraping stage 0 - collecting place URLs")

	// Regular expression to extract page number from URL
	pageNumberRegex := regexp.MustCompile(`page=(\d+)`)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href^='/pl/place/']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absoluteURL := e.Request.AbsoluteURL(link)

		// Extract the page number from the URL
		pageNumberMatch := pageNumberRegex.FindStringSubmatch(e.Request.URL.String())
		var pageNumber int
		if len(pageNumberMatch) > 1 {
			pageNumber, err = strconv.Atoi(pageNumberMatch[1])
			if err != nil {
				log.Printf("Error converting page number: %s", err)
				return
			}
		}

		// Write the data to the CSV file
		err := writer.Write([]string{absoluteURL, strconv.Itoa(pageNumber)})
		if err != nil {
			log.Println("Cannot write to file", err)
		}
		fmt.Printf("Place URL found on page %d: %s\n", pageNumber, absoluteURL)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Define the range of pages to visit
	// TODO: This is hardcoded now, but should be dynamically set,
	// based on the number of pages available on the website
	// and 'last visited' page number stored in a database
	for i := 1; i <= 198; i++ {
		url := fmt.Sprintf("https://%s/pl/places?page=%d", allowedDomain, i)
		err := c.Visit(url)
		if err != nil {
			log.Printf("Error visiting page %d: %s", i, err)
		}
	}

	// Wait for all requests to finish
	c.Wait()

	// Close the CSV file
	err = f.Close()
	if err != nil {
		log.Fatal("Cannot close file", err)
	}
	fmt.Println("Scraping stage 0 finished. Data saved to places.csv")

	// Visit random page from places.csv and scrape details
	fmt.Println("Starting scraping stage 1 - picking a random place to scrape details")

	// Open the CSV file
	csvFile, err := os.Open("places.csv")
	if err != nil {
		log.Fatal("Cannot open CSV file", err)
	}
	defer csvFile.Close()

	// Read the CSV file
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Cannot read CSV file", err)
	}

	// Skip header row
	if len(records) <= 1 {
		log.Fatal("CSV file has no data rows")
	}
	records = records[1:]

	// Get a random record
	if len(records) == 0 {
		log.Fatal("No places to scrape")
	}

	// Create directory for scraped data
	// Create 'runs' directory if it doesn't exist
	runsDir := "runs"
	if err := os.MkdirAll(runsDir, 0755); err != nil {
		log.Fatalf("Failed to create runs directory: %v", err)
	}

	// Create a subdirectory with human-readable timestamp
	runTime := time.Now().Format("2006-01-02_15-04-05")
	runDir := fmt.Sprintf("%s/%s", runsDir, runTime)
	if err := os.MkdirAll(runDir, 0755); err != nil {
		log.Fatalf("Failed to create run directory: %v", err)
	}

	fmt.Printf("Created run directory: %s\n", runDir)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(records))
	randomPlace := records[randomIndex]

	fmt.Printf("Selected random place: %s (found on page %s)\n", randomPlace[0], randomPlace[1])

	// Create a new collector for scraping place details

	fmt.Println("Scraping stage 1 finished. Data saved to places.csv")

}

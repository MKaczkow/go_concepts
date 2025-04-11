package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"web_crawler/models"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	// Initialize a PlaceDetails struct
	// Split URL by '/' and use the last part as the name
	parts := strings.Split(randomPlace[0], "/")
	placeName := ""
	if len(parts) > 0 {
		lastPart := parts[len(parts)-1]
		if lastPart != "" {
			placeName = lastPart
		}
	}

	// Initialize a AbandonedPlace struct, extract name and URL
	abandonedPlace := models.AbandonedPlace{
		Name: placeName,
		URL:  randomPlace[0],
	}

	fmt.Printf("Scraping details for place: %s\n", abandonedPlace.Name)
	fmt.Println("Place URL:", abandonedPlace.URL)

	// Set up callback for when a page is visited
	c.OnHTML("div[itemprop='description']", func(e *colly.HTMLElement) {

		// Extract the description
		description := e.Text
		if description != "" {
			abandonedPlace.Description = strings.TrimSpace(description)
			// fmt.Printf("Found description: %s\n", abandonedPlace.Description)
		} else {
			fmt.Println("No description found")
		}

	})

	// Extract the details
	// TBD

	// Extract comments
	// TBD

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %s", r.Request.URL, err)
	})

	// Visit the page
	err = c.Visit(abandonedPlace.URL)
	if err != nil {
		log.Printf("Failed to visit %s: %v", abandonedPlace.URL, err)
	}

	// Wait for scraping to finish
	c.Wait()

	// Save the place details to a JSON file
	detailsFilePath := filepath.Join(runDir, "place_details.json")
	detailsFile, err := os.Create(detailsFilePath)
	if err != nil {
		log.Fatalf("Failed to create details file: %v", err)
	}
	defer detailsFile.Close()

	// Marshal the data with indentation for readability
	jsonData, err := json.MarshalIndent(abandonedPlace, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal place details: %v", err)
	}

	// Write to file
	_, err = detailsFile.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write place details to file: %v", err)
	}

	fmt.Printf("Saved place details to %s\n", detailsFilePath)

	fmt.Println("Scraping stage 1 finished. Data saved to relevant run dir and place files")

}

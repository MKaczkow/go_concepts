package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

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
	abandonedPlace := AbandonedPlace{
		Name: randomPlace[0],
		URL:  randomPlace[1],
	}

	// Set up callback for when a page is visited
	c.OnHTML("body", func(e *colly.HTMLElement) {

		// Extract basic information from the page
		abandonedPlace.Name = e.FullText("h1")
		abandonedPlace.Details.Category = e.FullText("dd a[title^='Zobacz inne miejsca z tej kategorii']")
		abandonedPlace.Details.Status = e.FullText("dd span[style*='color:#ff0000']")
		abandonedPlace.Details.Location = e.FullText("dd[title^='Kołczewo']")
		abandonedPlace.Details.Coordinates = e.FullText("dd a[title='Zobacz na Google Maps']")

		// Extract meta information
		abandonedPlace.Description = e.FullText(".place-description")
		abandonedPlace.Details.AddedBy = e.FullText("dd a[title$='Profil']")
		abandonedPlace.Details.Accessibility = e.FullText(".definition-group:contains('Dostępność') dd")
		abandonedPlace.Details.Attractiveness = e.FullText(".definition-group:contains('Atrakcyjność') dd")
		abandonedPlace.Details.Views = e.FullText(".definition-group:contains('Wyświetlenia') dd")

		// Extract rating
		ratingText := e.FullText("dd[rel='tooltip'] span[itemprop='ratingValue']")
		if rating, err := strconv.ParseFloat(ratingText, 64); err == nil {
			abandonedPlace.Details.Rating = rating
		}

		// Extract image URLs
		// e.ForEach(".gallery img", func(_ int, el *colly.HTMLElement) {
		// 	abandonedPlace.Images = append(abandonedPlace.Images, el.Attr("src"))
		// })

		// Extract nearby places
		// e.ForEach("#places-nearby a", func(_ int, el *colly.HTMLElement) {
		// 	nearbyPlace := models.NearbyPlace{
		// 		Name:     el.FullText(".stripe span"),
		// 		URL:      el.Attr("href"),
		// 		Distance: el.FullText(".distance"),
		// 	}
		// 	abandonedPlace.NearbyPlaces = append(abandonedPlace.NearbyPlaces, nearbyPlace)
		// })

		// Extract comments
		e.ForEach(".fos_comment_comment_show", func(_ int, el *colly.HTMLElement) {
			comment := models.Comment{
				Author:  el.FullText(".media-heading a"),
				Date:    el.FullText(".media-date"),
				Content: el.FullText(".media-comment"),
			}
			abandonedPlace.Comments = append(abandonedPlace.Comments, comment)
		})
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %s", r.Request.URL, err)
	})

	// Visit the page
	err = c.Visit(randomPlace[1])
	if err != nil {
		log.Printf("Failed to visit %s: %v", randomPlace[1], err)
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

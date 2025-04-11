package stages

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"web_crawler/models"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RunStage0() {

	// Begin scrapping
	fmt.Println("Starting scraping stage 0 - collecting place URLs")

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

}

func RunStage1() {

	fmt.Println("Starting scraping stage 1 - collecting places details")

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
		colly.Async(true),
		colly.MaxDepth(1),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
		colly.CacheDir("./cache"),
	)

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

	if len(records) == 0 {
		log.Fatal("No places to scrape")
	}

	// Create directory for scraped data
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

	// Iterate over all places in the CSV ...
	// startIdx := 0
	// endIdx := len(records) - 1
	// ... or only specified range
	for placeIdx, randomPlace := range records {
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
			Name:       placeName,
			URL:        randomPlace[0],
			ScrapeDate: time.Now().Format("2006-01-02 15:04:05"),
		}

		fmt.Printf("Scraping details for place: %s\n", abandonedPlace.Name)
		fmt.Println("Place URL:", abandonedPlace.URL)

		// Extract the description
		c.OnHTML("div[itemprop='description'], .description, #description", func(e *colly.HTMLElement) {
			description := e.Text
			if description != "" {
				abandonedPlace.Description = strings.TrimSpace(description)
			} else {
				fmt.Println("No description found in this element")
			}
		})

		// Extract the details
		c.OnHTML("aside.col-lg-4 dl", func(e *colly.HTMLElement) {
			// Extract all definition groups
			e.ForEach("div.definition-group", func(_ int, div *colly.HTMLElement) {
				// Get the definition term and description
				dt := div.ChildText("dt")
				dd := div.ChildText("dd")

				switch dt {
				case "Dodano":
					abandonedPlace.Details.Added = strings.TrimSpace(dd)
				case "Dodał":
					abandonedPlace.Details.AddedBy = strings.TrimSpace(dd)
				case "Dostępność":
					abandonedPlace.Details.Accessibility = strings.TrimSpace(dd)
				case "Atrakcyjność":
					abandonedPlace.Details.Attractiveness = strings.TrimSpace(dd)
				case "Kategoria":
					abandonedPlace.Details.Category = strings.TrimSpace(dd)
				case "Koordynaty":
					// Coordinates are expected to be in the format "X, Y", e.g. "50.0437530, 18.2580460"
					coords := strings.Split(dd, ",")
					if len(coords) == 2 {
						latitude, err := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
						if err != nil {
							log.Printf("Failed to parse latitude: %v", err)
						} else {
							abandonedPlace.Details.Coordinates.Latitude = latitude
						}
						longitude, err := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
						if err != nil {
							log.Printf("Failed to parse longitude: %v", err)
						} else {
							abandonedPlace.Details.Coordinates.Longitude = longitude
						}
					} else {
						log.Printf("Unexpected coordinates format: %s", dd)
					}
				case "Lokalizacja":
					abandonedPlace.Details.Location = strings.TrimSpace(dd)
				case "Ocena":
					// Rating is expected to be in the format "X/Y", in 'school grading' system, i.e. 5 is max, 1 is min.
					// Rating is then parsed to float and stored in the struct
					ratingString := strings.TrimSpace(dd)
					ratingParts := strings.Split(ratingString, "/")
					if len(ratingParts) == 2 {
						rating, err := strconv.ParseFloat(strings.TrimSpace(ratingParts[0]), 64)
						if err != nil {
							log.Printf("Failed to parse rating: %v", err)
						} else {
							abandonedPlace.Details.Rating = rating
						}
					} else if len(ratingParts) == 1 {
						// Assuming this is the case with "Brak głosów", so just insert 0 nicely
						abandonedPlace.Details.Rating = 0
					} else {
						log.Printf("Unexpected rating format: %s", ratingString)
					}
				case "Status":
					abandonedPlace.Details.Status = strings.TrimSpace(dd)
				}
			})
			fmt.Println("Extracted details for place:", abandonedPlace.Name)
		})

		// Extract comments
		c.OnHTML("div#comments ul.media-list .fos_comment_comment_show", func(e *colly.HTMLElement) {
			comment := models.Comment{}
			comment.User = e.ChildText("h4.media-heading a")
			comment.Timestamp = e.ChildText("time.media-date")
			comment.CommentText = e.ChildText("p.media-comment")

			// Only add non-empty comments
			if comment.CommentText != "" {
				abandonedPlace.Comments = append(abandonedPlace.Comments, comment)
			}
		})

		// Extract hazards
		c.OnHTML("div.row.vcenter.margin-bottom-sm-15", func(e *colly.HTMLElement) {
			hazard := models.Hazard{}

			// Extract the description
			hazard.Description = e.ChildText("div.col-lg-10 > span:first-child")

			// Extract added by information
			addedByText := e.ChildText("div.col-lg-10 span.text-muted.small")
			if addedByText != "" {
				parts := strings.Split(addedByText, "Dodał")
				if len(parts) > 1 {
					userTimePart := strings.TrimSpace(parts[1])

					// Extract user - it's between the opening and closing tags
					userStart := strings.Index(userTimePart, ">")
					userEnd := strings.Index(userTimePart, "</a>")
					if userStart != -1 && userEnd != -1 && userStart < userEnd {
						hazard.AddedBy = strings.TrimSpace(userTimePart[userStart+1 : userEnd])
					}

					// Extract timestamp - it should be the last part
					timeStart := strings.LastIndex(userTimePart, ">")
					if timeStart != -1 && timeStart+1 < len(userTimePart) {
						hazard.Added = strings.TrimSpace(userTimePart[timeStart+1:])
					}
				}
			}

			// Only add if we have at least a description
			if hazard.Description != "" {
				abandonedPlace.Hazards = append(abandonedPlace.Hazards, hazard)
			}
		})

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
		detailsFilePath := filepath.Join(runDir, fmt.Sprintf("place_details_%05d.json", placeIdx))
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

	}

	fmt.Println("Scraping stage 1 finished. Data saved to relevant run dir and place files")

}

func RunStage2() {

	fmt.Println("Starting scraping stage 2 - collecting places details and putting to db")

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
		colly.Async(true),
		colly.MaxDepth(1),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
		colly.CacheDir("./cache"),
	)

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

	if len(records) == 0 {
		log.Fatal("No places to scrape")
	}

	// Initialize MongoDB connection
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017" // Default local MongoDB URI
	}

	// Connect to MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Iterate over all places in the CSV ...
	// startIdx := 0
	// endIdx := len(records) - 1
	// ... or only specified range
	for _, randomPlace := range records {
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
			Name:       placeName,
			URL:        randomPlace[0],
			ScrapeDate: time.Now().Format("2006-01-02 15:04:05"),
		}

		fmt.Printf("Scraping details for place: %s\n", abandonedPlace.Name)
		fmt.Println("Place URL:", abandonedPlace.URL)

		// Extract the description
		c.OnHTML("div[itemprop='description'], .description, #description", func(e *colly.HTMLElement) {
			description := e.Text
			if description != "" {
				abandonedPlace.Description = strings.TrimSpace(description)
			} else {
				fmt.Println("No description found in this element")
			}
		})

		// Extract the details
		c.OnHTML("aside.col-lg-4 dl", func(e *colly.HTMLElement) {
			// Extract all definition groups
			e.ForEach("div.definition-group", func(_ int, div *colly.HTMLElement) {
				// Get the definition term and description
				dt := div.ChildText("dt")
				dd := div.ChildText("dd")

				switch dt {
				case "Dodano":
					abandonedPlace.Details.Added = strings.TrimSpace(dd)
				case "Dodał":
					abandonedPlace.Details.AddedBy = strings.TrimSpace(dd)
				case "Dostępność":
					abandonedPlace.Details.Accessibility = strings.TrimSpace(dd)
				case "Atrakcyjność":
					abandonedPlace.Details.Attractiveness = strings.TrimSpace(dd)
				case "Kategoria":
					abandonedPlace.Details.Category = strings.TrimSpace(dd)
				case "Koordynaty":
					// Coordinates are expected to be in the format "X, Y", e.g. "50.0437530, 18.2580460"
					coords := strings.Split(dd, ",")
					if len(coords) == 2 {
						latitude, err := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
						if err != nil {
							log.Printf("Failed to parse latitude: %v", err)
						} else {
							abandonedPlace.Details.Coordinates.Latitude = latitude
						}
						longitude, err := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
						if err != nil {
							log.Printf("Failed to parse longitude: %v", err)
						} else {
							abandonedPlace.Details.Coordinates.Longitude = longitude
						}
					} else {
						log.Printf("Unexpected coordinates format: %s", dd)
					}
				case "Lokalizacja":
					abandonedPlace.Details.Location = strings.TrimSpace(dd)
				case "Ocena":
					// Rating is expected to be in the format "X/Y", in 'school grading' system, i.e. 5 is max, 1 is min.
					// Rating is then parsed to float and stored in the struct
					ratingString := strings.TrimSpace(dd)
					ratingParts := strings.Split(ratingString, "/")
					if len(ratingParts) == 2 {
						rating, err := strconv.ParseFloat(strings.TrimSpace(ratingParts[0]), 64)
						if err != nil {
							log.Printf("Failed to parse rating: %v", err)
						} else {
							abandonedPlace.Details.Rating = rating
						}
					} else if len(ratingParts) == 1 && ratingParts[0] == "Brak głosów" {
						// Assuming this is the case with "Brak głosów", so just insert 0 nicely
						abandonedPlace.Details.Rating = 0
					} else {
						log.Printf("Unexpected rating format: %s", ratingString)
					}
				case "Status":
					abandonedPlace.Details.Status = strings.TrimSpace(dd)
				}
			})
			// TODO: 'silent mode'
			// fmt.Println("Extracted details for place:", abandonedPlace.Name)
		})

		// Extract comments
		c.OnHTML("div#comments ul.media-list .fos_comment_comment_show", func(e *colly.HTMLElement) {
			comment := models.Comment{}
			comment.User = e.ChildText("h4.media-heading a")
			comment.Timestamp = e.ChildText("time.media-date")
			comment.CommentText = e.ChildText("p.media-comment")

			// Only add non-empty comments
			if comment.CommentText != "" {
				abandonedPlace.Comments = append(abandonedPlace.Comments, comment)
			}
		})

		// Extract hazards
		c.OnHTML("div.row.vcenter.margin-bottom-sm-15", func(e *colly.HTMLElement) {
			hazard := models.Hazard{}

			// Extract the description
			hazard.Description = e.ChildText("div.col-lg-10 > span:first-child")

			// Extract added by information
			addedByText := e.ChildText("div.col-lg-10 span.text-muted.small")
			if addedByText != "" {
				parts := strings.Split(addedByText, "Dodał")
				if len(parts) > 1 {
					userTimePart := strings.TrimSpace(parts[1])

					// Extract user - it's between the opening and closing tags
					userStart := strings.Index(userTimePart, ">")
					userEnd := strings.Index(userTimePart, "</a>")
					if userStart != -1 && userEnd != -1 && userStart < userEnd {
						hazard.AddedBy = strings.TrimSpace(userTimePart[userStart+1 : userEnd])
					}

					// Extract timestamp - it should be the last part
					timeStart := strings.LastIndex(userTimePart, ">")
					if timeStart != -1 && timeStart+1 < len(userTimePart) {
						hazard.Added = strings.TrimSpace(userTimePart[timeStart+1:])
					}
				}
			}

			// Only add if we have at least a description
			if hazard.Description != "" {
				abandonedPlace.Hazards = append(abandonedPlace.Hazards, hazard)
			}
		})

		// Set up error handling
		c.OnError(func(r *colly.Response, err error) {
			log.Printf("Error scraping %s: %s", r.Request.URL, err)
		})

		// Visit the page
		err = c.Visit(abandonedPlace.URL)
		if err != nil {
			log.Printf("Failed to visit %s: %v", abandonedPlace.URL, err)
		}

		// Save places details to MongoDB
		collection := client.Database("abandoned_places").Collection("places")

		// Check if this place already exists in the database
		cursor, err := collection.Find(
			context.TODO(),
			map[string]string{"url": abandonedPlace.URL},
		)
		if err != nil {
			log.Printf("Error checking for existing document: %v", err)
		}

		// If place exists, update it; otherwise insert new document
		count := 0
		if cursor.Next(context.TODO()) {
			count++
		}
		cursor.Close(context.TODO())

		if count > 0 {
			_, err = collection.ReplaceOne(
				context.TODO(),
				map[string]string{"url": abandonedPlace.URL},
				abandonedPlace,
			)
			if err != nil {
				log.Printf("Failed to update place in MongoDB: %v", err)
			} else {
				fmt.Printf("Updated place in MongoDB: %s\n", abandonedPlace.Name)
			}
		} else {
			_, err = collection.InsertOne(context.TODO(), abandonedPlace)
			if err != nil {
				log.Printf("Failed to insert place into MongoDB: %v", err)
			} else {
				fmt.Printf("Inserted place into MongoDB: %s\n", abandonedPlace.Name)
			}
		}

		// Wait for scraping to finish
		c.Wait()

	}

	fmt.Println("Scraping stage 2 finished. Data saved to database")
}

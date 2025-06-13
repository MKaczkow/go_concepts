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
	"web_crawler/utils"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RunStage0() {

	// Begin scrapping
	fmt.Println("Starting scraping stage 0 - collecting place URLs")

	// If places.csv file exists rename it for future reference
	if _, err := os.Stat("places.csv"); err == nil {
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		newPath := fmt.Sprintf("places_%s.csv", timestamp)
		if err := os.Rename("places.csv", newPath); err != nil {
			log.Printf("Warning: Could not rename places.csv: %v", err)
		}
	}

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
	// TODO: This is hardcoded in env now, but should be dynamically set,
	// based on the number of pages available on the website
	// and 'last visited' page number stored in a database

	// Read allowed domain from environment variable and cast to integerr
	maxPageNumber := os.Getenv("MAX_PAGE_NUMBER")
	if maxPageNumber == "" {
		log.Fatal("MAX_PAGE_NUMBER environment variable is required")
	}
	maxPageNumberInt, err := strconv.Atoi(maxPageNumber)
	if err != nil {
		log.Fatalf("Error converting MAX_PAGE_NUMBER to integer: %s", err)
	}
	if maxPageNumberInt <= 0 {
		log.Fatal("MAX_PAGE_NUMBER must be a positive integer")
	}

	for i := 1; i <= maxPageNumberInt; i++ {
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
		colly.UserAgent(utils.GetRandomUserAgent()),
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

			// Extract the hazard name and ID from the img title attribute
			imgTitle := e.ChildAttr("div.col-lg-2 img", "title")
			if h, ok := models.HazardTypes[imgTitle]; ok {
				hazard.ID = h.ID
				hazard.Name = h.Name
			} else {
				// Handle cases where the title might not match a known hazard type
				// For example, log a warning or set a default ID/Name
				hazard.Name = "Nieznane"
				hazard.ID = 7
				log.Printf("Unknown hazard type: %s, using default ID 7 and 'Nieznane' category", imgTitle)
			}

			// Extract the description
			hazard.Description = e.ChildText("div.col-lg-10 > span:first-child")

			// Extract added by information and date
			addedByText := e.ChildText("div.col-lg-10 span.text-muted.small")
			if addedByText != "" {
				// Extract user by selecting the text content of the <a> tag
				hazard.AddedBy = e.ChildText("div.col-lg-10 span.text-muted.small a.text-muted")

				// Extract timestamp
				// The timestamp is located directly after the </a> tag within the 'small' span.
				// We can get the full text and then remove the part containing "Dodał <username>"
				// A more robust way is to get the outerHTML of the 'small' span and parse it.
				smallSpanHTML, err := e.DOM.Find("div.col-lg-10 span.text-muted.small").Html()
				if err == nil {
					// Find the index of the closing </a> tag
					anchorEndIndex := strings.Index(smallSpanHTML, "</a>")
					if anchorEndIndex != -1 {
						// The timestamp starts right after "</a>"
						// We need to trim spaces and newlines
						hazard.Added = strings.TrimSpace(smallSpanHTML[anchorEndIndex+len("</a>"):])
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
		colly.UserAgent(utils.GetRandomUserAgent()),
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

func RunStage3() {

	fmt.Println("Starting scraping stage 3 - collecting users details")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read allowed domain from environment variable
	allowedDomain := os.Getenv("ALLOWED_DOMAIN")
	if allowedDomain == "" {
		log.Fatal("ALLOWED_DOMAIN environment variable is required")
	}

	// Initialize MongoDB connection
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
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

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomain),
		colly.Async(true),
		colly.MaxDepth(1),
		colly.UserAgent(utils.GetRandomUserAgent()),
		colly.CacheDir("./cache"),
	)

	// Extract user details
	c.OnHTML("div.tab-pane.profile", func(e *colly.HTMLElement) {
		user := models.User{
			ScrapeDate: time.Now().Format("2006-01-02 15:04:05"),
		}
		// Get username from the username dd element
		user.Username = strings.TrimSpace(e.ChildText("dl dd"))

		// Get join date and last login
		e.ForEach("dl dd", func(i int, dd *colly.HTMLElement) {
			switch i {
			case 1:
				user.JoinedDate = strings.TrimSpace(dd.Attr("title"))
			case 2:
				user.LastLogin = strings.TrimSpace(dd.Attr("title"))
			}
		})

		// Get stats
		// user.VisitedPlaces = len(e.DOM.ParentsUntil("~").Find("#visited-places .category-container").Nodes)
		// user.AddedPlaces = len(e.DOM.ParentsUntil("~").Find("#added-places .category-container").Nodes)
		// user.AddedPictures = len(e.DOM.ParentsUntil("~").Find("#added-pictures .category-container").Nodes)
		user.CommentsCount = len(e.DOM.ParentsUntil("~").Find("#added-comments li").Nodes)
		user.ChangesCount = len(e.DOM.ParentsUntil("~").Find("#added-changes li").Nodes)

		// Save user to MongoDB
		collection := client.Database("abandoned_places").Collection("users")

		// Check if user exists
		cursor, err := collection.Find(context.TODO(), map[string]string{"username": user.Username})
		if err != nil {
			log.Printf("Error checking for existing user: %v", err)
			return
		}

		count := 0
		if cursor.Next(context.TODO()) {
			count++
		}
		cursor.Close(context.TODO())

		if count > 0 {
			_, err = collection.ReplaceOne(
				context.TODO(),
				map[string]string{"username": user.Username},
				user,
			)
			if err != nil {
				log.Printf("Failed to update user in MongoDB: %v", err)
			} else {
				fmt.Printf("Updated user in MongoDB: %s\n", user.Username)
			}
		} else {
			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				log.Printf("Failed to insert user into MongoDB: %v", err)
			} else {
				fmt.Printf("Inserted user into MongoDB: %s\n", user.Username)
			}
		}
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %s", r.Request.URL, err)
	})

	// Query MongoDB for the highest user ID seen so far
	collection := client.Database("abandoned_places").Collection("users")
	lastUserDoc := bson.D{}
	opts := options.FindOne().SetSort(bson.D{{Key: "userid", Value: -1}})
	err = collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastUserDoc)

	// Assuming lastUserDoc is from MongoDB
	startID := 1
	if err == nil && len(lastUserDoc) > 0 {
		for _, elem := range lastUserDoc {
			if elem.Key == "userid" {
				if userID, ok := elem.Value.(int64); ok {
					startID = int(userID) + 1
					break
				}
			}
		}
	}

	// Read MAX_USER_PAGE_NUMBER from environment variable
	maxUserPageNumber := os.Getenv("MAX_USER_PAGE_NUMBER")
	if maxUserPageNumber == "" {
		log.Fatal("MAX_USER_PAGE_NUMBER environment variable is required")
	}
	maxUserPageNumberInt, err := strconv.Atoi(maxUserPageNumber)
	if err != nil {
		log.Fatalf("Error converting MAX_USER_PAGE_NUMBER to integer: %s", err)
	}
	if maxUserPageNumberInt <= 0 {
		log.Fatal("MAX_USER_PAGE_NUMBER must be a positive integer")
	}

	// Visit user profiles sequentially up to MAX_USER_PAGE_NUMBER
	maxFailures := 10
	consecutiveFailures := 0

	for i := startID; i <= maxUserPageNumberInt && consecutiveFailures < maxFailures; i++ {
		userURL := fmt.Sprintf("https://%s/pl/profile/show/%d", allowedDomain, i)
		err = c.Visit(userURL)
		if err != nil {
			consecutiveFailures++
			log.Printf("Failed to visit user %d: %v", i, err)
		} else {
			consecutiveFailures = 0
		}
	}

	// Wait for scraping to finish
	c.Wait()

	// Close MongoDB connection
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scraping stage 3 finished. Data saved to database")

}

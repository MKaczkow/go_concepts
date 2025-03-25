package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Create a file to save the scraped data
	f, err := os.Create("places.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer f.Close()

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
}

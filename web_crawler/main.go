package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {

	// Create a file to save the scraped data
	f, err := os.Create("data.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer f.Close()

	// Create a CSV writer to write data to the file
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		linkText := e.Text
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Write the data to the CSV file
		writer.Write([]string{linkText, link})
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}

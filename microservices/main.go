package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var omdbAPIURL string

func init() {
	apiKey, err := readAPIKeyFromFile("api_key")
	if err != nil {
		log.Fatal("Failed to read API key:", err)
	}

	omdbAPIURL = fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&", apiKey)
}

func readAPIKeyFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

// Rest of the code...

func main() {
	router := gin.Default()

	router.GET("/movies", getMovie)

	router.Run(":8080")
}

func getMovie(c *gin.Context) {
	title := c.Query("t")
	year := c.Query("y")
	id := c.Query("i")
	movieType := c.Query("type")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	// Make a request to OMDb API
	apiURL := fmt.Sprintf("%st=%s&i=%s&y=%s&type=%s&r=json", omdbAPIURL, title, id, year, movieType)
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie data"})
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": string(body)})
}

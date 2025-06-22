package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Serve the frontend
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Weather API endpoint
	router.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		lat := c.Query("lat")
		lng := c.Query("lng")
		apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")

		if apiKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "API key is not set"})
			return
		}

		var weatherURL string
		if city != "" {
			// City-based query
			weatherURL = "http://api.openweathermap.org/data/2.5/weather?q=" + url.QueryEscape(city) + "&appid=" + apiKey + "&units=metric"
		} else if lat != "" && lng != "" {
			// Coordinate-based query
			weatherURL = "http://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lng + "&appid=" + apiKey + "&units=metric"
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Either city or lat/lng coordinates are required"})
			return
		}

		fmt.Println(weatherURL)
		resp, err := http.Get(weatherURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
			return
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		log.Printf("Response Body: %s", string(bodyBytes))

		var weatherData interface{}
		if err := json.Unmarshal(bodyBytes, &weatherData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
			return
		}

		c.JSON(http.StatusOK, weatherData)
	})

	log.Println("Starting server on :58080")
	if err := router.Run(":58080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

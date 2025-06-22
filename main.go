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

func fetchWeather(city, lat, lng, apiKey string) (interface{}, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is not set")
	}

	var weatherURL string
	if city != "" {
		weatherURL = "http://api.openweathermap.org/data/2.5/weather?q=" + url.QueryEscape(city) + "&appid=" + apiKey + "&units=metric"
	} else if lat != "" && lng != "" {
		weatherURL = "http://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lng + "&appid=" + apiKey + "&units=metric"
	} else {
		return nil, fmt.Errorf("Either city or lat/lng coordinates are required")
	}

	fmt.Println(weatherURL)
	resp, err := http.Get(weatherURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch weather data")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body")
	}

	log.Printf("Response Body: %s", string(bodyBytes))

	var weatherData interface{}
	if err := json.Unmarshal(bodyBytes, &weatherData); err != nil {
		return nil, fmt.Errorf("Failed to parse weather data")
	}

	return weatherData, nil
}

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

		weatherData, err := fetchWeather(city, lat, lng, apiKey)
		if err != nil {
			if err.Error() == "API key is not set" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else if err.Error() == "Either city or lat/lng coordinates are required" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, weatherData)
	})

	log.Println("Starting server on :58080")
	if err := router.Run(":58080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

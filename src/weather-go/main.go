package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

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

func fetchTimezone(lat, lng, apiKey string) (interface{}, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is not set")
	}

	// Use OpenWeatherMap's timezone API
	timezoneURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lng, apiKey)

	resp, err := http.Get(timezoneURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch timezone data")
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read timezone response body")
	}

	var timezoneData interface{}
	if err := json.Unmarshal(bodyBytes, &timezoneData); err != nil {
		return nil, fmt.Errorf("Failed to parse timezone data")
	}

	return timezoneData, nil
}

func getLocalTime(lat, lng float64) (time.Time, error) {
	// Use a simple timezone calculation based on longitude
	// This is a basic approximation - for more accuracy, we'd need a proper timezone database
	hoursOffset := int(lng / 15.0) // Rough timezone calculation

	utc := time.Now().UTC()
	localTime := utc.Add(time.Duration(hoursOffset) * time.Hour)

	return localTime, nil
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

		// Add local time information using timezone from API response
		if weatherMap, ok := weatherData.(map[string]interface{}); ok {
			// Check if timezone information is available in the response
			if timezoneOffset, exists := weatherMap["timezone"]; exists {
				if offsetSeconds, ok := timezoneOffset.(float64); ok {
					// Calculate local time using the timezone offset
					utc := time.Now().UTC()
					localTime := utc.Add(time.Duration(offsetSeconds) * time.Second)

					weatherMap["local_time"] = localTime.Format("2006-01-02 15:04:05")
					weatherMap["timezone_offset_hours"] = int(offsetSeconds / 3600)
					weatherData = weatherMap
				}
			} else if lat != "" && lng != "" {
				// Fallback to longitude-based calculation if no timezone data
				if latFloat, err := strconv.ParseFloat(lat, 64); err == nil {
					if lngFloat, err := strconv.ParseFloat(lng, 64); err == nil {
						if localTime, err := getLocalTime(latFloat, lngFloat); err == nil {
							weatherMap["local_time"] = localTime.Format("2006-01-02 15:04:05")
							weatherMap["timezone_offset"] = int(lngFloat / 15.0)
							weatherData = weatherMap
						}
					}
				}
			}
		}

		c.JSON(http.StatusOK, weatherData)
	})

	// Timezone API endpoint
	router.GET("/timezone", func(c *gin.Context) {
		lat := c.Query("lat")
		lng := c.Query("lng")
		apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")

		if lat == "" || lng == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
			return
		}

		timezoneData, err := fetchTimezone(lat, lng, apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, timezoneData)
	})

	log.Println("Starting server on :58080")
	if err := router.Run(":58080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

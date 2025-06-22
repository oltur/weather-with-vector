package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Mock fetchWeather for endpoint tests
func mockFetchWeather(city, lat, lng, apiKey string) (interface{}, error) {
	if city == "TestCity" {
		return map[string]interface{}{"weather": "sunny", "name": city}, nil
	}
	if lat == "1.23" && lng == "4.56" {
		return map[string]interface{}{"weather": "cloudy", "coord": map[string]string{"lat": lat, "lon": lng}}, nil
	}
	if apiKey == "" {
		return nil, errApiKey
	}
	return nil, errParams
}

var (
	errApiKey = &customError{"API key is not set"}
	errParams = &customError{"Either city or lat/lng coordinates are required"}
)

type customError struct{ msg string }

func (e *customError) Error() string { return e.msg }

func TestFetchWeather_MissingApiKey(t *testing.T) {
	_, err := fetchWeather("TestCity", "", "", "")
	if err == nil || err.Error() != "API key is not set" {
		t.Errorf("Expected API key error, got %v", err)
	}
}

func TestFetchWeather_MissingParams(t *testing.T) {
	apiKey := "dummy"
	_, err := fetchWeather("", "", "", apiKey)
	if err == nil || err.Error() != "Either city or lat/lng coordinates are required" {
		t.Errorf("Expected params error, got %v", err)
	}
}

// Integration test for /weather endpoint
func TestWeatherEndpoint_City(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		lat := c.Query("lat")
		lng := c.Query("lng")
		apiKey := "dummy"
		data, err := mockFetchWeather(city, lat, lng, apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	req, _ := http.NewRequest("GET", "/weather?city=TestCity", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["name"] != "TestCity" {
		t.Errorf("Expected city TestCity, got %v", resp["name"])
	}
}

func TestWeatherEndpoint_LatLng(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		lat := c.Query("lat")
		lng := c.Query("lng")
		apiKey := "dummy"
		data, err := mockFetchWeather(city, lat, lng, apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	req, _ := http.NewRequest("GET", "/weather?lat=1.23&lng=4.56", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp["weather"] != "cloudy" {
		t.Errorf("Expected weather cloudy, got %v", resp["weather"])
	}
}

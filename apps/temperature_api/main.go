package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type TemperatureResponse struct {
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
	Location  string    `json:"location"`
	Status    string    `json:"status"`
	SensorID  string    `json:"sensor_id"`
}

func main() {
	router := gin.Default()

	getTemperature := func(c *gin.Context) {
		// Location is part of the resource URI so callers can request, for
		// example, /temperature/Living%20Room.
		location := c.Param("location")
		sensorID := c.Query("sensor_id")

		// Retain the query parameter as a backwards-compatible fallback for
		// sensor-ID lookups and existing clients.
		if location == "" {
			location = c.Query("location")
		}

		if location == "" {
			switch sensorID {
			case "1":
				location = "Living Room"
			case "2":
				location = "Bedroom"
			case "3":
				location = "Kitchen"
			default:
				location = "Unknown"
			}
		}

		if sensorID == "" {
			switch location {
			case "Living Room":
				sensorID = "1"
			case "Bedroom":
				sensorID = "2"
			case "Kitchen":
				sensorID = "3"
			default:
				sensorID = "0"
			}
		}

		randomTemp := 15.0 + rand.Float64()*(30.0-15.0)

		resp := TemperatureResponse{
			Value:     randomTemp,
			Unit:      "Celsius",
			Timestamp: time.Now(),
			Location:  location,
			Status:    "active",
			SensorID:  sensorID,
		}

		c.JSON(http.StatusOK, resp)
	}

	// Versioned public endpoint.
	router.GET("/api/v1/temperature/:location", getTemperature)
	// Kept for existing smart-home requests and backwards compatibility.
	router.GET("/temperature/:location", getTemperature)
	router.GET("/temperature", getTemperature)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router.Run(":" + port)
}

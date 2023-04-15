package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/sensor/pkg/domain"
	sHttp "github.com/pklimuk-eng-thesis/sensor/pkg/http"
	sService "github.com/pklimuk-eng-thesis/sensor/pkg/service"
)

func main() {
	// Initialization of the sensor
	sensor := domain.Sensor{Enabled: false, Detected: false}

	sensorService := sService.NewSensorService(&sensor)
	sensorHandler := sHttp.NewSensorHandler(sensorService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	sHttp.SetupRouter(r, sensorHandler)

	serviceAddress := os.Getenv("ADDRESS")
	if serviceAddress == "" {
		serviceAddress = ":8081"
	}
	log.Fatal(r.Run(serviceAddress))
}

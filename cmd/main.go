package main

import (
	"log"
	"os"

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
	sHttp.SetupRouter(r, sensorHandler)

	serviceAddress := os.Getenv("ADDRESS")
	if serviceAddress == "" {
		serviceAddress = ":8080"
	}
	log.Printf("Starting service at %s\n", serviceAddress)
	log.Fatal(r.Run(serviceAddress))
}

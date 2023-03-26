package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/sensor/pkg"
)

func main() {
	// Initialization of the sensor
	sensor := pkg.Sensor{Enabled: false, Detected: false}

	sensorService := pkg.NewSensorService(&sensor)
	sensorHandler := pkg.NewSensorHandler(sensorService)
	r := gin.Default()
	pkg.SetupRouter(r, sensorHandler)

	// Gets a service address from the environment variable or uses the default one
	// serviceAddress := viper.GetString("ADDRESS")
	serviceAddress := os.Getenv("ADDRESS")
	if serviceAddress == "" {
		serviceAddress = ":8080"
	}
	log.Printf("Starting service at %s\n", serviceAddress)
	log.Fatal(r.Run(serviceAddress))
}

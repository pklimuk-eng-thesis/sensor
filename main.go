package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/presence-sensor/pkg"
)

func main() {
	// Initialization of the sensor
	presenceSensor := pkg.PresenceSensor{SensorEnabled: false, PresenceDetected: false}

	presenceSensorService := pkg.NewPresenceSensorService(&presenceSensor)
	presenceSensorHandler := pkg.NewPresenceSensorHandler(presenceSensorService)
	r := gin.Default()
	pkg.SetupRouter(r, presenceSensorHandler)

	// Gets a service address from the environment variable or uses the default one
	// serviceAddress := viper.GetString("ADDRESS")
	serviceAddress := os.Getenv("ADDRESS")
	if serviceAddress == "" {
		serviceAddress = ":8080"
	}
	log.Printf("Starting service at %s\n", serviceAddress)
	log.Fatal(r.Run(serviceAddress))
}

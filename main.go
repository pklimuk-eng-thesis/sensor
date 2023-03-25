package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/presence-sensor/pkg"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Initialization of the sensor
	presenceSensor := pkg.PresenceSensor{SensorEnabled: false, PresenceDetected: false}

	presenceSensorService := pkg.NewPresenceSensorService(&presenceSensor)
	presenceSensorHandler := pkg.NewPresenceSensorHandler(presenceSensorService)
	r := gin.Default()
	pkg.SetupRouter(r, presenceSensorHandler)

	serviceAddress := viper.GetString("ADDRESS")
	log.Printf("Starting service at %s\n", serviceAddress)
	log.Fatal(r.Run(serviceAddress))
}

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

	r := gin.Default()
	r.GET("/isPresenceDetected", pkg.IsPresenceDetected)
	r.GET("/isSensorEnabled", pkg.IsSensorEnabled)
	r.POST("/setPresenceDetected", pkg.SetPresenceDetected)
	r.POST("/setSensorEnabled", pkg.SetSensorEnabled)

	log.Fatal(r.Run(viper.GetString("ADDRESS")))
}

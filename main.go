package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/presence-sensor/pkg"
)

func main() {
	r := gin.Default()
	r.GET("/isPresenceDetected", pkg.IsPresenceDetected)
	r.GET("/isSensorEnabled", pkg.IsSensorEnabled)
	r.POST("/setPresenceDetected", pkg.SetPresenceDetected)
	r.POST("/setSensorEnabled", pkg.SetSensorEnabled)
	r.Run("localhost:8080")
}

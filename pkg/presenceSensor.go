package pkg

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PresenceSensor struct {
	PresenceDetected bool `json:"presenceDetected"`
	SensorEnabled    bool `json:"sensorEnabled"`
}

// PresenceSensor1 is a global variable that holds the status of the sensor.
var PresenceSensor1 = &PresenceSensor{false, true}

func IsPresenceDetected(c *gin.Context) {
	if isEnabled := PresenceSensor1.SensorEnabled; !isEnabled {
		c.String(http.StatusBadRequest, "Sensor is disabled")
		return
	}
	c.IndentedJSON(http.StatusOK, &PresenceSensor1.PresenceDetected)
}

func IsSensorEnabled(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, &PresenceSensor1.SensorEnabled)
}

func SetPresenceDetected(c *gin.Context) {
	if isEnabled := PresenceSensor1.SensorEnabled; !isEnabled {
		c.String(http.StatusBadRequest, "Sensor is disabled")
		return
	}
	presenceDetected, err := strconv.ParseBool(c.Query("presenceDetected"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	PresenceSensor1.PresenceDetected = presenceDetected
	c.IndentedJSON(http.StatusOK, &PresenceSensor1.PresenceDetected)
}

func SetSensorEnabled(c *gin.Context) {
	sensorEnabled, err := strconv.ParseBool(c.Query("sensorEnabled"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	PresenceSensor1.SensorEnabled = sensorEnabled
	c.IndentedJSON(http.StatusOK, &PresenceSensor1.SensorEnabled)
}

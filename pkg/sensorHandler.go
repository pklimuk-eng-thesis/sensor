package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorHandler struct {
	service SensorService
}

func NewSensorHandler(service SensorService) *SensorHandler {
	return &SensorHandler{service}
}

func (h *SensorHandler) Detected(c *gin.Context) {
	detected, err := h.service.Detected()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &detected)
}

func (h *SensorHandler) IsSensorEnabled(c *gin.Context) {
	sensorEnabled := h.service.IsSensorEnabled()
	c.IndentedJSON(http.StatusOK, &sensorEnabled)
}

func (h *SensorHandler) ToggleDetected(c *gin.Context) {
	detected, err := h.service.ToggleDetected()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &detected)
}

func (h *SensorHandler) ToggleSensorEnabled(c *gin.Context) {
	sensorEnabled := h.service.ToggleSensorEnabled()
	c.IndentedJSON(http.StatusOK, &sensorEnabled)
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sService "github.com/pklimuk-eng-thesis/sensor/pkg/service"
)

type SensorHandler struct {
	service sService.SensorService
}

func NewSensorHandler(service sService.SensorService) *SensorHandler {
	return &SensorHandler{service}
}

func (h *SensorHandler) GetInfo(c *gin.Context) {
	sensor := h.service.GetInfo()
	c.IndentedJSON(http.StatusOK, &sensor)
}

func (h *SensorHandler) ToggleDetected(c *gin.Context) {
	sensor, err := h.service.ToggleDetected()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &sensor)
}

func (h *SensorHandler) ToggleSensorEnabled(c *gin.Context) {
	sensor := h.service.ToggleSensorEnabled()
	c.IndentedJSON(http.StatusOK, &sensor)
}
